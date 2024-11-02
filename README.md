## Go CratD2C

Golang execution layer implementation of the CratD2C protocol (based on Ethereum).

## Building the source

Building `crat` requires both a Go (version 1.21 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make CRAT
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-cratd2c project comes with several wrappers/executables found in the `cmd`
directory.

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|:----------:|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **`crat`** | Our main CratD2C CLI client. It is the entry point into the CratD2C network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ethereum network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and command line options. | |
|  `devp2p`  | Utilities to interact with nodes on the networking layer, without running a full blockchain.                                                                                                                                                                                                                                                                                                                                                   |
|  `abigen`  | Source code generator to convert CratD2C contract definitions into easy-to-use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined.                                                    |
| `bootnode` | Stripped down version of our CratD2C client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.                                                                                                                                                            |
|   `evm`    | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                           | |

## Running `crat`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://geth.ethereum.org/docs/fundamentals/command-line-options)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `geth` instance.

### Hardware Requirements

Minimum:

* CPU with 4+ cores
* 8GB RAM
* 1TB free storage space to sync the Mainnet
* 8 MBit/sec download Internet service

Recommended:

* Fast CPU with 8+ cores
* 32GB+ RAM
* High-performance SSD with at least 1TB of free space
* 25+ MBit/sec download Internet service

### Full node on the main CratD2C network

By far the most common scenario is people wanting to simply interact with the CratD2C
network: create accounts; transfer funds; deploy and interact with contracts. For this
particular use case, the user doesn't care about years-old historical data, so we can
sync quickly to the current state of the network. To do so:

```shell
$ crat console
```

This command will:
* Start `crat` in snap sync mode (default, can be changed with the `--syncmode` flag),
  causing it to download more data in exchange for avoiding processing the entire history
  of the Ethereum network, which is very CPU intensive.

### Configuration

As an alternative to passing the numerous flags to the `crat` binary, you can also pass a
configuration file via:

```shell
$ crat --config /path/to/your_config.toml
```

#### Docker quick start

One of the quickest ways to get CratD2C up and running on your machine is by using
Docker:

```shell
docker run -d --name crat-node -v /Users/Sammy/crat:/root \
           -p 8545:8545 -p 30303:30303 \
           crat/client-go
```

This will start `crat` in snap-sync mode with a DB memory allowance of 1GB, as the
above command does.  It will also create a persistent volume in your home directory for
saving your blockchain as well as map the default ports.

#### Defining the private genesis state

First, you'll need to create the genesis state of your networks, which all nodes need to be
aware of and agree upon. This consists of a small JSON file (e.g. call it `genesis.json`):

```json
{
  "config": {
    "chainId": 65349,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "CratDPoS": {
      "networkName": "CratD2C",
      "denom": "0x",
      "period": 2,
      "epoch": 900,
      "reward": 2,
      "rewardCheckpoint": 900,
      "gap": 450,
      "foudationWalletAddr": <addr>,
      "SkipV1Validation": false,
      "v2": {
        "switchBlock": 0,
        "config": {
          "switchRound": 0,
          "minePeriod": 2,
          "timeoutSyncThreshold": 3,
          "timeoutPeriod": 10,
          "certificateThreshold": 2
        },
        "allConfigs": {
          "0": {
            "switchRound": 0,
            "minePeriod": 2,
            "timeoutSyncThreshold": 3,
            "timeoutPeriod": 10,
            "certificateThreshold": 2
          }
        },
        "SkipV2Validation": false
      }
  },
  "alloc": {},
  "coinbase": "0x0000000000000000000000000000000000000000",
  "difficulty": "0x20000",
  "extraData": "",
  "gasLimit": "0x2fefd8",
  "nonce": "0x0000000000000042",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
}
```

The above fields should be fine for most purposes, although we'd recommend changing
the `nonce` to some random value so you prevent unknown remote nodes from being able
to connect to you. If you'd like to pre-fund some accounts for easier testing, create
the accounts and populate the `alloc` field with their addresses.

```json
"alloc": {
  "0x0000000000000000000000000000000000000001": {
    "balance": "111111111"
  },
  "0x0000000000000000000000000000000000000002": {
    "balance": "222222222"
  }
}
```

With the genesis state defined in the above JSON file, you'll need to initialize **every**
`crat` node with it prior to starting it up to ensure all blockchain parameters are correctly
set:

```shell
$ crat init path/to/genesis.json
```

#### Creating the rendezvous point

With all nodes that you want to run initialized to the desired genesis state, you'll need to
start a bootstrap node that others can use to find each other in your network and/or over
the internet. The clean way is to configure and run a dedicated bootnode:

```shell
$ bootnode --genkey=boot.key
$ bootnode --nodekey=boot.key
```

With the bootnode online, it will display an [`enode`]
that other nodes can use to connect to it and exchange peer information. Make sure to
replace the displayed IP address information (most probably `[::]`) with your externally
accessible IP to get the actual `enode` URL.

*Note: You could also use a full-fledged `crat` node as a bootnode, but it's the less
recommended way.*

## Contribution

Thank you for considering helping out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

TBA

## License

TBA
