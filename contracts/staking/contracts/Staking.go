// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CRATStakeManagerTestDelegatorPerValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type CRATStakeManagerTestDelegatorPerValidatorInfo struct {
	Amount             *big.Int
	StoredValidatorAcc *big.Int
	CalledForWithdraw  *big.Int
	LastClaim          *big.Int
	FixedReward        CRATStakeManagerTestFixedReward
	VariableReward     CRATStakeManagerTestVariableReward
}

// CRATStakeManagerTestFixedReward is an auto generated low-level Go binding around an user-defined struct.
type CRATStakeManagerTestFixedReward struct {
	Apr          *big.Int
	LastUpdate   *big.Int
	FixedReward  *big.Int
	TotalClaimed *big.Int
}

// CRATStakeManagerTestRoleSettings is an auto generated low-level Go binding around an user-defined struct.
type CRATStakeManagerTestRoleSettings struct {
	Apr              *big.Int
	ToSlash          *big.Int
	MinimumThreshold *big.Int
	ClaimCooldown    *big.Int
	WithdrawCooldown *big.Int
}

// CRATStakeManagerTestSlashPenaltyCalculation is an auto generated low-level Go binding around an user-defined struct.
type CRATStakeManagerTestSlashPenaltyCalculation struct {
	LastSlash        *big.Int
	PotentialPenalty *big.Int
}

// CRATStakeManagerTestValidatorInfoView is an auto generated low-level Go binding around an user-defined struct.
type CRATStakeManagerTestValidatorInfoView struct {
	Amount                 *big.Int
	Commission             *big.Int
	LastClaim              *big.Int
	CalledForWithdraw      *big.Int
	VestingEnd             *big.Int
	FixedReward            CRATStakeManagerTestFixedReward
	VariableReward         CRATStakeManagerTestVariableReward
	Penalty                CRATStakeManagerTestSlashPenaltyCalculation
	DelegatedAmount        *big.Int
	StoppedDelegatedAmount *big.Int
	DelegatorsAcc          *big.Int
	Delegators             []common.Address
	WithdrawAvailable      *big.Int
	ClaimAvailable         *big.Int
}

