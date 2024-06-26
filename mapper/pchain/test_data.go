package pchain

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/validator"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func buildImport() (*txs.ImportTx, map[string]*types.AccountIdentifier) {
	avaxAssetID, _ := ids.FromString("U8iRqJoiJm8xZHAacmvYyZVwqQx6uDNtQeP3CQ6fcgQk3JqnK")
	sourceChain, _ := ids.FromString("2JVSBoinj9C2J33VntvzYtVJNZdN2NKiwwKjcumHUWEb5DbBrm")
	outAddr1, _ := address.ParseToID("P-fuji1xm0r37l6gyf2mly4pmzc0tz6wnwqkugedh95fk")
	outAddr2, _ := address.ParseToID("P-fuji1fmragvegm5k26qzlt6vy0ghhdr508u6r4a5rxj")
	outAddr3, _ := address.ParseToID("P-fuji1j3sw805usytrsymfwxxrcwfqguyarumn45cllj")
	importAddr, _ := address.ParseToID("C-fuji1xm0r37l6gyf2mly4pmzc0tz6wnwqkugedh95fk")
	importedTxID, _ := ids.FromString("2DtYhzCvo9LRYMRJ6sCtYJ4aNPRpsibp46ETNyY6H5Cox1VLvX")
	impTx := &txs.ImportTx{
		BaseTx: txs.BaseTx{
			BaseTx: avax.BaseTx{
				NetworkID:    uint32(5),
				BlockchainID: [32]byte{},
				Outs: []*avax.TransferableOutput{
					{
						Asset: avax.Asset{ID: avaxAssetID},
						FxID:  [32]byte{},
						Out: &secp256k1fx.TransferOutput{
							Amt: 8000000,
							OutputOwners: secp256k1fx.OutputOwners{
								Locktime:  0,
								Threshold: 1,
								Addrs:     []ids.ShortID{outAddr1},
							},
						},
					},
					{ //  this will be skipped as it is multisig
						Asset: avax.Asset{ID: avaxAssetID},
						FxID:  [32]byte{},
						Out: &secp256k1fx.TransferOutput{
							Amt: 8000000,
							OutputOwners: secp256k1fx.OutputOwners{
								Locktime:  0,
								Threshold: 2,
								Addrs:     []ids.ShortID{outAddr1, outAddr2, outAddr3},
							},
						},
					},
					{ //  this will be skipped as it does not have any addresses
						Asset: avax.Asset{ID: avaxAssetID},
						FxID:  [32]byte{},
						Out: &secp256k1fx.TransferOutput{
							Amt: 1000000,
							OutputOwners: secp256k1fx.OutputOwners{
								Locktime:  0,
								Threshold: 0,
								Addrs:     []ids.ShortID{},
							},
						},
					},
				},
				Ins:  nil,
				Memo: []byte{},
			},
		},
		SourceChain: sourceChain,
		ImportedInputs: []*avax.TransferableInput{{
			UTXOID: avax.UTXOID{
				TxID:        importedTxID,
				OutputIndex: 0,
				Symbol:      false,
			},
			Asset: avax.Asset{ID: avaxAssetID},
			FxID:  [32]byte{},
			In: &secp256k1fx.TransferInput{
				Amt: 9000000,
				Input: secp256k1fx.Input{
					SigIndices: []uint32{},
				},
			},
		}},
	}

	inputTxAccounts := map[string]*types.AccountIdentifier{}
	inputTxAccounts[impTx.ImportedInputs[0].String()] = &types.AccountIdentifier{Address: importAddr.String()}

	return impTx, inputTxAccounts
}

