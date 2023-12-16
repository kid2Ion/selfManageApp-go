package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/di"
	"github.com/kid2Ion/selfManageApp-go/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/exp/slog"
)

func main() {
	// 環境変数適応
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	slog.Info("start server")
	e := echo.New()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	// todo: migration

	// auth init
	authClient, err := firebaseauth.NewClient()
	if err != nil {
		panic(err)
	}

	// di
	// hello
	helloHandler := di.InjectHelloHandler(authClient)
	router.InitHelloRouter(e, helloHandler)
	// user
	userHandler := di.InjectUserHandler(authClient)
	router.InitUserRouter(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