// CRATStakeManagerTestVariableReward is an auto generated low-level Go binding around an user-defined struct.
type CRATStakeManagerTestVariableReward struct {
	VariableReward *big.Int
	TotalClaimed   *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"DelegatorCalledForWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"DelegatorRevived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"DelegatorWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorCalledForWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"ValidatorDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRevived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorWithdrawed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWAP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YEAR_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"changeTestTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimAsDelegatorPerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAsValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegatorCallForWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegatorEarnedPerValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"validatorsArr\",\"type\":\"address[]\"}],\"name\":\"delegatorEarnedPerValidators\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"fixedRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"variableRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"depositAsDelegator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"depositAsValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingEnd\",\"type\":\"uint256\"}],\"name\":\"depositForValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forFixedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[3][]\",\"name\":\"amounts\",\"type\":\"uint256[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorsArr\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storedValidatorAcc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calledForWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaim\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.FixedReward\",\"name\":\"fixedReward\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.VariableReward\",\"name\":\"variableReward\",\"type\":\"tuple\"}],\"internalType\":\"structCRATStakeManagerTest.DelegatorPerValidatorInfo[]\",\"name\":\"delegatorPerValidatorArr\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"withdrawAvailable\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimAvailable\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getDelegatorsInfoPerValidator\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegators\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storedValidatorAcc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calledForWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaim\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.FixedReward\",\"name\":\"fixedReward\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.VariableReward\",\"name\":\"variableReward\",\"type\":\"tuple\"}],\"internalType\":\"structCRATStakeManagerTest.DelegatorPerValidatorInfo[]\",\"name\":\"delegatorPerValidatorArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStoppedValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[3][]\",\"name\":\"amounts\",\"type\":\"uint256[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calledForWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingEnd\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.FixedReward\",\"name\":\"fixedReward\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.VariableReward\",\"name\":\"variableReward\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lastSlash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"potentialPenalty\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.SlashPenaltyCalculation\",\"name\":\"penalty\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stoppedDelegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsAcc\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"delegators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimAvailable\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.ValidatorInfoView\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isDelegator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"restakeAsDelegator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restakeAsValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"reviveAsDelegator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reviveAsValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsAPR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsClaimCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsMinimum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsPercToSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsWithdrawCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setSlashReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValidatorsAPR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValidatorsAmountToSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValidatorsClaimCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValidatorsLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValidatorsMinimum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValidatorsWithdrawCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorsLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"slashReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toSlash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimCooldown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawCooldown\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.RoleSettings\",\"name\":\"validatorsSettings\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toSlash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimCooldown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawCooldown\",\"type\":\"uint256\"}],\"internalType\":\"structCRATStakeManagerTest.RoleSettings\",\"name\":\"delegatorsSettings\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stoppedDelegatorsPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stoppedValidatorsPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"totalDelegatorRewardPerValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelegatorsPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelegatorsRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"totalValidatorReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorsPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorsRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCallForWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawAsDelegator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAsValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawExcessFixedReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"delegators\",\"type\":\"address[]\"}],\"name\":\"withdrawForDelegators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawForValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Storage *StorageCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Storage *StorageSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Storage.Contract.DEFAULTADMINROLE(&_Storage.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Storage *StorageCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Storage.Contract.DEFAULTADMINROLE(&_Storage.CallOpts)
}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_Storage *StorageCaller) DISTRIBUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "DISTRIBUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_Storage *StorageSession) DISTRIBUTORROLE() ([32]byte, error) {
	return _Storage.Contract.DISTRIBUTORROLE(&_Storage.CallOpts)
}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_Storage *StorageCallerSession) DISTRIBUTORROLE() ([32]byte, error) {
	return _Storage.Contract.DISTRIBUTORROLE(&_Storage.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Storage *StorageCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Storage *StorageSession) PRECISION() (*big.Int, error) {
	return _Storage.Contract.PRECISION(&_Storage.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Storage *StorageCallerSession) PRECISION() (*big.Int, error) {
	return _Storage.Contract.PRECISION(&_Storage.CallOpts)
}

// SWAPROLE is a free data retrieval call binding the contract method 0xd4b7f403.
//
// Solidity: function SWAP_ROLE() view returns(bytes32)
func (_Storage *StorageCaller) SWAPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "SWAP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SWAPROLE is a free data retrieval call binding the contract method 0xd4b7f403.
//
// Solidity: function SWAP_ROLE() view returns(bytes32)
func (_Storage *StorageSession) SWAPROLE() ([32]byte, error) {
	return _Storage.Contract.SWAPROLE(&_Storage.CallOpts)
}

// SWAPROLE is a free data retrieval call binding the contract method 0xd4b7f403.
//
// Solidity: function SWAP_ROLE() view returns(bytes32)
func (_Storage *StorageCallerSession) SWAPROLE() ([32]byte, error) {
	return _Storage.Contract.SWAPROLE(&_Storage.CallOpts)
}

// YEARDURATION is a free data retrieval call binding the contract method 0x0dc0ba72.
//
// Solidity: function YEAR_DURATION() view returns(uint256)
func (_Storage *StorageCaller) YEARDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "YEAR_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YEARDURATION is a free data retrieval call binding the contract method 0x0dc0ba72.
//
// Solidity: function YEAR_DURATION() view returns(uint256)
func (_Storage *StorageSession) YEARDURATION() (*big.Int, error) {
	return _Storage.Contract.YEARDURATION(&_Storage.CallOpts)
}

// YEARDURATION is a free data retrieval call binding the contract method 0x0dc0ba72.
//
// Solidity: function YEAR_DURATION() view returns(uint256)
func (_Storage *StorageCallerSession) YEARDURATION() (*big.Int, error) {
	return _Storage.Contract.YEARDURATION(&_Storage.CallOpts)
}

// DelegatorEarnedPerValidator is a free data retrieval call binding the contract method 0x1cf208f8.
//
// Solidity: function delegatorEarnedPerValidator(address delegator, address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCaller) DelegatorEarnedPerValidator(opts *bind.CallOpts, delegator common.Address, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "delegatorEarnedPerValidator", delegator, validator)

	outstruct := new(struct {
		FixedReward    *big.Int
		VariableReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VariableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegatorEarnedPerValidator is a free data retrieval call binding the contract method 0x1cf208f8.
//
// Solidity: function delegatorEarnedPerValidator(address delegator, address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageSession) DelegatorEarnedPerValidator(delegator common.Address, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.DelegatorEarnedPerValidator(&_Storage.CallOpts, delegator, validator)
}

// DelegatorEarnedPerValidator is a free data retrieval call binding the contract method 0x1cf208f8.
//
// Solidity: function delegatorEarnedPerValidator(address delegator, address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCallerSession) DelegatorEarnedPerValidator(delegator common.Address, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.DelegatorEarnedPerValidator(&_Storage.CallOpts, delegator, validator)
}

// DelegatorEarnedPerValidators is a free data retrieval call binding the contract method 0x93cc48b4.
//
// Solidity: function delegatorEarnedPerValidators(address delegator, address[] validatorsArr) view returns(uint256[] fixedRewards, uint256[] variableRewards)
func (_Storage *StorageCaller) DelegatorEarnedPerValidators(opts *bind.CallOpts, delegator common.Address, validatorsArr []common.Address) (struct {
	FixedRewards    []*big.Int
	VariableRewards []*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "delegatorEarnedPerValidators", delegator, validatorsArr)

	outstruct := new(struct {
		FixedRewards    []*big.Int
		VariableRewards []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedRewards = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.VariableRewards = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// DelegatorEarnedPerValidators is a free data retrieval call binding the contract method 0x93cc48b4.
//
// Solidity: function delegatorEarnedPerValidators(address delegator, address[] validatorsArr) view returns(uint256[] fixedRewards, uint256[] variableRewards)
func (_Storage *StorageSession) DelegatorEarnedPerValidators(delegator common.Address, validatorsArr []common.Address) (struct {
	FixedRewards    []*big.Int
	VariableRewards []*big.Int
}, error) {
	return _Storage.Contract.DelegatorEarnedPerValidators(&_Storage.CallOpts, delegator, validatorsArr)
}

// DelegatorEarnedPerValidators is a free data retrieval call binding the contract method 0x93cc48b4.
//
// Solidity: function delegatorEarnedPerValidators(address delegator, address[] validatorsArr) view returns(uint256[] fixedRewards, uint256[] variableRewards)
func (_Storage *StorageCallerSession) DelegatorEarnedPerValidators(delegator common.Address, validatorsArr []common.Address) (struct {
	FixedRewards    []*big.Int
	VariableRewards []*big.Int
}, error) {
	return _Storage.Contract.DelegatorEarnedPerValidators(&_Storage.CallOpts, delegator, validatorsArr)
}

// ForFixedReward is a free data retrieval call binding the contract method 0x42273e92.
//
// Solidity: function forFixedReward() view returns(uint256)
func (_Storage *StorageCaller) ForFixedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "forFixedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForFixedReward is a free data retrieval call binding the contract method 0x42273e92.
//
// Solidity: function forFixedReward() view returns(uint256)
func (_Storage *StorageSession) ForFixedReward() (*big.Int, error) {
	return _Storage.Contract.ForFixedReward(&_Storage.CallOpts)
}

// ForFixedReward is a free data retrieval call binding the contract method 0x42273e92.
//
// Solidity: function forFixedReward() view returns(uint256)
func (_Storage *StorageCallerSession) ForFixedReward() (*big.Int, error) {
	return _Storage.Contract.ForFixedReward(&_Storage.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] validators, uint256[3][] amounts)
func (_Storage *StorageCaller) GetActiveValidators(opts *bind.CallOpts) (struct {
	Validators []common.Address
	Amounts    [][3]*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getActiveValidators")

	outstruct := new(struct {
		Validators []common.Address
		Amounts    [][3]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([][3]*big.Int)).(*[][3]*big.Int)

	return *outstruct, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] validators, uint256[3][] amounts)
func (_Storage *StorageSession) GetActiveValidators() (struct {
	Validators []common.Address
	Amounts    [][3]*big.Int
}, error) {
	return _Storage.Contract.GetActiveValidators(&_Storage.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] validators, uint256[3][] amounts)
func (_Storage *StorageCallerSession) GetActiveValidators() (struct {
	Validators []common.Address
	Amounts    [][3]*big.Int
}, error) {
	return _Storage.Contract.GetActiveValidators(&_Storage.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x18c3836b.
//
// Solidity: function getDelegatorInfo(address delegator) view returns(address[] validatorsArr, (uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256))[] delegatorPerValidatorArr, uint256[] withdrawAvailable, uint256[] claimAvailable)
func (_Storage *StorageCaller) GetDelegatorInfo(opts *bind.CallOpts, delegator common.Address) (struct {
	ValidatorsArr            []common.Address
	DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
	WithdrawAvailable        []*big.Int
	ClaimAvailable           []*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDelegatorInfo", delegator)

	outstruct := new(struct {
		ValidatorsArr            []common.Address
		DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
		WithdrawAvailable        []*big.Int
		ClaimAvailable           []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorsArr = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.DelegatorPerValidatorArr = *abi.ConvertType(out[1], new([]CRATStakeManagerTestDelegatorPerValidatorInfo)).(*[]CRATStakeManagerTestDelegatorPerValidatorInfo)
	outstruct.WithdrawAvailable = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.ClaimAvailable = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x18c3836b.
//
// Solidity: function getDelegatorInfo(address delegator) view returns(address[] validatorsArr, (uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256))[] delegatorPerValidatorArr, uint256[] withdrawAvailable, uint256[] claimAvailable)
func (_Storage *StorageSession) GetDelegatorInfo(delegator common.Address) (struct {
	ValidatorsArr            []common.Address
	DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
	WithdrawAvailable        []*big.Int
	ClaimAvailable           []*big.Int
}, error) {
	return _Storage.Contract.GetDelegatorInfo(&_Storage.CallOpts, delegator)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x18c3836b.
//
// Solidity: function getDelegatorInfo(address delegator) view returns(address[] validatorsArr, (uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256))[] delegatorPerValidatorArr, uint256[] withdrawAvailable, uint256[] claimAvailable)
func (_Storage *StorageCallerSession) GetDelegatorInfo(delegator common.Address) (struct {
	ValidatorsArr            []common.Address
	DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
	WithdrawAvailable        []*big.Int
	ClaimAvailable           []*big.Int
}, error) {
	return _Storage.Contract.GetDelegatorInfo(&_Storage.CallOpts, delegator)
}

// GetDelegatorsInfoPerValidator is a free data retrieval call binding the contract method 0xb2fdb177.
//
// Solidity: function getDelegatorsInfoPerValidator(address validator) view returns(address[] delegators, (uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256))[] delegatorPerValidatorArr)
func (_Storage *StorageCaller) GetDelegatorsInfoPerValidator(opts *bind.CallOpts, validator common.Address) (struct {
	Delegators               []common.Address
	DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDelegatorsInfoPerValidator", validator)

	outstruct := new(struct {
		Delegators               []common.Address
		DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.DelegatorPerValidatorArr = *abi.ConvertType(out[1], new([]CRATStakeManagerTestDelegatorPerValidatorInfo)).(*[]CRATStakeManagerTestDelegatorPerValidatorInfo)

	return *outstruct, err

}

// GetDelegatorsInfoPerValidator is a free data retrieval call binding the contract method 0xb2fdb177.
//
// Solidity: function getDelegatorsInfoPerValidator(address validator) view returns(address[] delegators, (uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256))[] delegatorPerValidatorArr)
func (_Storage *StorageSession) GetDelegatorsInfoPerValidator(validator common.Address) (struct {
	Delegators               []common.Address
	DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
}, error) {
	return _Storage.Contract.GetDelegatorsInfoPerValidator(&_Storage.CallOpts, validator)
}

// GetDelegatorsInfoPerValidator is a free data retrieval call binding the contract method 0xb2fdb177.
//
// Solidity: function getDelegatorsInfoPerValidator(address validator) view returns(address[] delegators, (uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256))[] delegatorPerValidatorArr)
func (_Storage *StorageCallerSession) GetDelegatorsInfoPerValidator(validator common.Address) (struct {
	Delegators               []common.Address
	DelegatorPerValidatorArr []CRATStakeManagerTestDelegatorPerValidatorInfo
}, error) {
	return _Storage.Contract.GetDelegatorsInfoPerValidator(&_Storage.CallOpts, validator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Storage *StorageCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Storage *StorageSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetRoleAdmin(&_Storage.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Storage *StorageCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetRoleAdmin(&_Storage.CallOpts, role)
}

// GetStoppedValidators is a free data retrieval call binding the contract method 0xb9c5ae38.
//
// Solidity: function getStoppedValidators() view returns(address[] validators, uint256[3][] amounts)
func (_Storage *StorageCaller) GetStoppedValidators(opts *bind.CallOpts) (struct {
	Validators []common.Address
	Amounts    [][3]*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getStoppedValidators")

	outstruct := new(struct {
		Validators []common.Address
		Amounts    [][3]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([][3]*big.Int)).(*[][3]*big.Int)

	return *outstruct, err

}

// GetStoppedValidators is a free data retrieval call binding the contract method 0xb9c5ae38.
//
// Solidity: function getStoppedValidators() view returns(address[] validators, uint256[3][] amounts)
func (_Storage *StorageSession) GetStoppedValidators() (struct {
	Validators []common.Address
	Amounts    [][3]*big.Int
}, error) {
	return _Storage.Contract.GetStoppedValidators(&_Storage.CallOpts)
}

// GetStoppedValidators is a free data retrieval call binding the contract method 0xb9c5ae38.
//
// Solidity: function getStoppedValidators() view returns(address[] validators, uint256[3][] amounts)
func (_Storage *StorageCallerSession) GetStoppedValidators() (struct {
	Validators []common.Address
	Amounts    [][3]*big.Int
}, error) {
	return _Storage.Contract.GetStoppedValidators(&_Storage.CallOpts)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns((uint256,uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,address[],uint256,uint256) info)
func (_Storage *StorageCaller) GetValidatorInfo(opts *bind.CallOpts, validator common.Address) (CRATStakeManagerTestValidatorInfoView, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getValidatorInfo", validator)

	if err != nil {
		return *new(CRATStakeManagerTestValidatorInfoView), err
	}

	out0 := *abi.ConvertType(out[0], new(CRATStakeManagerTestValidatorInfoView)).(*CRATStakeManagerTestValidatorInfoView)

	return out0, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns((uint256,uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,address[],uint256,uint256) info)
func (_Storage *StorageSession) GetValidatorInfo(validator common.Address) (CRATStakeManagerTestValidatorInfoView, error) {
	return _Storage.Contract.GetValidatorInfo(&_Storage.CallOpts, validator)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns((uint256,uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,address[],uint256,uint256) info)
func (_Storage *StorageCallerSession) GetValidatorInfo(validator common.Address) (CRATStakeManagerTestValidatorInfoView, error) {
	return _Storage.Contract.GetValidatorInfo(&_Storage.CallOpts, validator)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Storage *StorageCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Storage *StorageSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Storage.Contract.HasRole(&_Storage.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Storage *StorageCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Storage.Contract.HasRole(&_Storage.CallOpts, role, account)
}

// IsDelegator is a free data retrieval call binding the contract method 0xfd8ab482.
//
// Solidity: function isDelegator(address account) view returns(bool)
func (_Storage *StorageCaller) IsDelegator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isDelegator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegator is a free data retrieval call binding the contract method 0xfd8ab482.
//
// Solidity: function isDelegator(address account) view returns(bool)
func (_Storage *StorageSession) IsDelegator(account common.Address) (bool, error) {
	return _Storage.Contract.IsDelegator(&_Storage.CallOpts, account)
}

// IsDelegator is a free data retrieval call binding the contract method 0xfd8ab482.
//
// Solidity: function isDelegator(address account) view returns(bool)
func (_Storage *StorageCallerSession) IsDelegator(account common.Address) (bool, error) {
	return _Storage.Contract.IsDelegator(&_Storage.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) view returns(bool)
func (_Storage *StorageCaller) IsValidator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isValidator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) view returns(bool)
func (_Storage *StorageSession) IsValidator(account common.Address) (bool, error) {
	return _Storage.Contract.IsValidator(&_Storage.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) view returns(bool)
func (_Storage *StorageCallerSession) IsValidator(account common.Address) (bool, error) {
	return _Storage.Contract.IsValidator(&_Storage.CallOpts, account)
}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(uint256 validatorsLimit, address slashReceiver, (uint256,uint256,uint256,uint256,uint256) validatorsSettings, (uint256,uint256,uint256,uint256,uint256) delegatorsSettings)
func (_Storage *StorageCaller) Settings(opts *bind.CallOpts) (struct {
	ValidatorsLimit    *big.Int
	SlashReceiver      common.Address
	ValidatorsSettings CRATStakeManagerTestRoleSettings
	DelegatorsSettings CRATStakeManagerTestRoleSettings
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "settings")

	outstruct := new(struct {
		ValidatorsLimit    *big.Int
		SlashReceiver      common.Address
		ValidatorsSettings CRATStakeManagerTestRoleSettings
		DelegatorsSettings CRATStakeManagerTestRoleSettings
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorsLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SlashReceiver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ValidatorsSettings = *abi.ConvertType(out[2], new(CRATStakeManagerTestRoleSettings)).(*CRATStakeManagerTestRoleSettings)
	outstruct.DelegatorsSettings = *abi.ConvertType(out[3], new(CRATStakeManagerTestRoleSettings)).(*CRATStakeManagerTestRoleSettings)

	return *outstruct, err

}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(uint256 validatorsLimit, address slashReceiver, (uint256,uint256,uint256,uint256,uint256) validatorsSettings, (uint256,uint256,uint256,uint256,uint256) delegatorsSettings)
func (_Storage *StorageSession) Settings() (struct {
	ValidatorsLimit    *big.Int
	SlashReceiver      common.Address
	ValidatorsSettings CRATStakeManagerTestRoleSettings
	DelegatorsSettings CRATStakeManagerTestRoleSettings
}, error) {
	return _Storage.Contract.Settings(&_Storage.CallOpts)
}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(uint256 validatorsLimit, address slashReceiver, (uint256,uint256,uint256,uint256,uint256) validatorsSettings, (uint256,uint256,uint256,uint256,uint256) delegatorsSettings)
func (_Storage *StorageCallerSession) Settings() (struct {
	ValidatorsLimit    *big.Int
	SlashReceiver      common.Address
	ValidatorsSettings CRATStakeManagerTestRoleSettings
	DelegatorsSettings CRATStakeManagerTestRoleSettings
}, error) {
	return _Storage.Contract.Settings(&_Storage.CallOpts)
}

// StoppedDelegatorsPool is a free data retrieval call binding the contract method 0x89da31ea.
//
// Solidity: function stoppedDelegatorsPool() view returns(uint256)
func (_Storage *StorageCaller) StoppedDelegatorsPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "stoppedDelegatorsPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoppedDelegatorsPool is a free data retrieval call binding the contract method 0x89da31ea.
//
// Solidity: function stoppedDelegatorsPool() view returns(uint256)
func (_Storage *StorageSession) StoppedDelegatorsPool() (*big.Int, error) {
	return _Storage.Contract.StoppedDelegatorsPool(&_Storage.CallOpts)
}

// StoppedDelegatorsPool is a free data retrieval call binding the contract method 0x89da31ea.
//
// Solidity: function stoppedDelegatorsPool() view returns(uint256)
func (_Storage *StorageCallerSession) StoppedDelegatorsPool() (*big.Int, error) {
	return _Storage.Contract.StoppedDelegatorsPool(&_Storage.CallOpts)
}

// StoppedValidatorsPool is a free data retrieval call binding the contract method 0x011ceb69.
//
// Solidity: function stoppedValidatorsPool() view returns(uint256)
func (_Storage *StorageCaller) StoppedValidatorsPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "stoppedValidatorsPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoppedValidatorsPool is a free data retrieval call binding the contract method 0x011ceb69.
//
// Solidity: function stoppedValidatorsPool() view returns(uint256)
func (_Storage *StorageSession) StoppedValidatorsPool() (*big.Int, error) {
	return _Storage.Contract.StoppedValidatorsPool(&_Storage.CallOpts)
}

// StoppedValidatorsPool is a free data retrieval call binding the contract method 0x011ceb69.
//
// Solidity: function stoppedValidatorsPool() view returns(uint256)
func (_Storage *StorageCallerSession) StoppedValidatorsPool() (*big.Int, error) {
	return _Storage.Contract.StoppedValidatorsPool(&_Storage.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Storage *StorageCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Storage *StorageSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Storage.Contract.SupportsInterface(&_Storage.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Storage *StorageCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Storage.Contract.SupportsInterface(&_Storage.CallOpts, interfaceId)
}

// TestTime is a free data retrieval call binding the contract method 0x3c35a0c1.
//
// Solidity: function testTime() view returns(uint256)
func (_Storage *StorageCaller) TestTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "testTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestTime is a free data retrieval call binding the contract method 0x3c35a0c1.
//
// Solidity: function testTime() view returns(uint256)
func (_Storage *StorageSession) TestTime() (*big.Int, error) {
	return _Storage.Contract.TestTime(&_Storage.CallOpts)
}

// TestTime is a free data retrieval call binding the contract method 0x3c35a0c1.
//
// Solidity: function testTime() view returns(uint256)
func (_Storage *StorageCallerSession) TestTime() (*big.Int, error) {
	return _Storage.Contract.TestTime(&_Storage.CallOpts)
}

// TotalDelegatorRewardPerValidator is a free data retrieval call binding the contract method 0x8ce3e7c6.
//
// Solidity: function totalDelegatorRewardPerValidator(address delegator, address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCaller) TotalDelegatorRewardPerValidator(opts *bind.CallOpts, delegator common.Address, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalDelegatorRewardPerValidator", delegator, validator)

	outstruct := new(struct {
		FixedReward    *big.Int
		VariableReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VariableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalDelegatorRewardPerValidator is a free data retrieval call binding the contract method 0x8ce3e7c6.
//
// Solidity: function totalDelegatorRewardPerValidator(address delegator, address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageSession) TotalDelegatorRewardPerValidator(delegator common.Address, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalDelegatorRewardPerValidator(&_Storage.CallOpts, delegator, validator)
}

// TotalDelegatorRewardPerValidator is a free data retrieval call binding the contract method 0x8ce3e7c6.
//
// Solidity: function totalDelegatorRewardPerValidator(address delegator, address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCallerSession) TotalDelegatorRewardPerValidator(delegator common.Address, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalDelegatorRewardPerValidator(&_Storage.CallOpts, delegator, validator)
}

// TotalDelegatorsPool is a free data retrieval call binding the contract method 0x8c60d40c.
//
// Solidity: function totalDelegatorsPool() view returns(uint256)
func (_Storage *StorageCaller) TotalDelegatorsPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalDelegatorsPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelegatorsPool is a free data retrieval call binding the contract method 0x8c60d40c.
//
// Solidity: function totalDelegatorsPool() view returns(uint256)
func (_Storage *StorageSession) TotalDelegatorsPool() (*big.Int, error) {
	return _Storage.Contract.TotalDelegatorsPool(&_Storage.CallOpts)
}

// TotalDelegatorsPool is a free data retrieval call binding the contract method 0x8c60d40c.
//
// Solidity: function totalDelegatorsPool() view returns(uint256)
func (_Storage *StorageCallerSession) TotalDelegatorsPool() (*big.Int, error) {
	return _Storage.Contract.TotalDelegatorsPool(&_Storage.CallOpts)
}

// TotalDelegatorsRewards is a free data retrieval call binding the contract method 0x55e2ff03.
//
// Solidity: function totalDelegatorsRewards() view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCaller) TotalDelegatorsRewards(opts *bind.CallOpts) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalDelegatorsRewards")

	outstruct := new(struct {
		FixedReward    *big.Int
		VariableReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VariableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalDelegatorsRewards is a free data retrieval call binding the contract method 0x55e2ff03.
//
// Solidity: function totalDelegatorsRewards() view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageSession) TotalDelegatorsRewards() (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalDelegatorsRewards(&_Storage.CallOpts)
}

// TotalDelegatorsRewards is a free data retrieval call binding the contract method 0x55e2ff03.
//
// Solidity: function totalDelegatorsRewards() view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCallerSession) TotalDelegatorsRewards() (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalDelegatorsRewards(&_Storage.CallOpts)
}

// TotalValidatorReward is a free data retrieval call binding the contract method 0x302df125.
//
// Solidity: function totalValidatorReward(address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCaller) TotalValidatorReward(opts *bind.CallOpts, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalValidatorReward", validator)

	outstruct := new(struct {
		FixedReward    *big.Int
		VariableReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VariableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalValidatorReward is a free data retrieval call binding the contract method 0x302df125.
//
// Solidity: function totalValidatorReward(address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageSession) TotalValidatorReward(validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalValidatorReward(&_Storage.CallOpts, validator)
}

// TotalValidatorReward is a free data retrieval call binding the contract method 0x302df125.
//
// Solidity: function totalValidatorReward(address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCallerSession) TotalValidatorReward(validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalValidatorReward(&_Storage.CallOpts, validator)
}

// TotalValidatorsPool is a free data retrieval call binding the contract method 0x88b41e06.
//
// Solidity: function totalValidatorsPool() view returns(uint256)
func (_Storage *StorageCaller) TotalValidatorsPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalValidatorsPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValidatorsPool is a free data retrieval call binding the contract method 0x88b41e06.
//
// Solidity: function totalValidatorsPool() view returns(uint256)
func (_Storage *StorageSession) TotalValidatorsPool() (*big.Int, error) {
	return _Storage.Contract.TotalValidatorsPool(&_Storage.CallOpts)
}

// TotalValidatorsPool is a free data retrieval call binding the contract method 0x88b41e06.
//
// Solidity: function totalValidatorsPool() view returns(uint256)
func (_Storage *StorageCallerSession) TotalValidatorsPool() (*big.Int, error) {
	return _Storage.Contract.TotalValidatorsPool(&_Storage.CallOpts)
}

// TotalValidatorsRewards is a free data retrieval call binding the contract method 0x3b11858c.
//
// Solidity: function totalValidatorsRewards() view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCaller) TotalValidatorsRewards(opts *bind.CallOpts) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalValidatorsRewards")

	outstruct := new(struct {
		FixedReward    *big.Int
		VariableReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VariableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalValidatorsRewards is a free data retrieval call binding the contract method 0x3b11858c.
//
// Solidity: function totalValidatorsRewards() view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageSession) TotalValidatorsRewards() (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalValidatorsRewards(&_Storage.CallOpts)
}

// TotalValidatorsRewards is a free data retrieval call binding the contract method 0x3b11858c.
//
// Solidity: function totalValidatorsRewards() view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCallerSession) TotalValidatorsRewards() (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.TotalValidatorsRewards(&_Storage.CallOpts)
}

// ValidatorEarned is a free data retrieval call binding the contract method 0xd91f1c85.
//
// Solidity: function validatorEarned(address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCaller) ValidatorEarned(opts *bind.CallOpts, validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "validatorEarned", validator)

	outstruct := new(struct {
		FixedReward    *big.Int
		VariableReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VariableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorEarned is a free data retrieval call binding the contract method 0xd91f1c85.
//
// Solidity: function validatorEarned(address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageSession) ValidatorEarned(validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.ValidatorEarned(&_Storage.CallOpts, validator)
}

// ValidatorEarned is a free data retrieval call binding the contract method 0xd91f1c85.
//
// Solidity: function validatorEarned(address validator) view returns(uint256 fixedReward, uint256 variableReward)
func (_Storage *StorageCallerSession) ValidatorEarned(validator common.Address) (struct {
	FixedReward    *big.Int
	VariableReward *big.Int
}, error) {
	return _Storage.Contract.ValidatorEarned(&_Storage.CallOpts, validator)
}

// ChangeTestTime is a paid mutator transaction binding the contract method 0x80387b18.
//
// Solidity: function changeTestTime(uint256 value) returns()
func (_Storage *StorageTransactor) ChangeTestTime(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeTestTime", value)
}

// ChangeTestTime is a paid mutator transaction binding the contract method 0x80387b18.
//
// Solidity: function changeTestTime(uint256 value) returns()
func (_Storage *StorageSession) ChangeTestTime(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeTestTime(&_Storage.TransactOpts, value)
}

// ChangeTestTime is a paid mutator transaction binding the contract method 0x80387b18.
//
// Solidity: function changeTestTime(uint256 value) returns()
func (_Storage *StorageTransactorSession) ChangeTestTime(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeTestTime(&_Storage.TransactOpts, value)
}

// ClaimAsDelegatorPerValidator is a paid mutator transaction binding the contract method 0xd0019eb3.
//
// Solidity: function claimAsDelegatorPerValidator(address validator) returns()
func (_Storage *StorageTransactor) ClaimAsDelegatorPerValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "claimAsDelegatorPerValidator", validator)
}

// ClaimAsDelegatorPerValidator is a paid mutator transaction binding the contract method 0xd0019eb3.
//
// Solidity: function claimAsDelegatorPerValidator(address validator) returns()
func (_Storage *StorageSession) ClaimAsDelegatorPerValidator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ClaimAsDelegatorPerValidator(&_Storage.TransactOpts, validator)
}

// ClaimAsDelegatorPerValidator is a paid mutator transaction binding the contract method 0xd0019eb3.
//
// Solidity: function claimAsDelegatorPerValidator(address validator) returns()
func (_Storage *StorageTransactorSession) ClaimAsDelegatorPerValidator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ClaimAsDelegatorPerValidator(&_Storage.TransactOpts, validator)
}

// ClaimAsValidator is a paid mutator transaction binding the contract method 0xbd843124.
//
// Solidity: function claimAsValidator() returns()
func (_Storage *StorageTransactor) ClaimAsValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "claimAsValidator")
}

// ClaimAsValidator is a paid mutator transaction binding the contract method 0xbd843124.
//
// Solidity: function claimAsValidator() returns()
func (_Storage *StorageSession) ClaimAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.ClaimAsValidator(&_Storage.TransactOpts)
}

// ClaimAsValidator is a paid mutator transaction binding the contract method 0xbd843124.
//
// Solidity: function claimAsValidator() returns()
func (_Storage *StorageTransactorSession) ClaimAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.ClaimAsValidator(&_Storage.TransactOpts)
}

// DelegatorCallForWithdraw is a paid mutator transaction binding the contract method 0x871566ad.
//
// Solidity: function delegatorCallForWithdraw(address validator) returns()
func (_Storage *StorageTransactor) DelegatorCallForWithdraw(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "delegatorCallForWithdraw", validator)
}

// DelegatorCallForWithdraw is a paid mutator transaction binding the contract method 0x871566ad.
//
// Solidity: function delegatorCallForWithdraw(address validator) returns()
func (_Storage *StorageSession) DelegatorCallForWithdraw(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.DelegatorCallForWithdraw(&_Storage.TransactOpts, validator)
}

// DelegatorCallForWithdraw is a paid mutator transaction binding the contract method 0x871566ad.
//
// Solidity: function delegatorCallForWithdraw(address validator) returns()
func (_Storage *StorageTransactorSession) DelegatorCallForWithdraw(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.DelegatorCallForWithdraw(&_Storage.TransactOpts, validator)
}

// DepositAsDelegator is a paid mutator transaction binding the contract method 0x9cc4f420.
//
// Solidity: function depositAsDelegator(address validator) payable returns()
func (_Storage *StorageTransactor) DepositAsDelegator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "depositAsDelegator", validator)
}

// DepositAsDelegator is a paid mutator transaction binding the contract method 0x9cc4f420.
//
// Solidity: function depositAsDelegator(address validator) payable returns()
func (_Storage *StorageSession) DepositAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.DepositAsDelegator(&_Storage.TransactOpts, validator)
}

// DepositAsDelegator is a paid mutator transaction binding the contract method 0x9cc4f420.
//
// Solidity: function depositAsDelegator(address validator) payable returns()
func (_Storage *StorageTransactorSession) DepositAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.DepositAsDelegator(&_Storage.TransactOpts, validator)
}

// DepositAsValidator is a paid mutator transaction binding the contract method 0x9cb56b66.
//
// Solidity: function depositAsValidator(uint256 commission) payable returns()
func (_Storage *StorageTransactor) DepositAsValidator(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "depositAsValidator", commission)
}

// DepositAsValidator is a paid mutator transaction binding the contract method 0x9cb56b66.
//
// Solidity: function depositAsValidator(uint256 commission) payable returns()
func (_Storage *StorageSession) DepositAsValidator(commission *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DepositAsValidator(&_Storage.TransactOpts, commission)
}

// DepositAsValidator is a paid mutator transaction binding the contract method 0x9cb56b66.
//
// Solidity: function depositAsValidator(uint256 commission) payable returns()
func (_Storage *StorageTransactorSession) DepositAsValidator(commission *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DepositAsValidator(&_Storage.TransactOpts, commission)
}

// DepositForValidator is a paid mutator transaction binding the contract method 0xda01b469.
//
// Solidity: function depositForValidator(address sender, uint256 commission, uint256 vestingEnd) payable returns()
func (_Storage *StorageTransactor) DepositForValidator(opts *bind.TransactOpts, sender common.Address, commission *big.Int, vestingEnd *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "depositForValidator", sender, commission, vestingEnd)
}

// DepositForValidator is a paid mutator transaction binding the contract method 0xda01b469.
//
// Solidity: function depositForValidator(address sender, uint256 commission, uint256 vestingEnd) payable returns()
func (_Storage *StorageSession) DepositForValidator(sender common.Address, commission *big.Int, vestingEnd *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DepositForValidator(&_Storage.TransactOpts, sender, commission, vestingEnd)
}

// DepositForValidator is a paid mutator transaction binding the contract method 0xda01b469.
//
// Solidity: function depositForValidator(address sender, uint256 commission, uint256 vestingEnd) payable returns()
func (_Storage *StorageTransactorSession) DepositForValidator(sender common.Address, commission *big.Int, vestingEnd *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DepositForValidator(&_Storage.TransactOpts, sender, commission, vestingEnd)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] validators, uint256[] amounts) payable returns()
func (_Storage *StorageTransactor) DistributeRewards(opts *bind.TransactOpts, validators []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "distributeRewards", validators, amounts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] validators, uint256[] amounts) payable returns()
func (_Storage *StorageSession) DistributeRewards(validators []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DistributeRewards(&_Storage.TransactOpts, validators, amounts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x143ba4f3.
//
// Solidity: function distributeRewards(address[] validators, uint256[] amounts) payable returns()
func (_Storage *StorageTransactorSession) DistributeRewards(validators []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DistributeRewards(&_Storage.TransactOpts, validators, amounts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Storage *StorageTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Storage *StorageSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Storage.Contract.GrantRole(&_Storage.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Storage *StorageTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Storage.Contract.GrantRole(&_Storage.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _distributor, address _receiver) returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts, _distributor common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize", _distributor, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _distributor, address _receiver) returns()
func (_Storage *StorageSession) Initialize(_distributor common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, _distributor, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _distributor, address _receiver) returns()
func (_Storage *StorageTransactorSession) Initialize(_distributor common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, _distributor, _receiver)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Storage *StorageTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Storage *StorageSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RenounceRole(&_Storage.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Storage *StorageTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RenounceRole(&_Storage.TransactOpts, role, callerConfirmation)
}

// RestakeAsDelegator is a paid mutator transaction binding the contract method 0xd250217a.
//
// Solidity: function restakeAsDelegator(address validator) returns()
func (_Storage *StorageTransactor) RestakeAsDelegator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "restakeAsDelegator", validator)
}

// RestakeAsDelegator is a paid mutator transaction binding the contract method 0xd250217a.
//
// Solidity: function restakeAsDelegator(address validator) returns()
func (_Storage *StorageSession) RestakeAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RestakeAsDelegator(&_Storage.TransactOpts, validator)
}

// RestakeAsDelegator is a paid mutator transaction binding the contract method 0xd250217a.
//
// Solidity: function restakeAsDelegator(address validator) returns()
func (_Storage *StorageTransactorSession) RestakeAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RestakeAsDelegator(&_Storage.TransactOpts, validator)
}

// RestakeAsValidator is a paid mutator transaction binding the contract method 0xe317fd45.
//
// Solidity: function restakeAsValidator() returns()
func (_Storage *StorageTransactor) RestakeAsValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "restakeAsValidator")
}

// RestakeAsValidator is a paid mutator transaction binding the contract method 0xe317fd45.
//
// Solidity: function restakeAsValidator() returns()
func (_Storage *StorageSession) RestakeAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.RestakeAsValidator(&_Storage.TransactOpts)
}

// RestakeAsValidator is a paid mutator transaction binding the contract method 0xe317fd45.
//
// Solidity: function restakeAsValidator() returns()
func (_Storage *StorageTransactorSession) RestakeAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.RestakeAsValidator(&_Storage.TransactOpts)
}

// ReviveAsDelegator is a paid mutator transaction binding the contract method 0xa9f9ec5c.
//
// Solidity: function reviveAsDelegator(address validator) payable returns()
func (_Storage *StorageTransactor) ReviveAsDelegator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "reviveAsDelegator", validator)
}

// ReviveAsDelegator is a paid mutator transaction binding the contract method 0xa9f9ec5c.
//
// Solidity: function reviveAsDelegator(address validator) payable returns()
func (_Storage *StorageSession) ReviveAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ReviveAsDelegator(&_Storage.TransactOpts, validator)
}

// ReviveAsDelegator is a paid mutator transaction binding the contract method 0xa9f9ec5c.
//
// Solidity: function reviveAsDelegator(address validator) payable returns()
func (_Storage *StorageTransactorSession) ReviveAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ReviveAsDelegator(&_Storage.TransactOpts, validator)
}

// ReviveAsValidator is a paid mutator transaction binding the contract method 0x4f426dee.
//
// Solidity: function reviveAsValidator() payable returns()
func (_Storage *StorageTransactor) ReviveAsValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "reviveAsValidator")
}

// ReviveAsValidator is a paid mutator transaction binding the contract method 0x4f426dee.
//
// Solidity: function reviveAsValidator() payable returns()
func (_Storage *StorageSession) ReviveAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.ReviveAsValidator(&_Storage.TransactOpts)
}

// ReviveAsValidator is a paid mutator transaction binding the contract method 0x4f426dee.
//
// Solidity: function reviveAsValidator() payable returns()
func (_Storage *StorageTransactorSession) ReviveAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.ReviveAsValidator(&_Storage.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Storage *StorageTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Storage *StorageSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RevokeRole(&_Storage.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Storage *StorageTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RevokeRole(&_Storage.TransactOpts, role, account)
}

// SetDelegatorsAPR is a paid mutator transaction binding the contract method 0x0f42c389.
//
// Solidity: function setDelegatorsAPR(uint256 value) returns()
func (_Storage *StorageTransactor) SetDelegatorsAPR(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDelegatorsAPR", value)
}

// SetDelegatorsAPR is a paid mutator transaction binding the contract method 0x0f42c389.
//
// Solidity: function setDelegatorsAPR(uint256 value) returns()
func (_Storage *StorageSession) SetDelegatorsAPR(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsAPR(&_Storage.TransactOpts, value)
}

// SetDelegatorsAPR is a paid mutator transaction binding the contract method 0x0f42c389.
//
// Solidity: function setDelegatorsAPR(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetDelegatorsAPR(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsAPR(&_Storage.TransactOpts, value)
}

// SetDelegatorsClaimCooldown is a paid mutator transaction binding the contract method 0x3898d08c.
//
// Solidity: function setDelegatorsClaimCooldown(uint256 value) returns()
func (_Storage *StorageTransactor) SetDelegatorsClaimCooldown(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDelegatorsClaimCooldown", value)
}

// SetDelegatorsClaimCooldown is a paid mutator transaction binding the contract method 0x3898d08c.
//
// Solidity: function setDelegatorsClaimCooldown(uint256 value) returns()
func (_Storage *StorageSession) SetDelegatorsClaimCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsClaimCooldown(&_Storage.TransactOpts, value)
}

// SetDelegatorsClaimCooldown is a paid mutator transaction binding the contract method 0x3898d08c.
//
// Solidity: function setDelegatorsClaimCooldown(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetDelegatorsClaimCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsClaimCooldown(&_Storage.TransactOpts, value)
}

// SetDelegatorsMinimum is a paid mutator transaction binding the contract method 0xc3ca7e39.
//
// Solidity: function setDelegatorsMinimum(uint256 value) returns()
func (_Storage *StorageTransactor) SetDelegatorsMinimum(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDelegatorsMinimum", value)
}

// SetDelegatorsMinimum is a paid mutator transaction binding the contract method 0xc3ca7e39.
//
// Solidity: function setDelegatorsMinimum(uint256 value) returns()
func (_Storage *StorageSession) SetDelegatorsMinimum(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsMinimum(&_Storage.TransactOpts, value)
}

// SetDelegatorsMinimum is a paid mutator transaction binding the contract method 0xc3ca7e39.
//
// Solidity: function setDelegatorsMinimum(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetDelegatorsMinimum(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsMinimum(&_Storage.TransactOpts, value)
}

