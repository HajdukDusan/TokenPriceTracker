package server

import (
	"backendtask/blockchain"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start(priceSetter *blockchain.Contract) {
	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// create api endpoints
	CreateSymbolPriceEndpoint(e, priceSetter)
	CreateSymbolPriceHistoryEndpoint(e, priceSetter)

	// create socket channel
	CreateSocketChannel(e, priceSetter)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
