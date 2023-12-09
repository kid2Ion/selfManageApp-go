package router

import (
	"github.com/kid2Ion/selfManageApp-go/server"
	"github.com/labstack/echo"
)

// hello
func InitHelloRouter(e *echo.Echo, helloHandler server.HelloHandler) {
	e.GET("/", helloHandler.Hello())
}

// hogehoge
