package server

import (
	"fmt"
	"net/http"

	"github.com/kid2Ion/selfManageApp-go/usecase"
	"github.com/labstack/echo"
)

type (
	HelloHandler interface {
		Hello() echo.HandlerFunc
	}

	helloHandler struct {
		helloUsecase usecase.HelloUsecase
	}
)

// コンストラクタ
func NewHelloHandler(helloUsecase usecase.HelloUsecase) HelloHandler {
	return &helloHandler{helloUsecase: helloUsecase}
}

func (t *helloHandler) Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := t.helloUsecase.Hello()
		fmt.Println(res)
		return c.JSON(http.StatusOK, res)
	}
}
