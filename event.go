package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"contractRoundRobin/contract"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xf6967fa64d4927f02434ceec3aad278683799160")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	//合约ＡＢＩ
	contractAbi, err := abi.JSON(strings.NewReader(string(manager.ManagerABI)))
	if err != nil {
		log.Fatal("abi error: ", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// fmt.Println(vLog) // pointer to event log

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
}
