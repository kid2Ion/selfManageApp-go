package repository

import "github.com/kid2Ion/selfManageApp-go/domain/model"

type (
	ExpenseRepository interface {
		CreateIncome(i *model.Income) error
		GetExpenseUUID(e *model.Expense) (string, error)
		CreateExpense(e *model.Expense) error
	}
)
