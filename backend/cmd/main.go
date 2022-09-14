package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"server/api" // generated smart contract bindings

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var coins = map[string]int64{
	"ethereum": 0,
	"bitcoin":  0,
}

func updatePrices(client *ethclient.Client, contract *api.Api) {
	//https://api.coingecko.com/api/v3/coins/bitcoin/history?date=01-01-2020&localization=false
	for true {

		ids := ""
		index := 1
		for key := range coins {
			ids += key

			if index != len(coins) {
				ids += ","
			}
			index++
		}

		response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + ids + "&vs_currencies=usd")
		if err != nil {
			panic(err)
		}

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		response.Body.Close()

		//Convert the body to type string
		res := string(body)

		resBytes := []byte(res)                  // Converting the string "res" into byte array
		var jsonRes map[string]interface{}       // declaring a map for key names as string and values as interface{}
		err = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling

		if err != nil {
			panic(err)
		}

		// Type casting again so that interface{} -> map[string]interface{}

		for key := range coins {
			coin := jsonRes[key].(map[string]interface{}) // type the interface again to a map with key string type and value as interface
			price := coin["usd"].(float64)
			coins[key] = int64(price)
		}

		for key, value := range coins {

			auth := getAccountAuth(client, "7083aa8f4b536231bd55aa9bd49277cb03ce3ea0de097e273d529e89107196e7")

			reply, err := contract.Set(auth, key, big.NewInt(value))
			if err != nil {
				fmt.Println("Tx Failed!")
				fmt.Println(err)
			}

			fmt.Println(reply)
		}

		time.Sleep(60 * time.Second)
	}
}

type Contract struct {
	*api.Api
}

func (contract *Contract) getSymbolPrice(symbol string) (*big.Int, error) {
	reply, err := contract.PriceOf(&bind.CallOpts{}, symbol)
	return reply, err
}

type PriceChangeEvent struct {
	Symbol string
	Price  *big.Int
}

// var lastBlock = 

func fetchEvents(client *ethclient.Client, fromBlock int64, toBlock int64) {

	address := common.HexToAddress("0xCf72bbDc50e2a360F175441D9748271DFd4DC3AA")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			address,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
        fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
        fmt.Printf("Log Index: %d\n", vLog.Index)

		var transferEvent PriceChangeEvent

		err := contractAbi.UnpackIntoInterface(&transferEvent, "PriceChange", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(crypto.Keccak256Hash([]byte("ethereum")).Hex()) 

		transferEvent.Symbol = vLog.Topics[1].Hex()

		fmt.Printf("Symbol: %s\n", transferEvent.Symbol)
		fmt.Printf("Price: %s\n", transferEvent.Price.String())


        fmt.Printf("\n\n")
    }
}

func covertStringByte32(t string) [32]byte {
	var b32 [32]byte
	copy(b32[:], []byte(t))
	return b32
   }

func main() {

	// connect to rpc node
	client, err := ethclient.Dial("https://eth-goerli.g.alchemy.com/v2/dMsipp8n2rbpCS9-6ixSDYijn1zvLsBC")
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xCf72bbDc50e2a360F175441D9748271DFd4DC3AA")

	//creating api object to intract with smart contract function
	contract, err := api.NewApi(address, client)
	if err != nil {
		panic(err)
	}

	currBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	fetchEvents(client, 0, int64(currBlock))

	return

	e := echo.New()

	// Middleware
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/test", func(c echo.Context) error {
		reply, err := contract.MINDIF(&bind.CallOpts{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	// launch goroutines
	go updatePrices(client, contract)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// function to create auth for any account from its private key
func getAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth
}
