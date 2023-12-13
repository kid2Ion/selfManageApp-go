package router

import (
	"github.com/kid2Ion/selfManageApp-go/server"
	"github.com/labstack/echo"
)

// hello
func InitHelloRouter(e *echo.Echo, helloHandler server.HelloHandler) {
	e.GET("/", helloHandler.Hello())
}

// user
func InitUserRouter(e *echo.Echo, userHandler server.UserHandler) {
	e.POST("/user", userHandler.Create())
}
