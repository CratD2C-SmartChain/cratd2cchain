package XDCx

import (
	"github.com/CratD2C-SmartChain/cratd2cchain/accounts/abi/bind"
	"github.com/CratD2C-SmartChain/cratd2cchain/common"
	"github.com/CratD2C-SmartChain/cratd2cchain/contracts/XDCx/contract"
)

type LendingRelayerRegistration struct {
	*contract.LendingSession
	contractBackend bind.ContractBackend
}

func NewLendingRelayerRegistration(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*LendingRelayerRegistration, error) {
	smartContract, err := contract.NewLending(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &LendingRelayerRegistration{
		&contract.LendingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployLendingRelayerRegistration(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, relayerAddr common.Address, XDCxListtingAddr common.Address) (common.Address, *LendingRelayerRegistration, error) {
	contractAddr, _, _, err := contract.DeployLending(transactOpts, contractBackend, relayerAddr, XDCxListtingAddr)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewLendingRelayerRegistration(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
