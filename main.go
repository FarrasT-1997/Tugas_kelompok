package main

import (
	"fmt"
	"rest/api/config"
	"rest/api/router"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init_DB()
	config.InitPort()
	router.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
