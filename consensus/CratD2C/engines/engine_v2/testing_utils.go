package engine_v2

import (
	"github.com/CratD2C-SmartChain/cratd2cchain/common"
	"github.com/CratD2C-SmartChain/cratd2cchain/consensus"
	"github.com/CratD2C-SmartChain/cratd2cchain/core/types"
)

/*
	Testing tools
*/

func (x *CratD2C_v2) SetNewRoundFaker(blockChainReader consensus.ChainReader, newRound types.Round, resetTimer bool) {
	x.lock.Lock()
	defer x.lock.Unlock()
	// Reset a bunch of things
	if resetTimer {
		x.timeoutCount = 0
		x.timeoutWorker.Reset(blockChainReader)
	}
	x.currentRound = newRound
}

// for test only
func (x *CratD2C_v2) ProcessQCFaker(chain consensus.ChainReader, qc *types.QuorumCert) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.processQC(chain, qc)
}

// Utils for test to check currentRound value
func (x *CratD2C_v2) GetCurrentRoundFaker() types.Round {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.currentRound
}

// Utils for test to get current Pool size
func (x *CratD2C_v2) GetVotePoolSizeFaker(vote *types.Vote) int {
	return x.votePool.Size(vote)
}

// Utils for test to get Timeout Pool Size
func (x *CratD2C_v2) GetTimeoutPoolSizeFaker(timeout *types.Timeout) int {
	return x.timeoutPool.Size(timeout)
}

// WARN: This function is designed for testing purpose only!
// Utils for test to check currentRound values
func (x *CratD2C_v2) GetPropertiesFaker() (types.Round, *types.QuorumCert, *types.QuorumCert, *types.TimeoutCert, types.Round, *types.BlockInfo) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.currentRound, x.lockQuorumCert, x.highestQuorumCert, x.highestTimeoutCert, x.highestVotedRound, x.highestCommitBlock
}

// WARN: This function is designed for testing purpose only!
// Utils for tests to set engine specific values
func (x *CratD2C_v2) SetPropertiesFaker(highestQC *types.QuorumCert, highestTC *types.TimeoutCert) {
	x.highestQuorumCert = highestQC
	x.highestTimeoutCert = highestTC
}

func (x *CratD2C_v2) HygieneVotePoolFaker() {
	x.hygieneVotePool()
}

func (x *CratD2C_v2) GetVotePoolKeyListFaker() []string {
	return x.votePool.PoolObjKeysList()
}

func (x *CratD2C_v2) HygieneTimeoutPoolFaker() {
	x.hygieneTimeoutPool()
}

func (x *CratD2C_v2) GetTimeoutPoolKeyListFaker() []string {
	return x.timeoutPool.PoolObjKeysList()
}

// Fake the signer address, the signing function is incompatible
func (x *CratD2C_v2) AuthorizeFaker(signer common.Address) {
	x.signLock.Lock()
	defer x.signLock.Unlock()

	x.signer = signer
}

func (x *CratD2C_v2) GetForensicsFaker() *Forensics {
	return x.ForensicsProcessor
}
