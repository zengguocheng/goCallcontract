package main

import (
	"./contract"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//合约地址
	contractAddress := common.HexToAddress("0xf6967fa64d4927f02434ceec3aad278683799160")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(298),
		ToBlock:   big.NewInt(1034),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//合约ＡＢＩ
	contractAbi, err := abi.JSON(strings.NewReader(string(manager.ManagerABI)))
	if err != nil {
		log.Fatal("abi error: ", err)
	}

	for _, vLog := range logs {
		event := struct {
			Addr common.Address
			From common.Address
		}{}
		err := contractAbi.Unpack(&event, "SubmitContract", vLog.Data)
		if err != nil {
			log.Fatal("unpack error: ", err)
		}

		fmt.Println(event.From.Hex())
		fmt.Println(event.Addr.Hex())

		//这个topic模式的我还没搞明白是干啥
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0])
		fmt.Println("---------------------------------------------")
	}
}
