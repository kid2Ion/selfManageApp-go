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
	e.GET("/user", userHandler.Get())
	e.GET("/user/:userId", userHandler.GetByUserId())
	e.POST("/user", userHandler.Create())
	e.PUT("/user", userHandler.Update())
	e.DELETE("/user", userHandler.Delete())
}

// expense
func InitExpenseRouter(e *echo.Echo, expenseHandler server.ExpenseHandler) {
	e.POST("income/create", expenseHandler.CreateIncome())
	e.POST("outcome/create", expenseHandler.CreateOutcome())
	e.GET("expense/get", expenseHandler.GetExpense())
}
