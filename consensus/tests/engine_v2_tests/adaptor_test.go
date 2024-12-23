package engine_v2_tests

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/CratD2C-SmartChain/cratd2cchain/consensus/CratD2C"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/types"
	"github.com/CratD2C-SmartChain/cratd2cchain/params"
	"github.com/stretchr/testify/assert"
)

/*
func TestAdaptorShouldGetAuthorForDifferentConsensusVersion(t *testing.T) {
	blockchain, _, currentBlock, signer, signFn, _ := PrepareXDCTestBlockChainForV2Engine(t, 900, params.TestXDPoSMockChainConfig, nil)
	adaptor := blockchain.Engine().(*CratD2C.CratD2C)

	addressFromAdaptor, errorAdaptor := adaptor.Author(currentBlock.Header())
	if errorAdaptor != nil {
		t.Fatalf("Failed while trying to get Author from adaptor")
	}
	addressFromV1Engine, errV1 := adaptor.EngineV1.Author(currentBlock.Header())
	if errV1 != nil {
		t.Fatalf("Failed while trying to get Author from engine v1")
	}
	// Make sure the value is exactly the same as from V1 engine
	assert.Equal(t, addressFromAdaptor, addressFromV1Engine)

	// Insert one more block to make it above 10, which means now we are on v2 of consensus engine
	// Insert block 901

	merkleRoot := "b3e34cf1d3d80bcd2c5add880842892733e45979ddaf16e531f660fdf7ca5787"
	header := &types.Header{
		Root:       common.HexToHash(merkleRoot),
		Number:     big.NewInt(int64(901)),
		ParentHash: currentBlock.Hash(),
		Coinbase:   signer,
	}

	header.Extra = generateV2Extra(1, currentBlock, signer, signFn, nil)

	block901, err := createBlockFromHeader(blockchain, header, nil, signer, signFn, blockchain.Config())
	if err != nil {
		t.Fatal(err)
	}
	err = blockchain.InsertBlock(block901)
	assert.Nil(t, err)

	addressFromAdaptor, errorAdaptor = adaptor.Author(block901.Header())
	if errorAdaptor != nil {
		t.Fatalf("Failed while trying to get Author from adaptor")
	}
	addressFromV2Engine, errV2 := adaptor.EngineV2.Author(block901.Header())
	if errV2 != nil {
		t.Fatalf("Failed while trying to get Author from engine v2")
	}
	// Make sure the value is exactly the same as from V2 engine
	assert.Equal(t, addressFromAdaptor, addressFromV2Engine)
}
*/

