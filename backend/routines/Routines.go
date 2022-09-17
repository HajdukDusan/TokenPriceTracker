package routines

import (
	"backendtask/blockchain"
	"backendtask/config"
	"backendtask/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

func UpdatePrices(coins map[string]int64) {

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

		fmt.Println("New Prices:")

		for key := range coins {
			coin := jsonRes[key].(map[string]interface{}) // type the interface again to a map with key string type and value as interface
			price := coin["usd"].(float64)

			fmt.Printf("\t" + key + " -> %f USD\n", price)

			result, err := utils.ScaleFloatToInt(price)
			if err != nil {
				fmt.Println(err)
				continue
			}

			coins[key] = result
		}

		fmt.Println()

		// Init the block of txs
		SendTx := blockchain.CreateTxBlock()

		for symbol, price := range coins {

			// get symbol price from contract
			contractPrice, err := blockchain.PriceSetterContract.GetSymbolPrice(symbol)
			if err != nil {
				fmt.Println(err)
			}

			absDif := utils.GetAbs(big.NewInt(0).Sub(contractPrice, big.NewInt(price)))

			minDif := big.NewInt(2)

			if absDif.Mul(absDif, big.NewInt(100)).Cmp(minDif.Mul(contractPrice, minDif)) == 1 {
				tx, err := SendTx(symbol, big.NewInt(price))
				if err != nil {
					fmt.Println("Failed to send tx!")
					fmt.Println(err)
				} else {
					fmt.Printf("Tx successfully sent: %s\n", tx.Hash().Hex())
				}
			}
		}

		// stop thread for config.UpdateInterval seconds
		time.Sleep(time.Duration(config.UpdateInterval) * time.Second)
	}
}