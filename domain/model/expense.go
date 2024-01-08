package model

import "time"

type (
	Expense struct {
		ExpenseUUID string
		UserUUID    string
		Year        int
		Month       int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Income struct {
		IncomeUUID  string
		ExpenseUUID string
		Amount      int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Outcome struct {
		OutcomeUUID string
		ExpenseUUID string
		Title       string
		Day         int
		Amount      int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
