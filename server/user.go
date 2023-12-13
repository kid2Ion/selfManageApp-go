package server

import (
	"fmt"
	"net/http"

	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/usecase"
	"github.com/labstack/echo"
)

type (
	UserHandler interface {
		Create() echo.HandlerFunc
	}
	userHandler struct {
		userUsecase usecase.UserUsecase
		authClient  firebaseauth.AuthClient
	}
)

// コンストラクタ
func NewUserHandler(userUsecase usecase.UserUsecase, authClient firebaseauth.AuthClient) UserHandler {
	return &userHandler{userUsecase: userUsecase, authClient: authClient}
}

func (t *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		fUUID, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return err
		}
		req := new(usecase.UserReq)
		if err := c.Bind(req); err != nil {
			fmt.Println("bind失敗")
			return err
		}
		req.FUUID = fUUID
		res, err := t.userUsecase.Create(req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}
