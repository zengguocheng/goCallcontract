package main

import (
	"context"
	"contractRoundRobin/contract"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("conn err !!!")
	}

	address := common.HexToAddress("0x77a031e3Eb1a0f406b0d883E93bec9540cC10d1C")
	instance, err := manager.NewManager(address, conn)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("04e613ccfe7df93571a0965165c4c6f53b9b0b3ac633b2ffb71000fdcf07b5ff")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	tx, err := instance.SubmitAddress(auth, address)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
