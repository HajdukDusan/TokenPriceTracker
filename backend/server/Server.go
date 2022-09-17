package server

import (
	"backendtask/blockchain"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start(priceSetter *blockchain.Contract) {
	e := echo.New()

	// Middleware
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// create api endpoints
	CreateSymbolPriceEndpoint(e, priceSetter)
	CreateSymbolPriceHistoryEndpoint(e, priceSetter)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
