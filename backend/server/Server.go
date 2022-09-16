package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start(contractAddress string) {
	e := echo.New()

	// Middleware
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	CreateSymbolPriceHistoryEndpoint(e, contractAddress)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
