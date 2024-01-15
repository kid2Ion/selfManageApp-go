package infra

import (
	"time"

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

func (t *ExpenseRepository) GetIncome(expenseUUID string) (*model.Income, error) {
	q := `
select * from expense.incomes i
where i.expense_uuid = $1
;`

	rows, err := t.SqlHandler.DB.Query(q, expenseUUID)
	if err != nil {
		slog.Error("failed to fetch from db: %s", err.Error())
		return nil, err
	}
	res := new(model.Income)
	defer rows.Close()
	if rows.Next() {
		var rslt struct {
			IncomeUUID  string    `db:"income_uuid"`
			ExpenseUUID string    `db:"expense_uuid"`
			Amount      int       `db:"amount"`
			CreatedAt   time.Time `db:"created_at"`
			UpdatedAt   time.Time `db:"updated_at"`
		}
		if err := rows.Scan(&rslt.IncomeUUID, &rslt.ExpenseUUID, &rslt.Amount, &rslt.CreatedAt, &rslt.UpdatedAt); err != nil {
			slog.Error("failed to scan: %s", err.Error())
			return nil, err
		}
		res.IncomeUUID = rslt.IncomeUUID
		res.ExpenseUUID = rslt.ExpenseUUID
		res.Amount = rslt.Amount
		res.CreatedAt = rslt.CreatedAt
		res.UpdatedAt = rslt.UpdatedAt
	}
	return res, nil
}

func (t *ExpenseRepository) CreateOutcome(o *model.Outcome) error {
	cmd := `insert into expense.outcomes (outcome_uuid, expense_uuid, amount, title, day, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7);`
	_, err := t.SqlHandler.DB.Exec(cmd, o.OutcomeUUID, o.ExpenseUUID, o.Amount, o.Title, o.Day, o.CreatedAt, o.UpdatedAt)
	if err != nil {
		slog.Error("failed to create outcome:\n %s", err.Error())
		return err
	}
	return nil
}

func (t *ExpenseRepository) GetOutcomes(expenseUUID string) ([]model.Outcome, error) {
	q := `
select * from expense.outcomes o
where o.expense_uuid = $1
order by o.day
;`

	rows, err := t.SqlHandler.DB.Query(q, expenseUUID)
	if err != nil {
		slog.Error("failed to fetch from db: %s", err.Error())
		return nil, err
	}
	res := []model.Outcome{}
	defer rows.Close()
	var rslt struct {
		OutcomeUUID string    `db:"outcome_uuid"`
		ExpenseUUID string    `db:"expense_uuid"`
		Amount      int       `db:"amount"`
		Title       string    `db:"title"`
		Day         int       `db:"day"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
	}
	for rows.Next() {
		var outcome model.Outcome
		if err := rows.Scan(&rslt.OutcomeUUID, &rslt.ExpenseUUID, &rslt.Amount, &rslt.Title, &rslt.Day, &rslt.CreatedAt, &rslt.UpdatedAt); err != nil {
			slog.Error("failed to scan: %s", err.Error())
			return nil, err
		}
		outcome.OutcomeUUID = rslt.OutcomeUUID
		outcome.ExpenseUUID = rslt.ExpenseUUID
		outcome.Amount = rslt.Amount
		outcome.Title = rslt.Title
		outcome.Day = rslt.Day
		outcome.CreatedAt = rslt.CreatedAt
		outcome.UpdatedAt = rslt.UpdatedAt
		res = append(res, outcome)
	}
	return res, nil
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

// TODO: domain間のデータ取得は止める
func (t *ExpenseRepository) GetUserUUIDByFUUID(fUUID string) (string, error) {
	q := `
select u.user_uuid from user_setting.users u
where u.firebase_uuid = $1
;`

	rows, err := t.SqlHandler.DB.Query(q, fUUID)
	if err != nil {
		slog.Error("failed to fetch from db: %s", err.Error())
		return "", err
	}

	defer rows.Close()
	res := ""
	if rows.Next() {
		var userUUID string
		if err := rows.Scan(&userUUID); err != nil {
			slog.Error("failed to scan: %s", err.Error())
			return "", err
		}
		res = userUUID
	}

	return res, nil
}
