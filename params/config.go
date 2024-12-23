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

package params

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/CratD2C-SmartChain/cratd2cchain/common"
	"github.com/CratD2C-SmartChain/cratd2cchain/log"
)

const (
	ConsensusEngineVersion1 = "v1"
	ConsensusEngineVersion2 = "v2"
	Default                 = 0
)

var (
	XDCMainnetGenesisHash = common.HexToHash("866aade66493d5d09dbb9d40beeaa9ac8dce2dd7cc2019511ac727da139a92ef") // CRAT Mainnet genesis hash to enforce below configs on
	MainnetGenesisHash    = common.HexToHash("5fa43fd0cc2caa45b362c8c43c1f14df235f27a62135ca22f75fe17b47b3c5c3") // Mainnet genesis hash to enforce below configs on
	TestnetGenesisHash    = common.HexToHash("5fa43fd0cc2caa45b362c8c43c1f14df235f27a62135ca22f75fe17b47b3c5c3") // Testnet genesis hash to enforce below configs on
	DevnetGenesisHash     = common.HexToHash("ab6fd3cb7d1a489e03250c7d14c2d6d819a6a528d6380b31e8410951964ef423") // Devnet genesis hash to enforce below configs on
)

var (
	MainnetV2Configs = map[uint64]*V2Config{
		Default: {
			SwitchRound:          0,
			CertThreshold:        5, // To be confirmed once mainnet is ready
			TimeoutSyncThreshold: 3,
			TimeoutPeriod:        60,
			MinePeriod:           10,
		},
	}

	TestnetV2Configs = map[uint64]*V2Config{
		Default: {
			SwitchRound:          0,
			CertThreshold:        3,
			TimeoutSyncThreshold: 2,
			TimeoutPeriod:        4,
			MinePeriod:           2,
		},
	}

	DevnetV2Configs = map[uint64]*V2Config{
		Default: {
			SwitchRound:          0,
			CertThreshold:        2,
			TimeoutSyncThreshold: 5,
			TimeoutPeriod:        25,
			MinePeriod:           10,
		},
	}

	// CratD2CChain mainnet config
	CratD2CMainnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(50),
		HomesteadBlock: big.NewInt(1),
		EIP150Block:    big.NewInt(2),
		EIP150Hash:     common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:    big.NewInt(3),
		EIP158Block:    big.NewInt(3),
		ByzantiumBlock: big.NewInt(4),
		CratD2C: &CratD2CConfig{
			Period:              2,
			Epoch:               900,
			Reward:              5000,
			RewardCheckpoint:    900,
			Gap:                 450,
			FoudationWalletAddr: common.HexToAddress("xdc92a289fe95a85c53b8d0d113cbaef0c1ec98ac65"),
			V2: &V2{
				SwitchBlock:   common.TIPV2SwitchBlock,
				CurrentConfig: MainnetV2Configs[0],
				AllConfigs:    MainnetV2Configs,
			},
		},
	}

	UnitTestV2Configs = map[uint64]*V2Config{
		Default: {
			SwitchRound:          0,
			CertThreshold:        3,
			TimeoutSyncThreshold: 2,
			TimeoutPeriod:        10,
			MinePeriod:           2,
		},
		10: {
			SwitchRound:          10,
			CertThreshold:        5,
			TimeoutSyncThreshold: 2,
			TimeoutPeriod:        4,
			MinePeriod:           3,
		},
		899: {
			SwitchRound:          899,
			CertThreshold:        3,
			TimeoutSyncThreshold: 4,
			TimeoutPeriod:        5,
			MinePeriod:           2,
		},
		910: {
			SwitchRound:          915,
			CertThreshold:        5,
			TimeoutSyncThreshold: 4,
			TimeoutPeriod:        5,
			MinePeriod:           2,
		},
	}

	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainId:             big.NewInt(1),
		HomesteadBlock:      big.NewInt(1150000),
		DAOForkBlock:        big.NewInt(1920000),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2463000),
		EIP150Hash:          common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:         big.NewInt(2675000),
		EIP158Block:         big.NewInt(2675000),
		ByzantiumBlock:      big.NewInt(4370000),
		ConstantinopleBlock: nil,
		Ethash:              new(EthashConfig),
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainId:             big.NewInt(51),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: nil,
		CratD2C: &CratD2CConfig{
			Period:              2,
			Epoch:               900,
			Reward:              5000,
			RewardCheckpoint:    900,
			Gap:                 450,
			FoudationWalletAddr: common.HexToAddress("xdc746249c61f5832c5eed53172776b460491bdcd5c"),
			V2: &V2{
				SwitchBlock:      common.TIPV2SwitchBlock,
				CurrentConfig:    TestnetV2Configs[0],
				AllConfigs:       TestnetV2Configs,
				SkipV2Validation: true,
			},
		},
	}

	// DevnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	DevnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(551),
		HomesteadBlock: big.NewInt(1),
		EIP150Block:    big.NewInt(2),
		EIP150Hash:     common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:    big.NewInt(3),
		EIP158Block:    big.NewInt(3),
		ByzantiumBlock: big.NewInt(4),
		CratD2C: &CratD2CConfig{
			Period:              2,
			Epoch:               900,
			Reward:              5000,
			RewardCheckpoint:    900,
			Gap:                 450,
			FoudationWalletAddr: common.HexToAddress("0x746249c61f5832c5eed53172776b460491bdcd5c"),
			V2: &V2{
				SwitchBlock:   common.TIPV2SwitchBlock,
				CurrentConfig: DevnetV2Configs[0],
				AllConfigs:    DevnetV2Configs,
			},
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil}

	// AllXDPoSProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the CratD2C consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllXDPoSProtocolChanges  = &ChainConfig{big.NewInt(89), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, &CratD2CConfig{Period: 0, Epoch: 30000}}
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}, nil}

	// CratD2C config with v2 engine after block 0
	TestXDPoSMockChainConfig = &ChainConfig{
		big.NewInt(1337),
		big.NewInt(0),
		nil,
		false,
		big.NewInt(0),
		common.Hash{},
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		nil,
		new(EthashConfig),
		nil,
		&CratD2CConfig{
			Epoch:               900,
			Gap:                 450,
			SkipV1Validation:    true,
			FoudationWalletAddr: common.HexToAddress("0x0000000000000000000000000000000000000068"),
			Reward:              250,
			V2: &V2{
				SwitchBlock:      big.NewInt(0),
				CurrentConfig:    UnitTestV2Configs[0],
				AllConfigs:       UnitTestV2Configs,
				SkipV2Validation: true,
			},
			Denom:       "xdc",
			NetworkName: "xdc-subnet",
		},
	}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainId *big.Int `json:"chainId"` // Chain id identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash  *EthashConfig  `json:"ethash,omitempty"`
	Clique  *CliqueConfig  `json:"clique,omitempty"`
	CratD2C *CratD2CConfig `json:"CratD2C,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// CratD2CConfig is the consensus engine configs for delegated-proof-of-stake based sealing.
type CratD2CConfig struct {
	NetworkName         string         `json:"networkName"`         // The chain network name
	Denom               string         `json:"denom"`               // The denomination of the chain currency
	Period              uint64         `json:"period"`              // Number of seconds between blocks to enforce
	Epoch               uint64         `json:"epoch"`               // Epoch length to reset votes and checkpoint
	Reward              uint64         `json:"reward"`              // Block reward - unit Ether
	RewardCheckpoint    uint64         `json:"rewardCheckpoint"`    // Checkpoint block for calculate rewards.
	Gap                 uint64         `json:"gap"`                 // Gap time preparing for the next epoch
	FoudationWalletAddr common.Address `json:"foudationWalletAddr"` // Foundation Address Wallet
	SkipV1Validation    bool           //Skip Block Validation for testing purpose, V1 consensus only
	V2                  *V2            `json:"v2"`
}

type V2 struct {
	lock sync.RWMutex // Protects the signer fields

	SwitchBlock   *big.Int             `json:"switchBlock"`
	CurrentConfig *V2Config            `json:"config"`
	AllConfigs    map[uint64]*V2Config `json:"allConfigs"`
	configIndex   []uint64             //list of switch block of configs

	SkipV2Validation bool //Skip Block Validation for testing purpose, V2 consensus only
}

type V2Config struct {
	SwitchRound          uint64 `json:"switchRound"`          // v1 to v2 switch block number
	MinePeriod           int    `json:"minePeriod"`           // Miner mine period to mine a block
	TimeoutSyncThreshold int    `json:"timeoutSyncThreshold"` // send syncInfo after number of timeout
	TimeoutPeriod        int    `json:"timeoutPeriod"`        // Duration in ms
	CertThreshold        int    `json:"certificateThreshold"` // Necessary number of messages from master nodes to form a certificate
}

func (c *CratD2CConfig) String() string {
	return "CratD2C"
}

func (v *V2) UpdateConfig(round uint64) {
	v.lock.Lock()
	defer v.lock.Unlock()

	var index uint64

	//find the right config
	for i := range v.configIndex {
		if v.configIndex[i] <= round {
			index = v.configIndex[i]
			break
		}
	}
	// update to current config
	log.Info("[updateV2Config] Update config", "index", index, "round", round, "SwitchRound", v.AllConfigs[index].SwitchRound)
	v.CurrentConfig = v.AllConfigs[index]
}

func (v *V2) Config(round uint64) *V2Config {
	configRound := round - 1 //start from next block from SwitchRound number
	var index uint64

	//find the right config
	for i := range v.configIndex {
		if v.configIndex[i] <= configRound {
			index = v.configIndex[i]
			break
		}
	}
	return v.AllConfigs[index]
}

func (v *V2) BuildConfigIndex() {
	var list []uint64

	for i := range v.AllConfigs {
		list = append(list, i)
	}

	// sort, sort lib doesn't support type uint64, it's ok to have O(n^2)  because only few items in the list
	// Make it descending order
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	log.Info("[BuildConfigIndex] config list", "list", list)
	v.configIndex = list
}

func (v *V2) ConfigIndex() []uint64 {
	return v.configIndex
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.CratD2C != nil:
		engine = c.CratD2C
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Engine: %v}",
		c.ChainId,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAO returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(common.TIPXDCXCancellationFee, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(common.TIPXDCXCancellationFee, num)
}

// IsBerlin returns whether num is either equal to the Berlin fork block or greater.
func (c *ChainConfig) IsBerlin(num *big.Int) bool {
	return isForked(common.BerlinBlock, num)
}

// IsLondon returns whether num is either equal to the London fork block or greater.
func (c *ChainConfig) IsLondon(num *big.Int) bool {
	return isForked(common.LondonBlock, num)
}

// IsMerge returns whether num is either equal to the Merge fork block or greater.
// Different from Geth which uses `block.difficulty != nil`
func (c *ChainConfig) IsMerge(num *big.Int) bool {
	return isForked(common.MergeBlock, num)
}

// IsShanghai returns whether num is either equal to the Shanghai fork block or greater.
func (c *ChainConfig) IsShanghai(num *big.Int) bool {
	return isForked(common.ShanghaiBlock, num)
}

func (c *ChainConfig) IsTIP2019(num *big.Int) bool {
	return isForked(common.TIP2019Block, num)
}

func (c *ChainConfig) IsTIPSigning(num *big.Int) bool {
	return isForked(common.TIPSigning, num)
}

func (c *ChainConfig) IsTIPRandomize(num *big.Int) bool {
	return isForked(common.TIPRandomize, num)
}

func (c *ChainConfig) IsTIPNoHalvingMNReward(num *big.Int) bool {
	return isForked(common.TIPNoHalvingMNReward, num)
}
func (c *ChainConfig) IsTIPXDCX(num *big.Int) bool {
	if common.IsTestnet {
		return isForked(common.TIPXDCXTestnet, num)
	} else {
		return isForked(common.TIPXDCX, num)
	}
}

func (c *ChainConfig) IsTIPXDCXLending(num *big.Int) bool {
	return isForked(common.TIPXDCXLending, num)
}

func (c *ChainConfig) IsTIPXDCXCancellationFee(num *big.Int) bool {
	return isForked(common.TIPXDCXCancellationFee, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainId, newcfg.ChainId) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntatic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainId                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
	IsBerlin, IsLondon                                      bool
	IsMerge, IsShanghai                                     bool
}

func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainId := c.ChainId
	if chainId == nil {
		chainId = new(big.Int)
	}
	return Rules{
		ChainId:          new(big.Int).Set(chainId),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsIstanbul:       c.IsIstanbul(num),
		IsBerlin:         c.IsBerlin(num),
		IsLondon:         c.IsLondon(num),
		IsMerge:          c.IsMerge(num),
		IsShanghai:       c.IsShanghai(num),
	}
}
