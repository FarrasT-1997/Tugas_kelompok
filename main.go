package main

import (
	"fmt"
	"rest_api/config"
	"rest_api/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init_DB()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":8080")))
}