func buildExport() (*txs.ExportTx, map[string]*types.AccountIdentifier) {
	avaxAssetID, _ := ids.FromString("U8iRqJoiJm8xZHAacmvYyZVwqQx6uDNtQeP3CQ6fcgQk3JqnK")
	outAddr, _ := address.ParseToID("P-fuji1wmd9dfrqpud6daq0cde47u0r7pkrr46ep60399")
	exportOutAddr, _ := address.ParseToID("P-fuji1wmd9dfrqpud6daq0cde47u0r7pkrr46ep60399")
	txID, _ := ids.FromString("27LaDkrUrMY1bhVf2i8RARCrRwFjeRw7vEu8ntLQXracgLzL1v")
	destinationID, _ := ids.FromString("yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp")
	exTx := &txs.ExportTx{
		BaseTx: txs.BaseTx{
			BaseTx: avax.BaseTx{
				NetworkID:    uint32(5),
				BlockchainID: [32]byte{},
				Outs: []*avax.TransferableOutput{{
					Asset: avax.Asset{ID: avaxAssetID},
					FxID:  [32]byte{},
					Out: &secp256k1fx.TransferOutput{
						Amt: 2910137500,
						OutputOwners: secp256k1fx.OutputOwners{
							Locktime:  0,
							Threshold: 1,
							Addrs:     []ids.ShortID{outAddr},
						},
					},
				}},
				Ins: []*avax.TransferableInput{{
					UTXOID: avax.UTXOID{TxID: txID, OutputIndex: 0, Symbol: false},
					Asset:  avax.Asset{ID: avaxAssetID},
					FxID:   [32]byte{},
					In: &secp256k1fx.TransferInput{
						Amt:   2921137500,
						Input: secp256k1fx.Input{SigIndices: []uint32{}},
					},
				}},
				Memo: []byte{},
			},
		},
		DestinationChain: destinationID,
		ExportedOutputs: []*avax.TransferableOutput{{
			Asset: avax.Asset{ID: avaxAssetID},
			FxID:  [32]byte{},
			Out: &secp256k1fx.TransferOutput{
				Amt: 10000000,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  0,
					Threshold: 1,
					Addrs:     []ids.ShortID{exportOutAddr},
				},
			},
		}},
	}

	inputTxAccounts := map[string]*types.AccountIdentifier{}
	inputTxAccounts[exTx.Ins[0].String()] = &types.AccountIdentifier{Address: outAddr.String()}

	return exTx, inputTxAccounts
}

func buildAddDelegator() (*txs.AddDelegatorTx, map[string]*types.AccountIdentifier) {
	avaxAssetID, _ := ids.FromString("U8iRqJoiJm8xZHAacmvYyZVwqQx6uDNtQeP3CQ6fcgQk3JqnK")
	txID, _ := ids.FromString("2JQGX1MBdszAaeV6eApCZm7CBpc917qWiyQ2cygFRJ6WteDkre")
	outAddr, _ := address.ParseToID("P-fuji1gdkq8g208e3j4epyjmx65jglsw7vauh86l47ac")
	validatorID, _ := ids.NodeIDFromString("NodeID-BFa1padLXBj7VHa2JYvYGzcTBPQGjPhUy")
	stakeAddr, _ := address.ParseToID("P-fuji1l022sue7g2kzvrcuxughl30xkss2cj0az3e5r2")
	rewardAddr, _ := address.ParseToID("P-fuji1l022sue7g2kzvrcuxughl30xkss2cj0az3e5r2")
	tx := &txs.AddDelegatorTx{
		BaseTx: txs.BaseTx{
			BaseTx: avax.BaseTx{
				NetworkID:    uint32(5),
				BlockchainID: [32]byte{},
				Outs: []*avax.TransferableOutput{{
					Asset: avax.Asset{ID: avaxAssetID},
					FxID:  [32]byte{},
					Out: &secp256k1fx.TransferOutput{
						Amt: 996649063,
						OutputOwners: secp256k1fx.OutputOwners{
							Locktime:  9,
							Threshold: 1,
							Addrs:     []ids.ShortID{outAddr},
						},
					},
				}},
				Ins: []*avax.TransferableInput{{
					UTXOID: avax.UTXOID{TxID: txID, OutputIndex: 0, Symbol: false},
					Asset:  avax.Asset{ID: avaxAssetID},
					FxID:   [32]byte{},
					In: &secp256k1fx.TransferInput{
						Amt:   1996649063,
						Input: secp256k1fx.Input{SigIndices: []uint32{}},
					},
				}},
				Memo: []byte{},
			},
		},
		Validator: validator.Validator{
			NodeID: validatorID,
			Start:  1656058022,
			End:    1657872569,
			Wght:   1000000000,
		},
		StakeOuts: []*avax.TransferableOutput{{
			Asset: avax.Asset{ID: avaxAssetID},
			FxID:  [32]byte{},
			Out: &secp256k1fx.TransferOutput{
				Amt: 1000000000,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  0,
					Threshold: 1,
					Addrs:     []ids.ShortID{stakeAddr},
				},
			},
		}},
		DelegationRewardsOwner: &secp256k1fx.OutputOwners{
			Locktime:  0,
			Threshold: 1,
			Addrs:     []ids.ShortID{rewardAddr},
		},
	}

	inputTxAccounts := map[string]*types.AccountIdentifier{}
	inputTxAccounts[tx.Ins[0].String()] = &types.AccountIdentifier{Address: stakeAddr.String()}

	return tx, inputTxAccounts
}

