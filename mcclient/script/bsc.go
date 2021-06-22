package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/MoacLib/types"
	//"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
)

func main() {
	client, err := mcclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545")
	client.SetFuncPrefix("eth")
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
		return
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	currentBlock, _ := client.BlockNumber(context.Background())
	log.Printf("Chain ID: %d, Block Number: %d\n", chainID, currentBlock)

	///////////////////////////////////////////////
	///////////////////////////////////////////////
	privateKey, err := crypto.HexToECDSA(
		"393873d6bbc61b9d83ba923e08375b7bf8210a12bed4ea2016d96021e9378cc9",
	)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}
	log.Printf("0x%x, %d, %d, %d\n", fromAddress, chainID, nonce, balance)

	///////////////////////////////////////////////
	///////////////////////////////////////////////
	// build transactor
	Gwei := int64(1000000000)
	gasPrice := big.NewInt(int64(15) * Gwei)
	gasLimit := big.NewInt(int64(100000))

	/*
		transactor, _ := bind.NewKeyedTransactorWithChainID(
			privateKey,
			chainID,
		)
		transactor.GasPrice = big.NewInt(gasPrice)
		transactor.GasLimit = uint64(gasLimit)
		transactor.Nonce = big.NewInt(int64(nonce))
	*/

	///////////////////////////////////////////////
	///////////////////////////////////////////////
	// send ether
	value := big.NewInt(10000000000000000) // 0.01 ether
	toAddress := common.HexToAddress("0xda8ad06b2a20c6f92641d185c22f0479b00a90f3")
	data := []byte{}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewPanguSigner(chainID), privateKey)
	if err != nil {
		log.Printf("Sign tx err: %v", err)
	} else {
		log.Printf("Signed tx: %s", signedTx)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Send tx err: %v", err)
	}
}
