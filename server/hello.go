package server

import (
	"net/http"

	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/usecase"
	"github.com/labstack/echo"
)

type (
	HelloHandler interface {
		Hello() echo.HandlerFunc
	}

	helloHandler struct {
		helloUsecase usecase.HelloUsecase
		authClient   firebaseauth.AuthClient
	}
)

// コンストラクタ
func NewHelloHandler(helloUsecase usecase.HelloUsecase, authClient firebaseauth.AuthClient) HelloHandler {
	return &helloHandler{helloUsecase: helloUsecase, authClient: authClient}
}

func (t *helloHandler) Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return err
		}
		res := t.helloUsecase.Hello()
		return c.JSON(http.StatusOK, res)
	}
}