func buildValidatorTx() (*txs.AddValidatorTx, map[string]*types.AccountIdentifier) {
	avaxAssetID, _ := ids.FromString("U8iRqJoiJm8xZHAacmvYyZVwqQx6uDNtQeP3CQ6fcgQk3JqnK")

	txID, _ := ids.FromString("88tfp1Pkw9vyKrRtVNiMrghFBrre6Q6CzqPW1t7StDNX9PJEo")
	stakeAddr, _ := address.ParseToID("P-fuji1ljdzyey6vu3hgn3cwg4j5lpy0svd6arlxpj6je")
	rewardAddr, _ := address.ParseToID("P-fuji1ljdzyey6vu3hgn3cwg4j5lpy0svd6arlxpj6je")
	validatorID, _ := ids.NodeIDFromString("NodeID-CCecHmRK3ANe92VyvASxkNav26W4vAVpX")
	addvalidator := &txs.AddValidatorTx{
		BaseTx: txs.BaseTx{
			BaseTx: avax.BaseTx{
				NetworkID:    uint32(5),
				BlockchainID: [32]byte{},
				Outs:         nil,
				Ins: []*avax.TransferableInput{{
					UTXOID: avax.UTXOID{TxID: txID, OutputIndex: 0, Symbol: false},
					Asset:  avax.Asset{ID: avaxAssetID},
					FxID:   [32]byte{},
					In: &secp256k1fx.TransferInput{
						Amt:   2000000000,
						Input: secp256k1fx.Input{SigIndices: []uint32{1}},
					},
				}},
				Memo: []byte{},
			},
		},
		Validator: validator.Validator{
			NodeID: validatorID,
			Start:  1656084079,
			End:    1687620079,
			Wght:   2000000000,
		},
		StakeOuts: []*avax.TransferableOutput{{
			Asset: avax.Asset{ID: avaxAssetID},
			FxID:  [32]byte{},
			Out: &secp256k1fx.TransferOutput{
				Amt: 2000000000,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  0,
					Threshold: 1,
					Addrs:     []ids.ShortID{stakeAddr},
				},
			},
		}},
		RewardsOwner: &secp256k1fx.OutputOwners{
			Locktime:  0,
			Threshold: 1,
			Addrs:     []ids.ShortID{rewardAddr},
		},
		DelegationShares: 20000,
	}

	inputTxAccounts := map[string]*types.AccountIdentifier{}
	inputTxAccounts[addvalidator.Ins[0].String()] = &types.AccountIdentifier{Address: stakeAddr.String()}

	return addvalidator, inputTxAccounts
}