// SetDelegatorsPercToSlash is a paid mutator transaction binding the contract method 0xfc9d506c.
//
// Solidity: function setDelegatorsPercToSlash(uint256 value) returns()
func (_Storage *StorageTransactor) SetDelegatorsPercToSlash(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDelegatorsPercToSlash", value)
}

// SetDelegatorsPercToSlash is a paid mutator transaction binding the contract method 0xfc9d506c.
//
// Solidity: function setDelegatorsPercToSlash(uint256 value) returns()
func (_Storage *StorageSession) SetDelegatorsPercToSlash(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsPercToSlash(&_Storage.TransactOpts, value)
}

// SetDelegatorsPercToSlash is a paid mutator transaction binding the contract method 0xfc9d506c.
//
// Solidity: function setDelegatorsPercToSlash(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetDelegatorsPercToSlash(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsPercToSlash(&_Storage.TransactOpts, value)
}

// SetDelegatorsWithdrawCooldown is a paid mutator transaction binding the contract method 0x88caa107.
//
// Solidity: function setDelegatorsWithdrawCooldown(uint256 value) returns()
func (_Storage *StorageTransactor) SetDelegatorsWithdrawCooldown(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDelegatorsWithdrawCooldown", value)
}

// SetDelegatorsWithdrawCooldown is a paid mutator transaction binding the contract method 0x88caa107.
//
// Solidity: function setDelegatorsWithdrawCooldown(uint256 value) returns()
func (_Storage *StorageSession) SetDelegatorsWithdrawCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsWithdrawCooldown(&_Storage.TransactOpts, value)
}

