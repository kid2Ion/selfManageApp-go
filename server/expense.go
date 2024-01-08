package server

import (
	"net/http"

	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/usecase"
	"github.com/labstack/echo"
	"golang.org/x/exp/slog"
)

type (
	ExpenseHandler interface {
		CreateIncome() echo.HandlerFunc
		CreateOutcome() echo.HandlerFunc
	}
	expenseHandler struct {
		expenseUsecase usecase.ExpenseUsecase
		authClient     firebaseauth.AuthClient
	}
)

func NewExpenseHandler(
	expenseUsecase usecase.ExpenseUsecase,
	authClient firebaseauth.AuthClient,
) ExpenseHandler {
	return &expenseHandler{expenseUsecase: expenseUsecase, authClient: authClient}
}

func (t *expenseHandler) CreateIncome() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		req := new(usecase.IncomeReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind:\n %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		err = t.expenseUsecase.CreateIncome(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}

func (t *expenseHandler) CreateOutcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := t.authClient.VerifyIDToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		req := new(usecase.OutcomeReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind:\n %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		err = t.expenseUsecase.CreateOutcome(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
