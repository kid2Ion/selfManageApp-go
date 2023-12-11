package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/di"
	"github.com/kid2Ion/selfManageApp-go/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// 環境変数適応
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("start server")
	e := echo.New()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	//todo: logger

	// auth init
	authClient, err := firebaseauth.NewClient()
	if err != nil {
		panic(err)
	}

	// di
	helloHandler := di.InjectHandler(authClient)
	router.InitHelloRouter(e, helloHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
