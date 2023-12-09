package main

import (
	"fmt"

	"github.com/kid2Ion/selfManageApp-go/di"
	"github.com/kid2Ion/selfManageApp-go/router"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("start server")
	helloHandler := di.InjectHandler()
	e := echo.New()
	router.InitHelloRouter(e, helloHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
