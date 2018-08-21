package main

import (
	"./contract"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

const key = `{"address":"2450d888013ae20ac411823d514897059027b8e5","crypto":{"cipher":"aes-128-ctr","ciphertext":"8ad9f7f2c733f2517798f9ca58715a58e863635825eca1108132f0e038635992","cipherparams":{"iv":"2c0683c40ca5f1640daedc33de7105b7"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"6a928761dd4a1b7e704c6ee3541d443dde1893bcf716a52049f8c2d4bf8a09f0"},"mac":"3340276e4a26eb0298260058a6801b04a4d8b9f66e6bab625ad44d05110d9a18"},"id":"69196881-2e26-45c1-953a-f7c8b1205064","version":3}`

func main() {
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Deploy a new awesome contract for the binding demo
	address, tx, instanc, err := manager.DeployManager(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	var _ = instanc
}
