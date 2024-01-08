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
		CreateOutcome(o *OutcomeReq) error
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
	OutcomeReq struct {
		OutcomeUUID string `json:"outcome_uuid"`
		UserUUID    string `json:"user_uuid"`
		Year        int    `json:"year"`
		Month       int    `json:"month"`
		Amount      int    `json:"amount"`
		Title       string `json:"title"`
		Day         int    `json:"day"`
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
	//todo : 既存にincomeがないか（今はあればDBのユニーク制約に引っかかる）
	income := &model.Income{
		IncomeUUID:  uuid.New().String(),
		ExpenseUUID: eUUID,
		Amount:      i.Amount,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return t.expenseRepo.CreateIncome(income)
}

func (t *expenseUsecase) CreateOutcome(o *OutcomeReq) error {
	now := time.Now()
	e := &model.Expense{
		UserUUID: o.UserUUID,
		Year:     o.Year,
		Month:    o.Month,
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
	outcome := &model.Outcome{
		OutcomeUUID: uuid.New().String(),
		ExpenseUUID: eUUID,
		Title:       o.Title,
		Day:         o.Day,
		Amount:      o.Amount,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return t.expenseRepo.CreateOutcome(outcome)
}
