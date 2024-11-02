# CratDPoS - Crat's Consensus Mechanism

CratD2C Delegated Proof of Stake(CratDPoS) is another form of Proof of Stake(PoS) consensus mechanism to scale up to Thousands of transactions per second. CratDPOS concentrates block production in the hands of few semi-trusted entities in order to achieve more scalability than other Proof of Stake blockchains. CratDPOS leverages the power of stakeholder to resolve consensus issues in a fair and democratic way.

## Common Terms
* **Delegator:** 
* **Validator:** A semi-trusted entity responsible for validating and producing blocks.
* **Epoch:** A cycle of few blocks in which validator nodes create blocks in turn.

## Transition from PoW to CratDPoS
The traditional PoW mechanism, while effective in securing networks, is highly energy-intensive due to its reliance on solving complex cryptographic puzzles across all nodes in the network. This results in significant electricity consumption and poses environmental concerns. Recognizing these drawbacks, CratD2C transitioned away from PoW and adopted Proof of Stake (PoS), a more energy-efficient consensus methodology. Eventually, CratD2C advanced to CratD2C Delegated Proof of Stake (XDPoS), which further enhances the efficiency and security of the network.

To understand CratD2C Network's decision to adopt CratDPoS, it's essential to explore the various consensus mechanisms.

## Understanding Consensus: PoW vs. PoS vs. CratDPoS
In blockchain technology, consensus refers to a general agreement among nodes in the network, which is crucial for maintaining the integrity and accuracy of the distributed ledger. Unlike centralized systems like banks, where a central authority maintains records, blockchain relies on distributed ledgers to record information. Consensus ensures that all nodes agree on the state of the blockchain, including account balances, transactions, and more.

**Proof of Work (PoW)**
PoW relies on nodes solving cryptographic puzzles to validate transactions and create new blocks. This process requires significant computational power and consumes vast amounts of electricity, making it environmentally unsustainable and costly.

**Proof of Stake (PoS)**
PoS eliminates the need for energy-intensive computations. Instead of miners, PoS relies on validators who lock a portion of their cryptocurrency as a stake. Validators are chosen to propose the next block based on their stake and uptime. When a validator discovers a block that can be added to the blockchain, they validate it by placing a bet on it. The validators receive rewards proportional to their bets.

**CratD2C Delegated Proof of Stake (CratDPoS)**
CratDPoS is an evolution of PoS and offers greater efficiency. It uses a reputation-based system to achieve consensus, where master nodes create blocks in a round-robin manner. The network elects block producers (also known as witnesses) who are responsible for validating transactions and creating the next block. The key features of CratDPoS include:

* **Weighted Random Election of Block Producers:** Block producers are elected weighted randomly (weights depend on stake amount) and are limited in number. They are responsible for signing and creating blocks.
* **Block Validators:** These full nodes verify that the blocks created by block producers adhere to consensus rules.
* **Delegator Governance:** Unlike PoS, CratDPoS provides the community with more governance rights, allowing them to have a voice in validators election.

## Pros and Cons of CratD2C Delegated Proof of Stake
## Advantages

* **Speed:** CratDPoS enables faster transaction processing due to the limited number of witnesses involved.
* **Cost Efficiency:** The reduction in the number of participants in block creation lowers transaction costs.
* **Scalability:** CratDPoS improves the network's scalability by ensuring efficient block creation and validation.
* **Security:** The reputation system and regular monitoring of witnesses help maintain the integrity of the network. Malicious actors can be penalized according to performance.

## Disadvantages
* **Centralization Risk:** Since only validators can create blocks, there is a risk of centralization if these nodes collude.
* **Limited Participation:** The limited number of block producers (101) may reduce the inclusiveness of the consensus process.

## Why CratD2C Network Prefers CratDPoS
CratD2C Network chose CratDPoS for its network due to its superior speed, cost efficiency, and scalability. CratDPoS offers a more democratic, faster, and effective way to scale the network, making it the ideal consensus mechanism for CratD2C smartchain platform.


## Rewards Mechanism

The CratD2C Network operates on a Delegated Proof of Stake (CratDPoS) consensus mechanism, which allows for high transaction throughput, energy efficiency, and security. The network’s rewards mechanism is designed to incentivize validators and participants who contribute to the stability, security, and growth of the network. This document outlines how the CratD2C rewards mechanism works, who is eligible for rewards, and how rewards are distributed.

## Overview of CratD2C Rewards System
In the CratD2C Network, rewards are primarily distributed to validators who validate transactions, secure the network, and propose new blocks. Additionally, Crat token holders who delegate their tokens to validators also earn a share of the rewards. The rewards mechanism is structured to:

Incentivize decentralization by encouraging more participants to run validator or delegate their tokens.
Reward active participation in the network, such as block validation and staking.
Ensure the economic sustainability of the CratD2C Network.

## Validator's Rewards
Validators are critical to the CratD2C Network’s operations, responsible for validating transactions and maintaining consensus. To compensate for their role, validators receive rewards from every transaction in block they validate.

- **Eligibility:** To run a validator, an operator must stake a minimum of 100_000 Crat tokens.
- **Reward Distribution:** Validators are rewarded for validating blocks, and the rewards are distributed proportionally based on their performance and uptime.
- **Block Rewards:** No fixed reward is allocated for each block validated by the network.

**Note:** Validators with better uptime and performance will be secured for slashing procedures.

