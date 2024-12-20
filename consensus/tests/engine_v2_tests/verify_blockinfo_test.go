package engine_v2_tests

import (
	"fmt"
	"testing"

	"github.com/CratD2C-SmartChain/cratd2cchain/consensus/CratD2C"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/types"
	"github.com/CratD2C-SmartChain/cratd2cchain/params"
	"github.com/stretchr/testify/assert"
)

func TestShouldVerifyBlockInfo(t *testing.T) {
	// Block 1 is the first v2 block with round of 1
	blockchain, _, currentBlock, signer, signFn, _ := PrepareXDCTestBlockChainForV2Engine(t, 1, params.TestXDPoSMockChainConfig, nil)
	engineV2 := blockchain.Engine().(*CratD2C.CratD2C).EngineV2

	blockInfo := &types.BlockInfo{
		Hash:   currentBlock.Hash(),
		Round:  types.Round(1),
		Number: currentBlock.Number(),
	}
	err := engineV2.VerifyBlockInfo(blockchain, blockInfo, nil)
	assert.Nil(t, err)

	// Insert another Block, but it won't trigger commit
	blockNum := 2
	blockCoinBase := fmt.Sprintf("0x111000000000000000000000000000000%03d", blockNum)
	block2 := CreateBlock(blockchain, params.TestXDPoSMockChainConfig, currentBlock, blockNum, 2, blockCoinBase, signer, signFn, nil, nil, "")
	err = blockchain.InsertBlock(block2)
	assert.Nil(t, err)

	blockInfo = &types.BlockInfo{
		Hash:   block2.Hash(),
		Round:  types.Round(2),
		Number: block2.Number(),
	}
	err = engineV2.VerifyBlockInfo(blockchain, blockInfo, nil)
	assert.Nil(t, err)

	blockInfo = &types.BlockInfo{
		Hash:   currentBlock.Hash(),
		Round:  types.Round(2),
		Number: currentBlock.Number(),
	}
	err = engineV2.VerifyBlockInfo(blockchain, blockInfo, nil)
	assert.NotNil(t, err)

	blockInfo = &types.BlockInfo{
		Hash:   block2.Hash(),
		Round:  types.Round(3),
		Number: block2.Number(),
	}
	err = engineV2.VerifyBlockInfo(blockchain, blockInfo, nil)
	assert.NotNil(t, err)

	blockInfo = &types.BlockInfo{
		Hash:   block2.Hash(),
		Round:  types.Round(2),
		Number: currentBlock.Number(),
	}
	err = engineV2.VerifyBlockInfo(blockchain, blockInfo, nil)
	assert.NotNil(t, err)
}
