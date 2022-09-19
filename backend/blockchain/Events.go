package blockchain

import (
	"backendtask/api"
	"backendtask/model"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Is used to fetch contract PriceChange events.
//
// Params:
//
// - List of contract addresses to be scanned for events.
//
// - Block to start scanning from (nil will start from 0).
//
// - Block to end scanning (nil refers to the currentBlock).
//
// - Matrix for topics (indexed fields of events).
func FetchEvents(
	addresses []common.Address,
	fromBlock *big.Int,
	toBlock *big.Int,
	indexedValues [][]common.Hash,
) ([]model.PriceChangeEvent, error) {

	// check if the rpc connection is initialized
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

	// create a query for filtering logs
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: addresses,
		Topics:    indexedValues,
	}

	return getEventsFromLogs(query)
}


func SubscribeToEvent(addresses []common.Address, eventCh chan *model.PriceChangeEvent) {

	query := ethereum.FilterQuery{
        Addresses: addresses,
    }

    logs := make(chan types.Log)
    sub, err := Client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

	// get a contract abi to unpack log data
	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		fmt.Println("Subscribe: ", err)
        return
	}

    for {
        select {
        case err := <-sub.Err():
            fmt.Println(err)
        case log := <-logs:

			var event model.PriceChangeEvent
	
			err := contractAbi.UnpackIntoInterface(&event, "PriceChange", log.Data)
			if err != nil {
				fmt.Println("Error parsing event: ", err)
			}
	
			event.BlockNumber = log.BlockNumber
			event.SymbolHash = log.Topics[1].Hex()

			eventCh <- &event
        }
    }
}

// helper function for fetching and parsing logs
func getEventsFromLogs(query ethereum.FilterQuery) ([]model.PriceChangeEvent, error) {
	
	// get a contract abi to unpack log data
	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		return nil, err
	}

	// fetch all the logs that pass the query filter
	logs, err := Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	events := make([]model.PriceChangeEvent, len(logs))

	for indx, vLog := range logs {

		var event model.PriceChangeEvent

		err := contractAbi.UnpackIntoInterface(&event, "PriceChange", vLog.Data)
		if err != nil {
			return nil, err
		}

		event.BlockNumber = vLog.BlockNumber
		event.SymbolHash = vLog.Topics[1].Hex()

		events[indx] = event
	}
	return events, nil
}