// SetDelegatorsWithdrawCooldown is a paid mutator transaction binding the contract method 0x88caa107.
//
// Solidity: function setDelegatorsWithdrawCooldown(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetDelegatorsWithdrawCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDelegatorsWithdrawCooldown(&_Storage.TransactOpts, value)
}

// SetSlashReceiver is a paid mutator transaction binding the contract method 0x1a6933d5.
//
// Solidity: function setSlashReceiver(address receiver) returns()
func (_Storage *StorageTransactor) SetSlashReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setSlashReceiver", receiver)
}

// SetSlashReceiver is a paid mutator transaction binding the contract method 0x1a6933d5.
//
// Solidity: function setSlashReceiver(address receiver) returns()
func (_Storage *StorageSession) SetSlashReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetSlashReceiver(&_Storage.TransactOpts, receiver)
}

// SetSlashReceiver is a paid mutator transaction binding the contract method 0x1a6933d5.
//
// Solidity: function setSlashReceiver(address receiver) returns()
func (_Storage *StorageTransactorSession) SetSlashReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetSlashReceiver(&_Storage.TransactOpts, receiver)
}

// SetValidatorsAPR is a paid mutator transaction binding the contract method 0x0bd8999a.
//
// Solidity: function setValidatorsAPR(uint256 value) returns()
func (_Storage *StorageTransactor) SetValidatorsAPR(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setValidatorsAPR", value)
}

