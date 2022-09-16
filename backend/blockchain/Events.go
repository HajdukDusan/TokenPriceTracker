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
)

// Examples:
// {} or nil          matches any topic list
// {{A}}              matches topic A in first position
// {{}, {B}}          matches any topic in first position AND B in second position
// {{A}, {B}}         matches topic A in first position AND B in second position
// {{A, B}, {C, D}}   matches topic (A OR B) in first position AND (C OR D) in second position

func FetchEvents(
	addresses []common.Address,
	fromBlock *big.Int,
	toBlock *big.Int,
	indexedValues [][]common.Hash,
) ([]model.PriceChangeEvent, error) {

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

	events := make([]model.PriceChangeEvent, len(logs))

	for indx, vLog := range logs {

		var event model.PriceChangeEvent

		err := contractAbi.UnpackIntoInterface(&event, "PriceChange", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		event.BlockNumber = vLog.BlockNumber
		event.SymbolHash = vLog.Topics[1].Hex()

		events[indx] = event
	}

	return events, nil
}