/*
	func TestAdaptorGetMasternodesFromCheckpointHeader(t *testing.T) {
		blockchain, _, currentBlock, _, _, _ := PrepareXDCTestBlockChainForV2Engine(t, 1, params.TestXDPoSMockChainConfig, nil)
		adaptor := blockchain.Engine().(*CratD2C.CratD2C)
		headerV1 := currentBlock.Header()
		headerV1.Extra = common.Hex2Bytes("d7830100018358444388676f312e31352e38856c696e757800000000000000000278c350152e15fa6ffc712a5a73d704ce73e2e103d9e17ae3ff2c6712e44e25b09ac5ee91f6c9ff065551f0dcac6f00cae11192d462db709be3758ccef312ee5eea8d7bad5374c6a652150515d744508b61c1a4deb4e4e7bf057e4e3824c11fd2569bcb77a52905cda63b5a58507910bed335e4c9d87ae0ecdfafd400")
		masternodesV1 := adaptor.GetMasternodesFromCheckpointHeader(headerV1)
		headerV2 := currentBlock.Header()
		headerV2.Number.Add(blockchain.Config().CratD2C.V2.SwitchBlock, big.NewInt(1))
		headerV2.Validators = []common.Address{common.HexToAddress("0278c350152e15fa6ffc712a5a73d704ce73e2e1"), common.HexToAddress("03d9e17ae3ff2c6712e44e25b09ac5ee91f6c9ff"), common.HexToAddress("065551f0dcac6f00cae11192d462db709be3758c")}
		headerV2.Extra = []byte{2}
		masternodesV2 := adaptor.GetMasternodesFromCheckpointHeader(headerV2)
		assert.True(t, reflect.DeepEqual(masternodesV1, masternodesV2), "GetMasternodesFromCheckpointHeader in adaptor for v1 v2 not equal", "v1", masternodesV1, "v2", masternodesV2)
	}
*/
func TestAdaptorIsEpochSwitch(t *testing.T) {
	blockchain, _, currentBlock, _, _, _ := PrepareXDCTestBlockChainForV2Engine(t, 1, params.TestXDPoSMockChainConfig, nil)
	adaptor := blockchain.Engine().(*CratD2C.CratD2C)
	header := currentBlock.Header()
	// v2
	parentBlockInfo := &types.BlockInfo{
		Hash:   header.ParentHash,
		Round:  types.Round(0),
		Number: big.NewInt(0).Set(blockchain.Config().CratD2C.V2.SwitchBlock),
	}
	quorumCert := &types.QuorumCert{
		ProposedBlockInfo: parentBlockInfo,
		Signatures:        nil,
		GapNumber:         0,
	}
	extra := types.ExtraFields_v2{
		Round:      1,
		QuorumCert: quorumCert,
	}
	extraBytes, err := extra.EncodeToBytes()
	assert.Nil(t, err)
	header.Extra = extraBytes
	header.Number.Set(big.NewInt(0))
	isEpochSwitchBlock, _, err := adaptor.IsEpochSwitch(header)
	assert.Nil(t, err)
	assert.True(t, isEpochSwitchBlock, "header at block number 0 should be epoch switch", header)

	header.Number.Set(big.NewInt(1))
	isEpochSwitchBlock, _, err = adaptor.IsEpochSwitch(header)
	assert.Nil(t, err)
	assert.False(t, isEpochSwitchBlock, "header at block number 1 shouldn't be epoch switch", header)

	header.Number.Set(big.NewInt(899))
	isEpochSwitchBlock, _, err = adaptor.IsEpochSwitch(header)
	assert.Nil(t, err)
	assert.False(t, isEpochSwitchBlock, "header at block number 899 shouldn't be epoch switch", header)

	header.Number.Set(big.NewInt(900))
	isEpochSwitchBlock, _, err = adaptor.IsEpochSwitch(header)
	assert.Nil(t, err)
	assert.True(t, isEpochSwitchBlock, "header at block number 900 should be epoch switch", header)

	header.Number.Set(big.NewInt(901))
	isEpochSwitchBlock, _, err = adaptor.IsEpochSwitch(header)
	assert.Nil(t, err)
	assert.False(t, isEpochSwitchBlock, "header at block number 901 shouldn't be epoch switch", header)

}

func TestAdaptorGetMasternodesV2(t *testing.T) {
	// we skip test for v1 since it's hard to make a real genesis block
	blockchain, _, currentBlock, signer, signFn, _ := PrepareXDCTestBlockChainForV2Engine(t, 900, params.TestXDPoSMockChainConfig, nil)
	adaptor := blockchain.Engine().(*CratD2C.CratD2C)
	blockNum := 901
	blockCoinBase := "0x111000000000000000000000000000000123"
	currentBlock = CreateBlock(blockchain, params.TestXDPoSMockChainConfig, currentBlock, blockNum, 1, blockCoinBase, signer, signFn, nil, nil, "")

	// Block 0 is the first v2 block, and is treated as epoch switch block
	err := blockchain.InsertBlock(currentBlock)
	assert.Nil(t, err)
	masternodes1 := adaptor.GetMasternodes(blockchain, currentBlock.Header())
	assert.Equal(t, 20, len(masternodes1))
	masternodes1ByNumber := adaptor.GetMasternodesByNumber(blockchain, currentBlock.NumberU64())
	assert.True(t, reflect.DeepEqual(masternodes1, masternodes1ByNumber), "at block number", blockNum)
	for blockNum = 902; blockNum < 915; blockNum++ {
		currentBlock = CreateBlock(blockchain, params.TestXDPoSMockChainConfig, currentBlock, blockNum, int64(blockNum-900), blockCoinBase, signer, signFn, nil, nil, "")
		err = blockchain.InsertBlock(currentBlock)
		assert.Nil(t, err)
		masternodes2 := adaptor.GetMasternodes(blockchain, currentBlock.Header())
		assert.True(t, reflect.DeepEqual(masternodes1, masternodes2), "at block number", blockNum)
		masternodes2ByNumber := adaptor.GetMasternodesByNumber(blockchain, currentBlock.NumberU64())
		assert.True(t, reflect.DeepEqual(masternodes2, masternodes2ByNumber), "at block number", blockNum)
	}
}