// SetValidatorsAPR is a paid mutator transaction binding the contract method 0x0bd8999a.
//
// Solidity: function setValidatorsAPR(uint256 value) returns()
func (_Storage *StorageSession) SetValidatorsAPR(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsAPR(&_Storage.TransactOpts, value)
}

// SetValidatorsAPR is a paid mutator transaction binding the contract method 0x0bd8999a.
//
// Solidity: function setValidatorsAPR(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetValidatorsAPR(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsAPR(&_Storage.TransactOpts, value)
}

// SetValidatorsAmountToSlash is a paid mutator transaction binding the contract method 0xd5efdf5c.
//
// Solidity: function setValidatorsAmountToSlash(uint256 value) returns()
func (_Storage *StorageTransactor) SetValidatorsAmountToSlash(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setValidatorsAmountToSlash", value)
}

// SetValidatorsAmountToSlash is a paid mutator transaction binding the contract method 0xd5efdf5c.
//
// Solidity: function setValidatorsAmountToSlash(uint256 value) returns()
func (_Storage *StorageSession) SetValidatorsAmountToSlash(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsAmountToSlash(&_Storage.TransactOpts, value)
}

// SetValidatorsAmountToSlash is a paid mutator transaction binding the contract method 0xd5efdf5c.
//
// Solidity: function setValidatorsAmountToSlash(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetValidatorsAmountToSlash(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsAmountToSlash(&_Storage.TransactOpts, value)
}

// SetValidatorsClaimCooldown is a paid mutator transaction binding the contract method 0x209d3881.
//
// Solidity: function setValidatorsClaimCooldown(uint256 value) returns()
func (_Storage *StorageTransactor) SetValidatorsClaimCooldown(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setValidatorsClaimCooldown", value)
}

// SetValidatorsClaimCooldown is a paid mutator transaction binding the contract method 0x209d3881.
//
// Solidity: function setValidatorsClaimCooldown(uint256 value) returns()
func (_Storage *StorageSession) SetValidatorsClaimCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsClaimCooldown(&_Storage.TransactOpts, value)
}

// SetValidatorsClaimCooldown is a paid mutator transaction binding the contract method 0x209d3881.
//
// Solidity: function setValidatorsClaimCooldown(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetValidatorsClaimCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsClaimCooldown(&_Storage.TransactOpts, value)
}

// SetValidatorsLimit is a paid mutator transaction binding the contract method 0xf76bd541.
//
// Solidity: function setValidatorsLimit(uint256 value) returns()
func (_Storage *StorageTransactor) SetValidatorsLimit(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setValidatorsLimit", value)
}

// SetValidatorsLimit is a paid mutator transaction binding the contract method 0xf76bd541.
//
// Solidity: function setValidatorsLimit(uint256 value) returns()
func (_Storage *StorageSession) SetValidatorsLimit(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsLimit(&_Storage.TransactOpts, value)
}

