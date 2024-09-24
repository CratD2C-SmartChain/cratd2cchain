package main

import (
	"context"
	"fmt"
	"github.com/XinFinOrg/XDC-Subnet/contracts/XDCx/simulation"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/XinFinOrg/XDC-Subnet/accounts/abi/bind"
	"github.com/XinFinOrg/XDC-Subnet/common"
	"github.com/XinFinOrg/XDC-Subnet/contracts/XDCx"
	"github.com/XinFinOrg/XDC-Subnet/crypto"
	"github.com/XinFinOrg/XDC-Subnet/ethclient"
)

func main() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}

	MainKey, _ := crypto.HexToECDSA(os.Getenv("OWNER_KEY"))
	MainAddr := crypto.PubkeyToAddress(MainKey.PublicKey)
	coinbase := common.HexToAddress(os.Getenv("RELAYER_COINBASE"))
	fee, _ := strconv.Atoi(os.Getenv("FEE"))

	nonce, _ := client.NonceAt(context.Background(), MainAddr, nil)
	auth := bind.NewKeyedTransactor(MainKey)
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(250000000000000)

	auth.Nonce = big.NewInt(int64(nonce))

	registrationContract, _ := XDCx.NewRelayerRegistration(auth, common.HexToAddress("0x0342d186212b04E69eA682b3bed8e232b6b3361a"), client)

	tx, err := registrationContract.UpdateFee(coinbase, uint16(fee))
	if err != nil {
		fmt.Println("UpdateFee: failed!", err)
	}

	time.Sleep(5 * time.Second)
	r, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		fmt.Println("UpdateFee: Get receipt failed", err)
	}
	fmt.Println("UpdateFee: Done receipt status", r.Status)

}
