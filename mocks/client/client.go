// Code generated by mockery v2.12.3. DO NOT EDIT.

package client

import (
	big "math/big"

	api "github.com/ava-labs/avalanchego/api"

	client "github.com/ava-labs/avalanche-rosetta/client"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	info "github.com/ava-labs/avalanchego/api/info"

	interfaces "github.com/ava-labs/coreth/interfaces"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/ava-labs/avalanchego/utils/rpc"

	types "github.com/ava-labs/coreth/core/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BalanceAt provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) BalanceAt(_a0 context.Context, _a1 common.Address, _a2 *big.Int) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByHash provides a mock function with given fields: _a0, _a1
func (_m *Client) BlockByHash(_a0 context.Context, _a1 common.Hash) (*types.Block, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: _a0, _a1
func (_m *Client) BlockByNumber(_a0 context.Context, _a1 *big.Int) (*types.Block, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallContract provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) CallContract(_a0 context.Context, _a1 interfaces.CallMsg, _a2 *big.Int) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.CallMsg, *big.Int) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.CallMsg, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainID provides a mock function with given fields: _a0
func (_m *Client) ChainID(_a0 context.Context) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: _a0, _a1
func (_m *Client) EstimateGas(_a0 context.Context, _a1 interfaces.CallMsg) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.CallMsg) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.CallMsg) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAtomicUTXOs provides a mock function with given fields: ctx, addrs, sourceChain, limit, startAddress, startUTXOID
func (_m *Client) GetAtomicUTXOs(ctx context.Context, addrs []string, sourceChain string, limit uint32, startAddress string, startUTXOID string) ([][]byte, api.Index, error) {
	ret := _m.Called(ctx, addrs, sourceChain, limit, startAddress, startUTXOID)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, uint32, string, string) [][]byte); ok {
		r0 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 api.Index
	if rf, ok := ret.Get(1).(func(context.Context, []string, string, uint32, string, string) api.Index); ok {
		r1 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID)
	} else {
		r1 = ret.Get(1).(api.Index)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, string, uint32, string, string) error); ok {
		r2 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetContractInfo provides a mock function with given fields: _a0, _a1
func (_m *Client) GetContractInfo(_a0 common.Address, _a1 bool) (string, uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(common.Address, bool) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 uint8
	if rf, ok := ret.Get(1).(func(common.Address, bool) uint8); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(uint8)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(common.Address, bool) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNetworkName provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNetworkName(_a0 context.Context, _a1 ...rpc.Option) (string, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByHash provides a mock function with given fields: _a0, _a1
func (_m *Client) HeaderByHash(_a0 context.Context, _a1 common.Hash) (*types.Header, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Header); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: _a0, _a1
func (_m *Client) HeaderByNumber(_a0 context.Context, _a1 *big.Int) (*types.Header, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBootstrapped provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) IsBootstrapped(_a0 context.Context, _a1 string, _a2 ...rpc.Option) (bool, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, ...rpc.Option) bool); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NonceAt provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) NonceAt(_a0 context.Context, _a1 common.Address, _a2 *big.Int) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peers provides a mock function with given fields: _a0, _a1
func (_m *Client) Peers(_a0 context.Context, _a1 ...rpc.Option) ([]info.Peer, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []info.Peer
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) []info.Peer); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]info.Peer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: _a0, _a1
func (_m *Client) SendTransaction(_a0 context.Context, _a1 *types.Transaction) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SuggestGasPrice provides a mock function with given fields: _a0
func (_m *Client) SuggestGasPrice(_a0 context.Context) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TraceBlockByHash provides a mock function with given fields: _a0, _a1
func (_m *Client) TraceBlockByHash(_a0 context.Context, _a1 string) ([]*client.Call, [][]*client.FlatCall, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*client.Call
	if rf, ok := ret.Get(0).(func(context.Context, string) []*client.Call); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.Call)
		}
	}

	var r1 [][]*client.FlatCall
	if rf, ok := ret.Get(1).(func(context.Context, string) [][]*client.FlatCall); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]*client.FlatCall)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TraceTransaction provides a mock function with given fields: _a0, _a1
func (_m *Client) TraceTransaction(_a0 context.Context, _a1 string) (*client.Call, []*client.FlatCall, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *client.Call
	if rf, ok := ret.Get(0).(func(context.Context, string) *client.Call); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Call)
		}
	}

	var r1 []*client.FlatCall
	if rf, ok := ret.Get(1).(func(context.Context, string) []*client.FlatCall); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*client.FlatCall)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionByHash provides a mock function with given fields: _a0, _a1
func (_m *Client) TransactionByHash(_a0 context.Context, _a1 common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionReceipt provides a mock function with given fields: _a0, _a1
func (_m *Client) TransactionReceipt(_a0 context.Context, _a1 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxPoolContent provides a mock function with given fields: _a0
func (_m *Client) TxPoolContent(_a0 context.Context) (*client.TxPoolContent, error) {
	ret := _m.Called(_a0)

	var r0 *client.TxPoolContent
	if rf, ok := ret.Get(0).(func(context.Context) *client.TxPoolContent); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.TxPoolContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewClientT interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t NewClientT) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