// SetValidatorsLimit is a paid mutator transaction binding the contract method 0xf76bd541.
//
// Solidity: function setValidatorsLimit(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetValidatorsLimit(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsLimit(&_Storage.TransactOpts, value)
}

// SetValidatorsMinimum is a paid mutator transaction binding the contract method 0xc01b96ad.
//
// Solidity: function setValidatorsMinimum(uint256 value) returns()
func (_Storage *StorageTransactor) SetValidatorsMinimum(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setValidatorsMinimum", value)
}

// SetValidatorsMinimum is a paid mutator transaction binding the contract method 0xc01b96ad.
//
// Solidity: function setValidatorsMinimum(uint256 value) returns()
func (_Storage *StorageSession) SetValidatorsMinimum(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsMinimum(&_Storage.TransactOpts, value)
}

// SetValidatorsMinimum is a paid mutator transaction binding the contract method 0xc01b96ad.
//
// Solidity: function setValidatorsMinimum(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetValidatorsMinimum(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsMinimum(&_Storage.TransactOpts, value)
}

// SetValidatorsWithdrawCooldown is a paid mutator transaction binding the contract method 0x1bc4ffdd.
//
// Solidity: function setValidatorsWithdrawCooldown(uint256 value) returns()
func (_Storage *StorageTransactor) SetValidatorsWithdrawCooldown(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setValidatorsWithdrawCooldown", value)
}

// SetValidatorsWithdrawCooldown is a paid mutator transaction binding the contract method 0x1bc4ffdd.
//
// Solidity: function setValidatorsWithdrawCooldown(uint256 value) returns()
func (_Storage *StorageSession) SetValidatorsWithdrawCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsWithdrawCooldown(&_Storage.TransactOpts, value)
}

// SetValidatorsWithdrawCooldown is a paid mutator transaction binding the contract method 0x1bc4ffdd.
//
// Solidity: function setValidatorsWithdrawCooldown(uint256 value) returns()
func (_Storage *StorageTransactorSession) SetValidatorsWithdrawCooldown(value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetValidatorsWithdrawCooldown(&_Storage.TransactOpts, value)
}

// Slash is a paid mutator transaction binding the contract method 0x8b8c24c1.
//
// Solidity: function slash(address[] validators) returns()
func (_Storage *StorageTransactor) Slash(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "slash", validators)
}

// Slash is a paid mutator transaction binding the contract method 0x8b8c24c1.
//
// Solidity: function slash(address[] validators) returns()
func (_Storage *StorageSession) Slash(validators []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Slash(&_Storage.TransactOpts, validators)
}

// Slash is a paid mutator transaction binding the contract method 0x8b8c24c1.
//
// Solidity: function slash(address[] validators) returns()
func (_Storage *StorageTransactorSession) Slash(validators []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Slash(&_Storage.TransactOpts, validators)
}

// ValidatorCallForWithdraw is a paid mutator transaction binding the contract method 0x50778ba3.
//
// Solidity: function validatorCallForWithdraw() returns()
func (_Storage *StorageTransactor) ValidatorCallForWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "validatorCallForWithdraw")
}

// ValidatorCallForWithdraw is a paid mutator transaction binding the contract method 0x50778ba3.
//
// Solidity: function validatorCallForWithdraw() returns()
func (_Storage *StorageSession) ValidatorCallForWithdraw() (*types.Transaction, error) {
	return _Storage.Contract.ValidatorCallForWithdraw(&_Storage.TransactOpts)
}

// ValidatorCallForWithdraw is a paid mutator transaction binding the contract method 0x50778ba3.
//
// Solidity: function validatorCallForWithdraw() returns()
func (_Storage *StorageTransactorSession) ValidatorCallForWithdraw() (*types.Transaction, error) {
	return _Storage.Contract.ValidatorCallForWithdraw(&_Storage.TransactOpts)
}

// WithdrawAsDelegator is a paid mutator transaction binding the contract method 0x373944b5.
//
// Solidity: function withdrawAsDelegator(address validator) returns()
func (_Storage *StorageTransactor) WithdrawAsDelegator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawAsDelegator", validator)
}

// WithdrawAsDelegator is a paid mutator transaction binding the contract method 0x373944b5.
//
// Solidity: function withdrawAsDelegator(address validator) returns()
func (_Storage *StorageSession) WithdrawAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawAsDelegator(&_Storage.TransactOpts, validator)
}

// WithdrawAsDelegator is a paid mutator transaction binding the contract method 0x373944b5.
//
// Solidity: function withdrawAsDelegator(address validator) returns()
func (_Storage *StorageTransactorSession) WithdrawAsDelegator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawAsDelegator(&_Storage.TransactOpts, validator)
}

// WithdrawAsValidator is a paid mutator transaction binding the contract method 0xe8a6de2b.
//
// Solidity: function withdrawAsValidator() returns()
func (_Storage *StorageTransactor) WithdrawAsValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawAsValidator")
}

// WithdrawAsValidator is a paid mutator transaction binding the contract method 0xe8a6de2b.
//
// Solidity: function withdrawAsValidator() returns()
func (_Storage *StorageSession) WithdrawAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.WithdrawAsValidator(&_Storage.TransactOpts)
}

// WithdrawAsValidator is a paid mutator transaction binding the contract method 0xe8a6de2b.
//
// Solidity: function withdrawAsValidator() returns()
func (_Storage *StorageTransactorSession) WithdrawAsValidator() (*types.Transaction, error) {
	return _Storage.Contract.WithdrawAsValidator(&_Storage.TransactOpts)
}

// WithdrawExcessFixedReward is a paid mutator transaction binding the contract method 0xef672fa1.
//
// Solidity: function withdrawExcessFixedReward(uint256 amount) returns()
func (_Storage *StorageTransactor) WithdrawExcessFixedReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawExcessFixedReward", amount)
}

// WithdrawExcessFixedReward is a paid mutator transaction binding the contract method 0xef672fa1.
//
// Solidity: function withdrawExcessFixedReward(uint256 amount) returns()
func (_Storage *StorageSession) WithdrawExcessFixedReward(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawExcessFixedReward(&_Storage.TransactOpts, amount)
}

// WithdrawExcessFixedReward is a paid mutator transaction binding the contract method 0xef672fa1.
//
// Solidity: function withdrawExcessFixedReward(uint256 amount) returns()
func (_Storage *StorageTransactorSession) WithdrawExcessFixedReward(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawExcessFixedReward(&_Storage.TransactOpts, amount)
}

// WithdrawForDelegators is a paid mutator transaction binding the contract method 0xaf31b0e9.
//
// Solidity: function withdrawForDelegators(address validator, address[] delegators) returns()
func (_Storage *StorageTransactor) WithdrawForDelegators(opts *bind.TransactOpts, validator common.Address, delegators []common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawForDelegators", validator, delegators)
}

// WithdrawForDelegators is a paid mutator transaction binding the contract method 0xaf31b0e9.
//
// Solidity: function withdrawForDelegators(address validator, address[] delegators) returns()
func (_Storage *StorageSession) WithdrawForDelegators(validator common.Address, delegators []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawForDelegators(&_Storage.TransactOpts, validator, delegators)
}

// WithdrawForDelegators is a paid mutator transaction binding the contract method 0xaf31b0e9.
//
// Solidity: function withdrawForDelegators(address validator, address[] delegators) returns()
func (_Storage *StorageTransactorSession) WithdrawForDelegators(validator common.Address, delegators []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawForDelegators(&_Storage.TransactOpts, validator, delegators)
}

// WithdrawForValidator is a paid mutator transaction binding the contract method 0x2f35c691.
//
// Solidity: function withdrawForValidator(address validator) returns()
func (_Storage *StorageTransactor) WithdrawForValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawForValidator", validator)
}

// WithdrawForValidator is a paid mutator transaction binding the contract method 0x2f35c691.
//
// Solidity: function withdrawForValidator(address validator) returns()
func (_Storage *StorageSession) WithdrawForValidator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawForValidator(&_Storage.TransactOpts, validator)
}

