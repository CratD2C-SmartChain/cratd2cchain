#!/bin/bash

# Variables
params=""

# constants
KEYSTORE_DIR="keystore"
DATA_DIR="/work/cratd2cchain"

# networkid
if [[ ! -z $NETWORK_ID ]]; then
  echo "Network Id set to ${NETWORK_ID}"
  params="$params --networkid $NETWORK_ID"
else
  echo "NETWORK_ID environment variable has not been set. Default to 65349"
  params="$params --networkid 65349"
fi

# bootnodes
if [[ ! -z $BOOTNODES ]]; then
  echo "Received bootnodes of ${BOOTNODES}"
  params="$params --bootnodes $BOOTNODES"
fi

# extip
if [[ ! -z $EXTIP ]]; then
echo "Set the NAT to extip:${EXTIP}"
  params="$params --nat extip:${EXTIP}"
fi

# port
if [[ ! -z $PORT ]]; then
echo "Set PORT to ${PORT}"
  params="$params --port ${PORT}"
else
  params="$params --port 30303"
fi

# rpcport
if [[ ! -z $RPCPORT ]]; then
echo "Set RPCPORT to ${RPCPORT}"
  params="$params --rpcport ${RPCPORT}"
else
  params="$params --rpcport 8545"
fi

# wsport
if [[ ! -z $WSPORT ]]; then
echo "Set WSPORT ${WSPORT}"
  params="$params --wsport ${WSPORT}"
else
  params="$params --wsport 8555"
fi

# syncmode
if [[ ! -z $SYNC_MODE ]]; then
  echo "Set the sync mode to ${SYNC_MODE}"
  params="$params --syncmode ${SYNC_MODE}"
else
  params="$params --syncmode full"
fi

# log level
if [[ ! -z $LOG_LEVEL ]]; then
  echo "Set the log level to ${LOG_LEVEL}"
  params="$params --verbosity ${LOG_LEVEL}"
else
  params="$params --verbosity 2"
fi

# GC mode
if [[ ! -z $GC_MODE ]]; then
  echo "Set the GC mode to ${GC_MODE}"
  params="$params --gcmode ${GC_MODE}"
else
  params="$params --gcmode full"
fi

# RPC API
if [ ! -d $DATA_DIR/keystore ]
then
  if test -z "$PRIVATE_KEY"
  then
    echo "PRIVATE_KEY environment variable has not been set, randomly creating a new one"

    wallet=$(CRAT --datadir $DATA_DIR account new \
      --password /work/.pwd 2>&1 | grep -oE '0x[a-fA-F0-9]{40}' | head -n1)

    if [ -z "$wallet" ]; then
      echo "ERROR: Failed to create new account!"
      exit 1
    fi

    echo "Created new wallet: $wallet"

    CRAT --datadir $DATA_DIR init /work/genesis.json

  else
    echo "${PRIVATE_KEY}" > ./private_key
    echo "Creating account from private key"

    wallet=$(CRAT --datadir $DATA_DIR account import \
      --password /work/.pwd \
      ./private_key 2>&1 | awk -v FS="({|})" '{print $2}')

    if [ -z "$wallet" ]; then
      echo "ERROR: Failed to import account!"
      exit 1
    fi

    CRAT --datadir $DATA_DIR init /work/genesis.json
    rm -f ./private_key
  fi
else
  echo "Wallet already exist, re-use the same one"
  wallet=$(CRAT --datadir $DATA_DIR account list | head -n 1 | awk -v FS="({|})" '{print $2}')
fi



echo "Using wallet: $wallet"



# Stats server
if [[ ! -z $STATS_SERVICE_ADDRESS ]]; then
  echo "Setting up stats server communication to ${STATS_SERVICE_ADDRESS} with name ${INSTANCE_NAME}-${wallet}"
  statsSecret="subnet-stats-server"
  if [ ! -z $STATS_SECRET ]; then
    statsSecret="${STATS_SECRET}"
  fi
  
  statsHostName=""
  if [ ! -z $INSTANCE_NAME ]; then
    statsHostName="${INSTANCE_NAME}-${wallet}"
  else
    statsHostName="${wallet}"
  fi
  
  netstats="${statsHostName}:${statsSecret}@${STATS_SERVICE_ADDRESS}"
  echo "Sending events to stats service at ${netstats}"
  
  params="$params --ethstats ${netstats}"
else
  echo "STATS_SERVICE_ADDRESS not set. Skipping the stats server set up. Won't emit any messages"
fi

echo "Using wallet $wallet"
params="$params --unlock $wallet"

echo "Starting nodes with bootnodes of: $BOOTNODES ..."

CRAT $params \
--datadir $DATA_DIR \
--rpc \
--rpccorsdomain "*" \
--rpcaddr 127.0.0.1 \
--rpcvhosts "*" \
--password /work/.pwd \
--gasprice "1" \
--targetgaslimit "420000000" \
--ws --wsaddr=127.0.0.1 \
--wsorigins "*" 2>&1 >>$DATA_DIR/cratd2c.log | tee --append $DATA_DIR/cratd2c.log