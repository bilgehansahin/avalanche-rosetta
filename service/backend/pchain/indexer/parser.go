package indexer

import (
	"context"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/hashing"

	pBlocks "github.com/ava-labs/avalanchego/vms/platformvm/blocks"
	pGenesis "github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	proposerBlk "github.com/ava-labs/avalanchego/vms/proposervm/block"

	"github.com/ava-labs/avalanche-rosetta/client"
	"github.com/ava-labs/avalanche-rosetta/mapper"
)

var genesisTimestamp = time.Date(2020, time.September, 10, 0, 0, 0, 0, time.UTC).Unix()

// Parser defines the interface for a P-chain indexer parser
type Parser interface {
	// GetGenesisBlock parses and returns the Genesis block
	GetGenesisBlock(ctx context.Context) (*ParsedGenesisBlock, error)
	// GetPlatformHeight returns the current block height of P-chain
	GetPlatformHeight(ctx context.Context) (uint64, error)
	// ParseCurrentBlock parses and returns the current tip of P-chain
	ParseCurrentBlock(ctx context.Context) (*ParsedBlock, error)
	// ParseBlockAtHeight parses and returns the block at the specified index
	ParseBlockAtHeight(ctx context.Context, height uint64) (*ParsedBlock, error)
	// ParseBlockWithHash parses and returns the block with the specified hash
	ParseBlockWithHash(ctx context.Context, hash string) (*ParsedBlock, error)
}

// Interface compliance
var _ Parser = &parser{}

type parser struct {
	networkID uint32
	aliaser   ids.Aliaser

	codec        codec.Manager
	codecVersion uint16

	pChainClient client.PChainClient
}

// NewParser creates a new P-chain indexer parser
func NewParser(pChainClient client.PChainClient) (Parser, error) {
	aliaser := ids.NewAliaser()
	err := aliaser.Alias(constants.PlatformChainID, mapper.PChainNetworkIdentifier)
	if err != nil {
		return nil, err
	}

	networkID, err := pChainClient.GetNetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &parser{
		codec:        pBlocks.Codec,
		codecVersion: pBlocks.Version,
		pChainClient: pChainClient,
		aliaser:      aliaser,
		networkID:    networkID,
	}, nil
}

func (p *parser) GetPlatformHeight(ctx context.Context) (uint64, error) {
	return p.pChainClient.GetHeight(ctx)
}

func (p *parser) GetGenesisBlock(ctx context.Context) (*ParsedGenesisBlock, error) {
	bytes, _, err := genesis.FromConfig(genesis.GetConfig(p.networkID))
	if err != nil {
		return nil, err
	}

	genesisState, err := pGenesis.Parse(bytes)
	if err != nil {
		return nil, err
	}

	genesisTimestamp := time.Unix(int64(genesisState.Timestamp), 0)

	var genesisTxs []*txs.Tx
	genesisTxs = append(genesisTxs, genesisState.Validators...)
	genesisTxs = append(genesisTxs, genesisState.Chains...)

	// Genesis commit block's parent ID is the hash of genesis state
	var genesisParentID ids.ID = hashing.ComputeHash256Array(bytes)

	// Genesis Block is not indexed by the indexer, but its block ID can be accessed from block 0's parent id
	genesisChildBlock, err := p.ParseBlockAtHeight(ctx, 1)
	if err != nil {
		return nil, err
	}

	genesisCtx := &snow.Context{
		BCLookup:  p.aliaser,
		NetworkID: p.networkID,
	}
	for _, utxo := range genesisState.UTXOs {
		utxo.UTXO.Out.InitCtx(genesisCtx)
	}

	genesisBlockID := genesisChildBlock.ParentID

	return &ParsedGenesisBlock{
		ParsedBlock: ParsedBlock{
			ParentID:  genesisParentID,
			Height:    0,
			BlockID:   genesisBlockID,
			BlockType: "GenesisBlock",
			Timestamp: genesisTimestamp.UnixMilli(),
			Txs:       genesisTxs,
			Proposer:  Proposer{},
		},
		GenesisBlockData: GenesisBlockData{
			Message:       genesisState.Message,
			InitialSupply: genesisState.InitialSupply,
			UTXOs:         genesisState.UTXOs,
		},
	}, nil
}

func (p *parser) ParseCurrentBlock(ctx context.Context) (*ParsedBlock, error) {
	height, err := p.GetPlatformHeight(ctx)
	if err != nil {
		return nil, err
	}

	return p.ParseBlockAtHeight(ctx, height)
}