// WithdrawForValidator is a paid mutator transaction binding the contract method 0x2f35c691.
//
// Solidity: function withdrawForValidator(address validator) returns()
func (_Storage *StorageTransactorSession) WithdrawForValidator(validator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawForValidator(&_Storage.TransactOpts, validator)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactorSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// StorageDelegatorCalledForWithdrawIterator is returned from FilterDelegatorCalledForWithdraw and is used to iterate over the raw logs and unpacked data for DelegatorCalledForWithdraw events raised by the Storage contract.
type StorageDelegatorCalledForWithdrawIterator struct {
	Event *StorageDelegatorCalledForWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDelegatorCalledForWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDelegatorCalledForWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDelegatorCalledForWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDelegatorCalledForWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDelegatorCalledForWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDelegatorCalledForWithdraw represents a DelegatorCalledForWithdraw event raised by the Storage contract.
type StorageDelegatorCalledForWithdraw struct {
	Delegator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorCalledForWithdraw is a free log retrieval operation binding the contract event 0xac1632fe30069d2cf3bdfba44fbafd937138e87816f8d618027d918246a27553.
//
// Solidity: event DelegatorCalledForWithdraw(address delegator)
func (_Storage *StorageFilterer) FilterDelegatorCalledForWithdraw(opts *bind.FilterOpts) (*StorageDelegatorCalledForWithdrawIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DelegatorCalledForWithdraw")
	if err != nil {
		return nil, err
	}
	return &StorageDelegatorCalledForWithdrawIterator{contract: _Storage.contract, event: "DelegatorCalledForWithdraw", logs: logs, sub: sub}, nil
}

// WatchDelegatorCalledForWithdraw is a free log subscription operation binding the contract event 0xac1632fe30069d2cf3bdfba44fbafd937138e87816f8d618027d918246a27553.
//
// Solidity: event DelegatorCalledForWithdraw(address delegator)
func (_Storage *StorageFilterer) WatchDelegatorCalledForWithdraw(opts *bind.WatchOpts, sink chan<- *StorageDelegatorCalledForWithdraw) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DelegatorCalledForWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDelegatorCalledForWithdraw)
				if err := _Storage.contract.UnpackLog(event, "DelegatorCalledForWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorCalledForWithdraw is a log parse operation binding the contract event 0xac1632fe30069d2cf3bdfba44fbafd937138e87816f8d618027d918246a27553.
//
// Solidity: event DelegatorCalledForWithdraw(address delegator)
func (_Storage *StorageFilterer) ParseDelegatorCalledForWithdraw(log types.Log) (*StorageDelegatorCalledForWithdraw, error) {
	event := new(StorageDelegatorCalledForWithdraw)
	if err := _Storage.contract.UnpackLog(event, "DelegatorCalledForWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDelegatorClaimedIterator is returned from FilterDelegatorClaimed and is used to iterate over the raw logs and unpacked data for DelegatorClaimed events raised by the Storage contract.
type StorageDelegatorClaimedIterator struct {
	Event *StorageDelegatorClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDelegatorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDelegatorClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDelegatorClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDelegatorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDelegatorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDelegatorClaimed represents a DelegatorClaimed event raised by the Storage contract.
type StorageDelegatorClaimed struct {
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorClaimed is a free log retrieval operation binding the contract event 0x2cc5bb4b67c11ca2dd990733e7d4a46832ca8d875de3765b3262d70db9b52d81.
//
// Solidity: event DelegatorClaimed(address delegator, uint256 amount)
func (_Storage *StorageFilterer) FilterDelegatorClaimed(opts *bind.FilterOpts) (*StorageDelegatorClaimedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DelegatorClaimed")
	if err != nil {
		return nil, err
	}
	return &StorageDelegatorClaimedIterator{contract: _Storage.contract, event: "DelegatorClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorClaimed is a free log subscription operation binding the contract event 0x2cc5bb4b67c11ca2dd990733e7d4a46832ca8d875de3765b3262d70db9b52d81.
//
// Solidity: event DelegatorClaimed(address delegator, uint256 amount)
func (_Storage *StorageFilterer) WatchDelegatorClaimed(opts *bind.WatchOpts, sink chan<- *StorageDelegatorClaimed) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DelegatorClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDelegatorClaimed)
				if err := _Storage.contract.UnpackLog(event, "DelegatorClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorClaimed is a log parse operation binding the contract event 0x2cc5bb4b67c11ca2dd990733e7d4a46832ca8d875de3765b3262d70db9b52d81.
//
// Solidity: event DelegatorClaimed(address delegator, uint256 amount)
func (_Storage *StorageFilterer) ParseDelegatorClaimed(log types.Log) (*StorageDelegatorClaimed, error) {
	event := new(StorageDelegatorClaimed)
	if err := _Storage.contract.UnpackLog(event, "DelegatorClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDelegatorDepositedIterator is returned from FilterDelegatorDeposited and is used to iterate over the raw logs and unpacked data for DelegatorDeposited events raised by the Storage contract.
type StorageDelegatorDepositedIterator struct {
	Event *StorageDelegatorDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDelegatorDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDelegatorDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDelegatorDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDelegatorDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDelegatorDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDelegatorDeposited represents a DelegatorDeposited event raised by the Storage contract.
type StorageDelegatorDeposited struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorDeposited is a free log retrieval operation binding the contract event 0x26c220cc833a77ab0dcc3447efd56fa098f32095b1bc8144d12b12fead9bb9a4.
//
// Solidity: event DelegatorDeposited(address delegator, address validator, uint256 amount)
func (_Storage *StorageFilterer) FilterDelegatorDeposited(opts *bind.FilterOpts) (*StorageDelegatorDepositedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DelegatorDeposited")
	if err != nil {
		return nil, err
	}
	return &StorageDelegatorDepositedIterator{contract: _Storage.contract, event: "DelegatorDeposited", logs: logs, sub: sub}, nil
}

// WatchDelegatorDeposited is a free log subscription operation binding the contract event 0x26c220cc833a77ab0dcc3447efd56fa098f32095b1bc8144d12b12fead9bb9a4.
//
// Solidity: event DelegatorDeposited(address delegator, address validator, uint256 amount)
func (_Storage *StorageFilterer) WatchDelegatorDeposited(opts *bind.WatchOpts, sink chan<- *StorageDelegatorDeposited) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DelegatorDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDelegatorDeposited)
				if err := _Storage.contract.UnpackLog(event, "DelegatorDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorDeposited is a log parse operation binding the contract event 0x26c220cc833a77ab0dcc3447efd56fa098f32095b1bc8144d12b12fead9bb9a4.
//
// Solidity: event DelegatorDeposited(address delegator, address validator, uint256 amount)
func (_Storage *StorageFilterer) ParseDelegatorDeposited(log types.Log) (*StorageDelegatorDeposited, error) {
	event := new(StorageDelegatorDeposited)
	if err := _Storage.contract.UnpackLog(event, "DelegatorDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDelegatorRevivedIterator is returned from FilterDelegatorRevived and is used to iterate over the raw logs and unpacked data for DelegatorRevived events raised by the Storage contract.
type StorageDelegatorRevivedIterator struct {
	Event *StorageDelegatorRevived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDelegatorRevivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDelegatorRevived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDelegatorRevived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDelegatorRevivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDelegatorRevivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDelegatorRevived represents a DelegatorRevived event raised by the Storage contract.
type StorageDelegatorRevived struct {
	Delegator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRevived is a free log retrieval operation binding the contract event 0x040076653a9e6786d4eab46346be729e6ee17e8070a2637953772f4704f84c73.
//
// Solidity: event DelegatorRevived(address delegator)
func (_Storage *StorageFilterer) FilterDelegatorRevived(opts *bind.FilterOpts) (*StorageDelegatorRevivedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DelegatorRevived")
	if err != nil {
		return nil, err
	}
	return &StorageDelegatorRevivedIterator{contract: _Storage.contract, event: "DelegatorRevived", logs: logs, sub: sub}, nil
}

// WatchDelegatorRevived is a free log subscription operation binding the contract event 0x040076653a9e6786d4eab46346be729e6ee17e8070a2637953772f4704f84c73.
//
// Solidity: event DelegatorRevived(address delegator)
func (_Storage *StorageFilterer) WatchDelegatorRevived(opts *bind.WatchOpts, sink chan<- *StorageDelegatorRevived) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DelegatorRevived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDelegatorRevived)
				if err := _Storage.contract.UnpackLog(event, "DelegatorRevived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRevived is a log parse operation binding the contract event 0x040076653a9e6786d4eab46346be729e6ee17e8070a2637953772f4704f84c73.
//
// Solidity: event DelegatorRevived(address delegator)
func (_Storage *StorageFilterer) ParseDelegatorRevived(log types.Log) (*StorageDelegatorRevived, error) {
	event := new(StorageDelegatorRevived)
	if err := _Storage.contract.UnpackLog(event, "DelegatorRevived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDelegatorWithdrawedIterator is returned from FilterDelegatorWithdrawed and is used to iterate over the raw logs and unpacked data for DelegatorWithdrawed events raised by the Storage contract.
type StorageDelegatorWithdrawedIterator struct {
	Event *StorageDelegatorWithdrawed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDelegatorWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDelegatorWithdrawed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDelegatorWithdrawed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDelegatorWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDelegatorWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDelegatorWithdrawed represents a DelegatorWithdrawed event raised by the Storage contract.
type StorageDelegatorWithdrawed struct {
	Delegator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorWithdrawed is a free log retrieval operation binding the contract event 0x4d2482436185a9dc70f48332b654c5190763276c13afe701d1dfb7306af4861e.
//
// Solidity: event DelegatorWithdrawed(address delegator)
func (_Storage *StorageFilterer) FilterDelegatorWithdrawed(opts *bind.FilterOpts) (*StorageDelegatorWithdrawedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DelegatorWithdrawed")
	if err != nil {
		return nil, err
	}
	return &StorageDelegatorWithdrawedIterator{contract: _Storage.contract, event: "DelegatorWithdrawed", logs: logs, sub: sub}, nil
}

// WatchDelegatorWithdrawed is a free log subscription operation binding the contract event 0x4d2482436185a9dc70f48332b654c5190763276c13afe701d1dfb7306af4861e.
//
// Solidity: event DelegatorWithdrawed(address delegator)
func (_Storage *StorageFilterer) WatchDelegatorWithdrawed(opts *bind.WatchOpts, sink chan<- *StorageDelegatorWithdrawed) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DelegatorWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDelegatorWithdrawed)
				if err := _Storage.contract.UnpackLog(event, "DelegatorWithdrawed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorWithdrawed is a log parse operation binding the contract event 0x4d2482436185a9dc70f48332b654c5190763276c13afe701d1dfb7306af4861e.
//
// Solidity: event DelegatorWithdrawed(address delegator)
func (_Storage *StorageFilterer) ParseDelegatorWithdrawed(log types.Log) (*StorageDelegatorWithdrawed, error) {
	event := new(StorageDelegatorWithdrawed)
	if err := _Storage.contract.UnpackLog(event, "DelegatorWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Storage contract.
type StorageRoleAdminChangedIterator struct {
	Event *StorageRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRoleAdminChanged represents a RoleAdminChanged event raised by the Storage contract.
type StorageRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Storage *StorageFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StorageRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StorageRoleAdminChangedIterator{contract: _Storage.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Storage *StorageFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRoleAdminChanged)
				if err := _Storage.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Storage *StorageFilterer) ParseRoleAdminChanged(log types.Log) (*StorageRoleAdminChanged, error) {
	event := new(StorageRoleAdminChanged)
	if err := _Storage.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Storage contract.
type StorageRoleGrantedIterator struct {
	Event *StorageRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRoleGranted represents a RoleGranted event raised by the Storage contract.
type StorageRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Storage *StorageFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StorageRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StorageRoleGrantedIterator{contract: _Storage.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Storage *StorageFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StorageRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRoleGranted)
				if err := _Storage.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Storage *StorageFilterer) ParseRoleGranted(log types.Log) (*StorageRoleGranted, error) {
	event := new(StorageRoleGranted)
	if err := _Storage.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Storage contract.
type StorageRoleRevokedIterator struct {
	Event *StorageRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRoleRevoked represents a RoleRevoked event raised by the Storage contract.
type StorageRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Storage *StorageFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StorageRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StorageRoleRevokedIterator{contract: _Storage.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Storage *StorageFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StorageRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRoleRevoked)
				if err := _Storage.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Storage *StorageFilterer) ParseRoleRevoked(log types.Log) (*StorageRoleRevoked, error) {
	event := new(StorageRoleRevoked)
	if err := _Storage.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageValidatorCalledForWithdrawIterator is returned from FilterValidatorCalledForWithdraw and is used to iterate over the raw logs and unpacked data for ValidatorCalledForWithdraw events raised by the Storage contract.
type StorageValidatorCalledForWithdrawIterator struct {
	Event *StorageValidatorCalledForWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageValidatorCalledForWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageValidatorCalledForWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageValidatorCalledForWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageValidatorCalledForWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageValidatorCalledForWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageValidatorCalledForWithdraw represents a ValidatorCalledForWithdraw event raised by the Storage contract.
type StorageValidatorCalledForWithdraw struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorCalledForWithdraw is a free log retrieval operation binding the contract event 0xa9ef4c6a79539b205149d607c389b83c2e9c16ddfa216be5cae7e35a040124a3.
//
// Solidity: event ValidatorCalledForWithdraw(address validator)
func (_Storage *StorageFilterer) FilterValidatorCalledForWithdraw(opts *bind.FilterOpts) (*StorageValidatorCalledForWithdrawIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ValidatorCalledForWithdraw")
	if err != nil {
		return nil, err
	}
	return &StorageValidatorCalledForWithdrawIterator{contract: _Storage.contract, event: "ValidatorCalledForWithdraw", logs: logs, sub: sub}, nil
}

// WatchValidatorCalledForWithdraw is a free log subscription operation binding the contract event 0xa9ef4c6a79539b205149d607c389b83c2e9c16ddfa216be5cae7e35a040124a3.
//
// Solidity: event ValidatorCalledForWithdraw(address validator)
func (_Storage *StorageFilterer) WatchValidatorCalledForWithdraw(opts *bind.WatchOpts, sink chan<- *StorageValidatorCalledForWithdraw) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ValidatorCalledForWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageValidatorCalledForWithdraw)
				if err := _Storage.contract.UnpackLog(event, "ValidatorCalledForWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorCalledForWithdraw is a log parse operation binding the contract event 0xa9ef4c6a79539b205149d607c389b83c2e9c16ddfa216be5cae7e35a040124a3.
//
// Solidity: event ValidatorCalledForWithdraw(address validator)
func (_Storage *StorageFilterer) ParseValidatorCalledForWithdraw(log types.Log) (*StorageValidatorCalledForWithdraw, error) {
	event := new(StorageValidatorCalledForWithdraw)
	if err := _Storage.contract.UnpackLog(event, "ValidatorCalledForWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageValidatorClaimedIterator is returned from FilterValidatorClaimed and is used to iterate over the raw logs and unpacked data for ValidatorClaimed events raised by the Storage contract.
type StorageValidatorClaimedIterator struct {
	Event *StorageValidatorClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageValidatorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageValidatorClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageValidatorClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageValidatorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageValidatorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageValidatorClaimed represents a ValidatorClaimed event raised by the Storage contract.
type StorageValidatorClaimed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorClaimed is a free log retrieval operation binding the contract event 0xc20af58baaae9e9347897f7e714ad24bf44b9bc415c3ad4bc843cc5f06e0c82a.
//
// Solidity: event ValidatorClaimed(address validator, uint256 amount)
func (_Storage *StorageFilterer) FilterValidatorClaimed(opts *bind.FilterOpts) (*StorageValidatorClaimedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ValidatorClaimed")
	if err != nil {
		return nil, err
	}
	return &StorageValidatorClaimedIterator{contract: _Storage.contract, event: "ValidatorClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorClaimed is a free log subscription operation binding the contract event 0xc20af58baaae9e9347897f7e714ad24bf44b9bc415c3ad4bc843cc5f06e0c82a.
//
// Solidity: event ValidatorClaimed(address validator, uint256 amount)
func (_Storage *StorageFilterer) WatchValidatorClaimed(opts *bind.WatchOpts, sink chan<- *StorageValidatorClaimed) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ValidatorClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageValidatorClaimed)
				if err := _Storage.contract.UnpackLog(event, "ValidatorClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorClaimed is a log parse operation binding the contract event 0xc20af58baaae9e9347897f7e714ad24bf44b9bc415c3ad4bc843cc5f06e0c82a.
//
// Solidity: event ValidatorClaimed(address validator, uint256 amount)
func (_Storage *StorageFilterer) ParseValidatorClaimed(log types.Log) (*StorageValidatorClaimed, error) {
	event := new(StorageValidatorClaimed)
	if err := _Storage.contract.UnpackLog(event, "ValidatorClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageValidatorDepositedIterator is returned from FilterValidatorDeposited and is used to iterate over the raw logs and unpacked data for ValidatorDeposited events raised by the Storage contract.
type StorageValidatorDepositedIterator struct {
	Event *StorageValidatorDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageValidatorDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageValidatorDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageValidatorDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageValidatorDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageValidatorDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageValidatorDeposited represents a ValidatorDeposited event raised by the Storage contract.
type StorageValidatorDeposited struct {
	Validator  common.Address
	Amount     *big.Int
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposited is a free log retrieval operation binding the contract event 0x071464ed23d89f47be15656b77a0b638a220b7f85ba6bf9db649c41f64142a45.
//
// Solidity: event ValidatorDeposited(address validator, uint256 amount, uint256 commission)
func (_Storage *StorageFilterer) FilterValidatorDeposited(opts *bind.FilterOpts) (*StorageValidatorDepositedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ValidatorDeposited")
	if err != nil {
		return nil, err
	}
	return &StorageValidatorDepositedIterator{contract: _Storage.contract, event: "ValidatorDeposited", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposited is a free log subscription operation binding the contract event 0x071464ed23d89f47be15656b77a0b638a220b7f85ba6bf9db649c41f64142a45.
//
// Solidity: event ValidatorDeposited(address validator, uint256 amount, uint256 commission)
func (_Storage *StorageFilterer) WatchValidatorDeposited(opts *bind.WatchOpts, sink chan<- *StorageValidatorDeposited) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ValidatorDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageValidatorDeposited)
				if err := _Storage.contract.UnpackLog(event, "ValidatorDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeposited is a log parse operation binding the contract event 0x071464ed23d89f47be15656b77a0b638a220b7f85ba6bf9db649c41f64142a45.
//
// Solidity: event ValidatorDeposited(address validator, uint256 amount, uint256 commission)
func (_Storage *StorageFilterer) ParseValidatorDeposited(log types.Log) (*StorageValidatorDeposited, error) {
	event := new(StorageValidatorDeposited)
	if err := _Storage.contract.UnpackLog(event, "ValidatorDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageValidatorRevivedIterator is returned from FilterValidatorRevived and is used to iterate over the raw logs and unpacked data for ValidatorRevived events raised by the Storage contract.
type StorageValidatorRevivedIterator struct {
	Event *StorageValidatorRevived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageValidatorRevivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageValidatorRevived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageValidatorRevived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageValidatorRevivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageValidatorRevivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageValidatorRevived represents a ValidatorRevived event raised by the Storage contract.
type StorageValidatorRevived struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRevived is a free log retrieval operation binding the contract event 0x235b820853e16eaf1ebaf364190b4dc7d8065e1ed2c99b1fcb3590ed1a95b57a.
//
// Solidity: event ValidatorRevived(address validator)
func (_Storage *StorageFilterer) FilterValidatorRevived(opts *bind.FilterOpts) (*StorageValidatorRevivedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ValidatorRevived")
	if err != nil {
		return nil, err
	}
	return &StorageValidatorRevivedIterator{contract: _Storage.contract, event: "ValidatorRevived", logs: logs, sub: sub}, nil
}

// WatchValidatorRevived is a free log subscription operation binding the contract event 0x235b820853e16eaf1ebaf364190b4dc7d8065e1ed2c99b1fcb3590ed1a95b57a.
//
// Solidity: event ValidatorRevived(address validator)
func (_Storage *StorageFilterer) WatchValidatorRevived(opts *bind.WatchOpts, sink chan<- *StorageValidatorRevived) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ValidatorRevived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageValidatorRevived)
				if err := _Storage.contract.UnpackLog(event, "ValidatorRevived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRevived is a log parse operation binding the contract event 0x235b820853e16eaf1ebaf364190b4dc7d8065e1ed2c99b1fcb3590ed1a95b57a.
//
// Solidity: event ValidatorRevived(address validator)
func (_Storage *StorageFilterer) ParseValidatorRevived(log types.Log) (*StorageValidatorRevived, error) {
	event := new(StorageValidatorRevived)
	if err := _Storage.contract.UnpackLog(event, "ValidatorRevived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageValidatorWithdrawedIterator is returned from FilterValidatorWithdrawed and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawed events raised by the Storage contract.
type StorageValidatorWithdrawedIterator struct {
	Event *StorageValidatorWithdrawed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageValidatorWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageValidatorWithdrawed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageValidatorWithdrawed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageValidatorWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageValidatorWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageValidatorWithdrawed represents a ValidatorWithdrawed event raised by the Storage contract.
type StorageValidatorWithdrawed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawed is a free log retrieval operation binding the contract event 0x954451082939aa012a7d1e42733189160877cbae4eda956a0f3d2271b6b8365e.
//
// Solidity: event ValidatorWithdrawed(address validator)
func (_Storage *StorageFilterer) FilterValidatorWithdrawed(opts *bind.FilterOpts) (*StorageValidatorWithdrawedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ValidatorWithdrawed")
	if err != nil {
		return nil, err
	}
	return &StorageValidatorWithdrawedIterator{contract: _Storage.contract, event: "ValidatorWithdrawed", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawed is a free log subscription operation binding the contract event 0x954451082939aa012a7d1e42733189160877cbae4eda956a0f3d2271b6b8365e.
//
// Solidity: event ValidatorWithdrawed(address validator)
func (_Storage *StorageFilterer) WatchValidatorWithdrawed(opts *bind.WatchOpts, sink chan<- *StorageValidatorWithdrawed) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ValidatorWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageValidatorWithdrawed)
				if err := _Storage.contract.UnpackLog(event, "ValidatorWithdrawed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorWithdrawed is a log parse operation binding the contract event 0x954451082939aa012a7d1e42733189160877cbae4eda956a0f3d2271b6b8365e.
//
// Solidity: event ValidatorWithdrawed(address validator)
func (_Storage *StorageFilterer) ParseValidatorWithdrawed(log types.Log) (*StorageValidatorWithdrawed, error) {
	event := new(StorageValidatorWithdrawed)
	if err := _Storage.contract.UnpackLog(event, "ValidatorWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
