package infra

import (
	"github.com/kid2Ion/selfManageApp-go/domain/model"
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
	"golang.org/x/exp/slog"
)

type (
	ExpenseRepository struct {
		SqlHandler
	}
)

func NewExpenseRepository(sqlHandler SqlHandler) repository.ExpenseRepository {
	return &ExpenseRepository{SqlHandler: sqlHandler}
}

func (t *ExpenseRepository) CreateIncome(i *model.Income) error {
	cmd := `insert into expense.incomes (income_uuid, expense_uuid, amount, created_at, updated_at) values ($1, $2, $3, $4, $5);`
	_, err := t.SqlHandler.DB.Exec(cmd, i.IncomeUUID, i.ExpenseUUID, i.Amount, i.CreatedAt, i.UpdatedAt)
	if err != nil {
		slog.Error("failed to create income:\n %s", err.Error())
		return err
	}
	return nil
}

func (t *ExpenseRepository) GetExpenseUUID(e *model.Expense) (string, error) {
	q := `select e.expense_uuid from expense.expenses e
where e.user_uuid = $1
and e.year = $2
and e.month = $3
;`

	rows, err := t.SqlHandler.DB.Query(q, e.UserUUID, e.Year, e.Month)
	if err != nil {
		slog.Error("failed to fetch from db: %s", err.Error())
		return "", err
	}

	defer rows.Close()
	res := ""
	if rows.Next() {
		var eUUID string
		if err := rows.Scan(&eUUID); err != nil {
			slog.Error("failed to scan: %s", err.Error())
			return "", err
		}
		res = eUUID
	}

	return res, nil
}

func (t *ExpenseRepository) CreateExpense(e *model.Expense) error {
	cmd := `insert into expense.expenses (expense_uuid, user_uuid, year, month, created_at, updated_at) values ($1, $2, $3, $4, $5, $6);`
	_, err := t.SqlHandler.DB.Exec(cmd, e.ExpenseUUID, e.UserUUID, e.Year, e.Month, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		slog.Error("failed to create expense:\n %s", err.Error())
		return err
	}
	return nil
}