func TestGetCurrentEpochSwitchBlock(t *testing.T) {
	blockchain, _, currentBlock, signer, signFn, _ := PrepareXDCTestBlockChainForV2Engine(t, 900, params.TestXDPoSMockChainConfig, nil)
	adaptor := blockchain.Engine().(*CratD2C.CratD2C)

	currentCheckpointNumber, epochNum, err := adaptor.GetCurrentEpochSwitchBlock(blockchain, big.NewInt(900))
	assert.Nil(t, err)
	assert.Equal(t, uint64(900), currentCheckpointNumber)
	assert.Equal(t, uint64(1), epochNum)

	blockNum := 901
	blockCoinBase := "0x111000000000000000000000000000000123"
	currentBlock = CreateBlock(blockchain, params.TestXDPoSMockChainConfig, currentBlock, blockNum, 1, blockCoinBase, signer, signFn, nil, nil, "")
	err = blockchain.InsertBlock(currentBlock)
	assert.Nil(t, err)
	currentCheckpointNumber, epochNum, err = adaptor.GetCurrentEpochSwitchBlock(blockchain, currentBlock.Number())
	assert.Nil(t, err)
	assert.Equal(t, uint64(900), currentCheckpointNumber)
	assert.Equal(t, uint64(1), epochNum)

	for blockNum = 902; blockNum < 915; blockNum++ {
		currentBlock = CreateBlock(blockchain, params.TestXDPoSMockChainConfig, currentBlock, blockNum, int64(blockNum-900), blockCoinBase, signer, signFn, nil, nil, "")

		err = blockchain.InsertBlock(currentBlock)
		assert.Nil(t, err)
		currentCheckpointNumber, epochNum, err := adaptor.GetCurrentEpochSwitchBlock(blockchain, currentBlock.Number())
		assert.Nil(t, err)
		assert.Equal(t, uint64(900), currentCheckpointNumber)
		assert.Equal(t, uint64(1), epochNum)
	}
}

func TestGetParentBlock(t *testing.T) {
	blockchain, _, block900, signer, signFn, _ := PrepareXDCTestBlockChainForV2Engine(t, 900, params.TestXDPoSMockChainConfig, nil)
	adaptor := blockchain.Engine().(*CratD2C.CratD2C)
	block899 := blockchain.GetBlockByNumber(899) //highest QC

	// V2
	blockNum := 901
	blockCoinBase := "0x111000000000000000000000000000000123"
	block901 := CreateBlock(blockchain, params.TestXDPoSMockChainConfig, block900, blockNum, 1, blockCoinBase, signer, signFn, nil, nil, "")
	err := blockchain.InsertBlock(block901)
	assert.Nil(t, err)

	// let's inject another one, but the highestedQC has not been updated, so it shall still point to 900
	blockNum = 902
	block902 := CreateBlock(blockchain, params.TestXDPoSMockChainConfig, block901, blockNum, 1, blockCoinBase, signer, signFn, nil, nil, "")
	err = blockchain.InsertBlock(block902)
	assert.Nil(t, err)
	block := adaptor.FindParentBlockToAssign(blockchain, block902)

	assert.Equal(t, block899.Hash(), block.Hash())
}
