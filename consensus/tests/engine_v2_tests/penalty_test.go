package engine_v2_tests

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/CratD2C-SmartChain/cratd2cchain/common"
	"github.com/CratD2C-SmartChain/cratd2cchain/consensus/CratD2C"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/types"
	"github.com/CratD2C-SmartChain/cratd2cchain/eth/hooks"
	"github.com/CratD2C-SmartChain/cratd2cchain/params"
	"github.com/stretchr/testify/assert"
)

func TestHookPenaltyV2TwoEpoch(t *testing.T) {
	b, err := json.Marshal(params.TestXDPoSMockChainConfig)
	assert.Nil(t, err)
	configString := string(b)

	var config params.ChainConfig
	err = json.Unmarshal([]byte(configString), &config)
	assert.Nil(t, err)
	// set V2 switch to 0
	config.CratD2C.V2.SwitchBlock.SetUint64(0)
	conf := &config
	blockchain, _, block1348, signer, signFn := PrepareXDCTestBlockChainWith128Candidates(t, int(config.CratD2C.Epoch+config.CratD2C.Gap)-2, conf)
	adaptor := blockchain.Engine().(*CratD2C.CratD2C)
	hooks.AttachConsensusV2Hooks(adaptor, blockchain, conf)
	assert.NotNil(t, adaptor.EngineV2.HookPenalty)
	header001 := blockchain.GetHeaderByNumber(1)
	masternodes := adaptor.GetMasternodesFromCheckpointHeader(header001)
	header450 := blockchain.GetHeaderByNumber(450)
	penalty, err := adaptor.EngineV2.HookPenalty(blockchain, header450.Number, header450.ParentHash, masternodes, config.CratD2C)
	assert.Nil(t, err)
	assert.Equal(t, 125, len(penalty)) // we have 4 miners created for all blocks, 3 is among 128 masternode candidates (test issue, no need to fix). So 125 candidates left are penalties

	header1335 := blockchain.GetHeaderByNumber(config.CratD2C.Epoch + config.CratD2C.Gap - uint64(common.MergeSignRange))
	tx, err := signingTxWithKey(header1335, 0, voterKey)
	assert.Nil(t, err)
	block1349 := CreateBlock(blockchain, conf, block1348, 1349, 1358, signer.Hex(), signer, signFn, nil, nil, "c2bf7b59be5184fc1148be5db14692b2dc89a1b345895d3e8d0ee7b8a7607450")
	err = blockchain.InsertBlock(block1349)
	assert.Nil(t, err)
	adaptor.CacheSigningTxs(block1349.Hash(), []*types.Transaction{tx})
	// the following penalty across two epochs: 1349 - 901, 900 - 450
	penalty, err = adaptor.EngineV2.HookPenalty(blockchain, big.NewInt(0).Add(block1349.Header().Number, big.NewInt(1)), block1349.Hash(), masternodes, config.CratD2C)
	assert.Nil(t, err)
	assert.Equal(t, 124, len(penalty)) // 1 master node has signing tx cached, so it comes back. 125-1 candidates are penalties
}
