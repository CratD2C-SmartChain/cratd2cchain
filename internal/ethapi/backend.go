// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package ethapi implements the general Ethereum API functions.
package ethapi

import (
	"context"
	"math/big"

	"github.com/CratD2C-SmartChain/cratd2cchain/XDCx/tradingstate"
	"github.com/CratD2C-SmartChain/cratd2cchain/XDCxlending"
	"github.com/CratD2C-SmartChain/cratd2cchain/accounts/abi/bind"

	"github.com/CratD2C-SmartChain/cratd2cchain/XDCx"

	"github.com/CratD2C-SmartChain/cratd2cchain/accounts"
	"github.com/CratD2C-SmartChain/cratd2cchain/common"
	"github.com/CratD2C-SmartChain/cratd2cchain/consensus"
	"github.com/CratD2C-SmartChain/cratd2cchain/core"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/state"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/types"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/vm"
	"github.com/CratD2C-SmartChain/cratd2cchain/eth/downloader"
	"github.com/CratD2C-SmartChain/cratd2cchain/ethdb"
	"github.com/CratD2C-SmartChain/cratd2cchain/event"
	"github.com/CratD2C-SmartChain/cratd2cchain/params"
	"github.com/CratD2C-SmartChain/cratd2cchain/rpc"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Ethereum API
	Downloader() *downloader.Downloader
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() ethdb.Database
	EventMux() *event.TypeMux
	AccountManager() *accounts.Manager
	XDCxService() *XDCx.XDCX
	LendingService() *XDCxlending.Lending

	// BlockChain API
	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetTd(blockHash common.Hash) *big.Int
	GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, XDCxState *tradingstate.TradingStateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error)
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription

	// TxPool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeTxPreEvent(chan<- core.TxPreEvent) event.Subscription

	// Order Pool Transaction
	SendOrderTx(ctx context.Context, signedTx *types.OrderTransaction) error
	OrderTxPoolContent() (map[common.Address]types.OrderTransactions, map[common.Address]types.OrderTransactions)
	OrderStats() (pending int, queued int)
	SendLendingTx(ctx context.Context, signedTx *types.LendingTransaction) error

	ChainConfig() *params.ChainConfig
	CurrentBlock() *types.Block
	GetIPCClient() (bind.ContractBackend, error)
	GetEngine() consensus.Engine
	GetRewardByHash(hash common.Hash) map[string]map[string]map[string]*big.Int

	GetVotersRewards(common.Address) map[common.Address]*big.Int
	GetVotersCap(checkpoint *big.Int, masterAddr common.Address, voters []common.Address) map[common.Address]*big.Int
	GetEpochDuration() *big.Int
	GetMasternodesCap(checkpoint uint64) map[common.Address]*big.Int
	GetBlocksHashCache(blockNr uint64) []common.Hash
	AreTwoBlockSamePath(newBlock common.Hash, oldBlock common.Hash) bool
	GetOrderNonce(address common.Hash) (uint64, error)
}

func GetAPIs(apiBackend Backend, chainReader consensus.ChainReader) []rpc.API {
	nonceLock := new(AddrLocker)
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend, chainReader),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "XDCx",
			Version:   "1.0",
			Service:   NewPublicXDCXTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(apiBackend),
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend, nonceLock),
			Public:    false,
		},
	}
}
