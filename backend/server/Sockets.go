package server

import (
	"backendtask/blockchain"
	"backendtask/config"
	"backendtask/model"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// Subscribes to event and broadcasts that event through a web socket when it happens
func CreateSocketChannel(e *echo.Echo, priceSetter *blockchain.Contract) {

	ch := make(chan *model.PriceChangeEvent)

	initWebSocketEndpoint(e, ch)

	addr := common.HexToAddress(priceSetter.Address)

	go blockchain.SubscribeToEvent(
		[]common.Address{addr},
		ch,
	)
}

// Helper func for creating the web socket endpoint
func initWebSocketEndpoint(e *echo.Echo, ch chan *model.PriceChangeEvent) {
	
	// load symbol hashes so we can return the symbol from its hash value
	symbolHashes := make(map[string]string)
	for _, symbol := range config.Symbols {
		hash := crypto.Keccak256Hash([]byte(symbol))
		symbolHashes[hash.Hex()] = symbol
	}

	upgrader := websocket.Upgrader{}

	e.GET("/ws", func(c echo.Context) error {

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()
	
		for {
			event := <-ch

			eventDTO, err := ConvertPriceEventToDTO(*event, symbolHashes)
			if err != nil {
				fmt.Println("Failed to convert object!")
				continue
			}

			fmt.Println("\nPrice change event broadcasted!")
			fmt.Println(eventDTO)

			json, _ := json.Marshal(eventDTO)

			// write to ws
			err = ws.WriteMessage(websocket.TextMessage, []byte(string(json)))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	})
}