func (p *parser) ParseBlockAtHeight(ctx context.Context, height uint64) (*ParsedBlock, error) {
	// P-chain indexer does not include genesis and store block at height 1 with index 0.
	// Therefore containers are looked up with index = height - 1.
	// Note that genesis does not cause a problem here as it is handled in a separate code path
	container, err := p.pChainClient.GetContainerByIndex(ctx, height-1)
	if err != nil {
		return nil, err
	}

	return p.parseBlockBytes(container.Bytes)
}

func (p *parser) ParseBlockWithHash(ctx context.Context, hash string) (*ParsedBlock, error) {
	hashID, err := ids.FromString(hash)
	if err != nil {
		return nil, err
	}

	container, err := p.pChainClient.GetContainerByID(ctx, hashID)
	if err != nil {
		return nil, err
	}

	return p.parseBlockBytes(container.Bytes)
}

func (p *parser) parseBlockBytes(proposerBytes []byte) (*ParsedBlock, error) {
	proposer, bytes, err := getProposerFromBytes(proposerBytes)
	if err != nil {
		return nil, fmt.Errorf("fetching proposer from block bytes errored with %w", err)
	}

	blk, err := pBlocks.Parse(p.codec, bytes)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling block bytes errored with %w", err)
	}

	parsedBlock := &ParsedBlock{
		Height:    blk.Height(),
		BlockID:   blk.ID(),
		BlockType: fmt.Sprintf("%T", blk),
		Proposer:  proposer,
	}

	blockTimestamp := time.Time{}

	switch castBlk := blk.(type) {
	case *pBlocks.ApricotProposalBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = castBlk.Txs()

		// If the block has an advance time tx, use its timestamp as the block timestamp
		for _, tx := range parsedBlock.Txs {
			if att, ok := tx.Unsigned.(*txs.AdvanceTimeTx); ok {
				blockTimestamp = att.Timestamp()
				break
			}
		}

	case *pBlocks.BanffProposalBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = castBlk.Txs()
		parsedBlock.Txs = append(parsedBlock.Txs, castBlk.Transactions...)
		blockTimestamp = castBlk.Timestamp()

	case *pBlocks.ApricotAtomicBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = castBlk.Txs()

	case *pBlocks.ApricotStandardBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = castBlk.Transactions

	case *pBlocks.BanffStandardBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = castBlk.Transactions
		blockTimestamp = castBlk.Timestamp()

	case *pBlocks.ApricotAbortBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = []*txs.Tx{}

	case *pBlocks.BanffAbortBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = []*txs.Tx{}
		blockTimestamp = castBlk.Timestamp()

	case *pBlocks.BanffCommitBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = []*txs.Tx{}
		blockTimestamp = castBlk.Timestamp()

	case *pBlocks.ApricotCommitBlock:
		parsedBlock.ParentID = castBlk.PrntID
		parsedBlock.Txs = []*txs.Tx{}

	default:
		return nil, fmt.Errorf("no handler exists for block type %T", castBlk)
	}

	// If no timestamp was found in a given block (pre-Banff) we fallback to proposer timestamp as used by Snowman++
	// if available.
	//
	// For pre-Snowman++, proposer timestamp does not exist either. In that case, fallback to the genesis timestamp.
	// This is needed as opposed to simply return 0 timestamp as Rosetta validation expects timestamps to be available
	// after 1/1/2000.
	if blockTimestamp.IsZero() {
		if proposer.Timestamp > genesisTimestamp {
			blockTimestamp = time.Unix(proposer.Timestamp, 0)
		} else {
			blockTimestamp = time.Unix(genesisTimestamp, 0)
		}
	}
	parsedBlock.Timestamp = blockTimestamp.UnixMilli()

	return parsedBlock, nil
}

func getProposerFromBytes(bytes []byte) (Proposer, []byte, error) {
	proposer, _, err := proposerBlk.Parse(bytes)
	if err != nil || proposer == nil {
		return Proposer{}, bytes, nil
	}

	switch castBlock := proposer.(type) {
	case proposerBlk.SignedBlock:
		return Proposer{
			ID:           castBlock.ID(),
			NodeID:       castBlock.Proposer(),
			PChainHeight: castBlock.PChainHeight(),
			Timestamp:    castBlock.Timestamp().Unix(),
			ParentID:     castBlock.ParentID(),
		}, castBlock.Block(), nil
	case proposerBlk.Block:
		return Proposer{}, castBlock.Block(), nil
	default:
		return Proposer{}, bytes, fmt.Errorf("no handler exists for proposer block type %T", castBlock)
	}
}
