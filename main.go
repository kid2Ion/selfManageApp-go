package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kid2Ion/selfManageApp-go/di"
	"github.com/kid2Ion/selfManageApp-go/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("start server")
	helloHandler := di.InjectHandler()
	e := echo.New()

	// 単純でないリクエストを許可
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	//todo: IPアクセス制限
	// CSRFでやる？

	router.InitHelloRouter(e, helloHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
