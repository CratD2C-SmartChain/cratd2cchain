package engine_v2

import (
	"math/big"

	"github.com/CratD2C-SmartChain/cratd2cchain/common"
	"github.com/CratD2C-SmartChain/cratd2cchain/consensus"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/types"
)

// TODO: what should be new difficulty
func (x *CratD2C_v2) calcDifficulty(chain consensus.ChainReader, parent *types.Header, signer common.Address) *big.Int {
	return big.NewInt(1)
}
