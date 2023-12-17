package server

import (
	"net/http"

	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/usecase"
	"github.com/labstack/echo"
	"golang.org/x/exp/slog"
)

type (
	UserHandler interface {
		Get() echo.HandlerFunc
		Create() echo.HandlerFunc
		Update() echo.HandlerFunc
		Delete() echo.HandlerFunc
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

func (t *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		fUUID, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		res, err := t.userUsecase.Get(&usecase.UserReq{
			FUUID: fUUID,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (t *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		fUUID, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		req := new(usecase.UserReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind:\n %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		req.FUUID = fUUID
		res, err := t.userUsecase.Create(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (t *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		fUUID, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		req := new(usecase.UserReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind:\n %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		req.FUUID = fUUID
		res, err := t.userUsecase.Update(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (t *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		req := new(usecase.UserReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind:\n %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		if err = t.userUsecase.Delete(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
