// Copyright 2016 The go-ethereum Authors
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

package ethclient

import "github.com/CratD2C-SmartChain/cratd2cchain"

// Verify that Client implements the ethereum interfaces.
var (
	_ = XDPoSChain.ChainReader(&Client{})
	_ = XDPoSChain.TransactionReader(&Client{})
	_ = XDPoSChain.ChainStateReader(&Client{})
	_ = XDPoSChain.ChainSyncReader(&Client{})
	_ = XDPoSChain.ContractCaller(&Client{})
	_ = XDPoSChain.GasEstimator(&Client{})
	_ = XDPoSChain.GasPricer(&Client{})
	_ = XDPoSChain.LogFilterer(&Client{})
	_ = XDPoSChain.PendingStateReader(&Client{})
	// _ = ethereum.PendingStateEventer(&Client{})
	_ = XDPoSChain.PendingContractCaller(&Client{})
)
