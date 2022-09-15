package blockchain

import (
	"backendtask/api"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PriceChangeEvent struct {
	SymbolHash string
	Price  *big.Int
}

type PriceChangeEventDTO struct {
	Symbol string
	Price int64
}


func (pce PriceChangeEvent) Print() {
	fmt.Printf("Symbol: %s\n", pce.SymbolHash)
	fmt.Printf("Price: %s\n", pce.Price.String())
}


func FetchEventsBySymbolsAndTimeStamp(contractAddress string, fromTimestamp int64, toTimestamp int64, symbols []string) ([]PriceChangeEventDTO, error) {

	// TODO TIMESTAMPS

	indexedValues := make([]common.Hash, len(symbols))
	symbolHashMap := make(map[string]string)

	// hash the array of symbols because they are indexed in events
	for indx := range symbols {
		hash := crypto.Keccak256Hash([]byte(symbols[indx]))
		indexedValues[indx] = hash
		symbolHashMap[hash.Hex()] = symbols[indx]
	}

	events, err := fetchEvents(
		[]common.Address{common.HexToAddress(contractAddress)},
		nil,
		nil,
		[][]common.Hash{
			{},
			indexedValues,
		},
	)
	if err != nil {
		return nil, err
	}

	eventsDTO := make([]PriceChangeEventDTO, len(events))

	// get symbol from symbolHashMap with symbol hash as key 
	for indx, event := range events {
		eventsDTO[indx] = PriceChangeEventDTO{
			Symbol: symbolHashMap[event.SymbolHash],
			Price: event.Price.Int64(),
		}
	}

	return eventsDTO, nil
}

// Examples:
// {} or nil          matches any topic list
// {{A}}              matches topic A in first position
// {{}, {B}}          matches any topic in first position AND B in second position
// {{A}, {B}}         matches topic A in first position AND B in second position
// {{A, B}, {C, D}}   matches topic (A OR B) in first position AND (C OR D) in second position

func fetchEvents(addresses []common.Address, fromBlock *big.Int, toBlock *big.Int, indexedValues [][]common.Hash) ([]PriceChangeEvent, error) {

	// check if the rpc conn is initialized
	if Client == nil {
		return nil, errors.New("RPC connection is not initialized, try to call Connect() first...")
	}

	// if toBlock is nil get current blockNumber
	if toBlock == nil {
		currBlock, err := Client.BlockNumber(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		toBlock = big.NewInt(int64(currBlock))
	}

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: addresses,
		Topics:    indexedValues,
	}

	logs, err := Client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	events := make([]PriceChangeEvent, len(logs))

	for indx, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)

		var event PriceChangeEvent

		err := contractAbi.UnpackIntoInterface(&event, "PriceChange", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		event.SymbolHash = vLog.Topics[1].Hex()

		event.Print()

		events[indx] = event
	}

	return events, nil
}
