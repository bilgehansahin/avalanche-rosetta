package pchain

import (
	"context"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks"
	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/ava-labs/avalanche-rosetta/client"
	"github.com/ava-labs/avalanche-rosetta/mapper"
	"github.com/ava-labs/avalanche-rosetta/service"
	"github.com/ava-labs/avalanche-rosetta/service/backend/pchain/indexer"
)

var (
	_ service.ConstructionBackend = &Backend{}
	_ service.NetworkBackend      = &Backend{}
	_ service.AccountBackend      = &Backend{}
	_ service.BlockBackend        = &Backend{}
)

type Backend struct {
	fac                    *crypto.FactorySECP256K1R
	networkID              *types.NetworkIdentifier
	networkHRP             string
	pClient                client.PChainClient
	indexerParser          indexer.Parser
	getUTXOsPageSize       uint32
	codec                  codec.Manager
	codecVersion           uint16
	genesisBlock           *indexer.ParsedGenesisBlock
	genesisBlockIdentifier *types.BlockIdentifier
	chainIDs               map[string]string
	avaxAssetID            ids.ID
}

// NewBackend creates a P-chain service backend
func NewBackend(
	pClient client.PChainClient,
	indexerParser indexer.Parser,
	assetID ids.ID,
	networkIdentifier *types.NetworkIdentifier,
) (*Backend, error) {
	backEnd := &Backend{
		fac:              &crypto.FactorySECP256K1R{},
		networkID:        networkIdentifier,
		pClient:          pClient,
		getUTXOsPageSize: 1024,
		codec:            blocks.Codec,
		codecVersion:     blocks.Version,
		indexerParser:    indexerParser,
		avaxAssetID:      assetID,
	}

	err := backEnd.initChainIDs()
	if err != nil {
		return nil, err
	}

	backEnd.networkHRP, err = mapper.GetHRP(networkIdentifier)
	if err != nil {
		return nil, err
	}
	return backEnd, err
}

// ShouldHandleRequest returns whether a given request should be handled by this backend
func (b *Backend) ShouldHandleRequest(req interface{}) bool {
	switch r := req.(type) {
	case *types.AccountBalanceRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.AccountCoinsRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.BlockRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.BlockTransactionRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionDeriveRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionMetadataRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionPreprocessRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionPayloadsRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionParseRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionCombineRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionHashRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.ConstructionSubmitRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	case *types.NetworkRequest:
		return b.isNetworkSupported(r.NetworkIdentifier) && b.isPChain(r.NetworkIdentifier)
	}

	return false
}

func (b *Backend) getGenesisBlock(ctx context.Context) (*indexer.ParsedGenesisBlock, error) {
	if b.genesisBlock != nil {
		return b.genesisBlock, nil
	}
	genesisBlock, err := b.indexerParser.GetGenesisBlock(ctx)
	if err != nil {
		return nil, err
	}
	b.genesisBlock = genesisBlock
	b.genesisBlockIdentifier = &types.BlockIdentifier{
		Index: int64(genesisBlock.Height),
		Hash:  genesisBlock.BlockID.String(),
	}
	return genesisBlock, nil
}

func (b *Backend) initChainIDs() error {
	ctx := context.Background()
	b.chainIDs = map[string]string{
		ids.Empty.String(): mapper.PChainNetworkIdentifier,
	}

	cChainID, err := b.pClient.GetBlockchainID(ctx, mapper.CChainNetworkIdentifier)
	if err != nil {
		return err
	}
	b.chainIDs[cChainID.String()] = mapper.CChainNetworkIdentifier

	xChainID, err := b.pClient.GetBlockchainID(ctx, mapper.XChainNetworkIdentifier)
	if err != nil {
		return err
	}
	b.chainIDs[xChainID.String()] = mapper.XChainNetworkIdentifier

	return nil
}

// isPChain checks network identifier to make sure sub-network identifier set to "P"
func (b *Backend) isPChain(reqNetworkID *types.NetworkIdentifier) bool {
	return reqNetworkID != nil &&
		reqNetworkID.SubNetworkIdentifier != nil &&
		reqNetworkID.SubNetworkIdentifier.Network == mapper.PChainNetworkIdentifier
}

func (b *Backend) isNetworkSupported(reqNetworkID *types.NetworkIdentifier) bool {
	return reqNetworkID != nil && reqNetworkID.Network == b.networkID.Network
}
