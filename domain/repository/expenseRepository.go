package repository

import "github.com/kid2Ion/selfManageApp-go/domain/model"

type (
	ExpenseRepository interface {
		CreateIncome(i *model.Income) error
		GetIncome(expenseUUID string) (*model.Income, error)
		CreateOutcome(o *model.Outcome) error
		GetOutcomes(expenseUUID string) ([]model.Outcome, error)
		GetExpenseUUID(e *model.Expense) (string, error)
		CreateExpense(e *model.Expense) error
	}
)
