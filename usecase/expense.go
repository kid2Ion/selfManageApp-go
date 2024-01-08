package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/kid2Ion/selfManageApp-go/domain/model"
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
)

type (
	ExpenseUsecase interface {
		CreateIncome(i *IncomeReq) error
	}
	expenseUsecase struct {
		expenseRepo repository.ExpenseRepository
	}
	IncomeReq struct {
		IncomeUUID string `json:"income_uuid"`
		UserUUID   string `json:"user_uuid"`
		Year       int    `json:"year"`
		Month      int    `json:"month"`
		Amount     int    `json:"amount"`
	}
)

func NewExpenseUsecase(
	expenseRepo repository.ExpenseRepository,
) ExpenseUsecase {
	return &expenseUsecase{expenseRepo: expenseRepo}
}

func (t *expenseUsecase) CreateIncome(i *IncomeReq) error {
	now := time.Now()
	e := &model.Expense{
		UserUUID: i.UserUUID,
		Year:     i.Year,
		Month:    i.Month,
	}
	// expense探す
	eUUID, err := t.expenseRepo.GetExpenseUUID(e)
	if err != nil {
		return err
	}
	// expenseが存在しない場合作成
	if eUUID == "" {
		e.ExpenseUUID = uuid.New().String()
		e.CreatedAt = now
		e.UpdatedAt = now
		err = t.expenseRepo.CreateExpense(e)
		if err != nil {
			return err
		}
		eUUID = e.ExpenseUUID
	}
	income := &model.Income{
		IncomeUUID:  uuid.New().String(),
		ExpenseUUID: eUUID,
		Amount:      i.Amount,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return t.expenseRepo.CreateIncome(income)
}
