package infra

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type SqlHandler struct {
	DB *sql.DB
}

// todo migrationファイル

func NewSqlHandler() *SqlHandler {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	cmd := `
		-- userスキーマ作成
		create schema if not exists user_setting;
		comment on schema user_setting is 'user domainのデータを記録するスキーマです';
		-- usersテーブル作成
		create table if not exists user_setting.users (
			user_uuid uuid not null,
			mail text not null,
			name text not null,
			firebase_uuid text not null,
			created_at timestamp,
			updated_at timestamp,
			constraint user_pkey primary key (user_uuid)
		);
		comment on table user_setting.users is 'user情報を管理するテーブルです';
		comment on column user_setting.users.user_uuid is 'レコードを一意に識別するIDです';
		comment on column user_setting.users.mail is 'メールアドレスです';
		comment on column user_setting.users.name is '名前です';
		comment on column user_setting.users.firebase_uuid is 'firebase上でuserを一意に識別するUIDです';
		comment on column user_setting.users.created_at is 'レコードを初めに登録した日です';
		comment on column user_setting.users.updated_at is 'レコードを更新した日です';
		-- expenseスキーマ作成
		create schema if not exists expense;
		comment on schema expense is 'expense domainのデータを記録するスキーマです';
		-- expensesテーブル作成
		create table if not exists expense.expenses (
			expense_uuid uuid not null,
			user_uuid uuid not null,
			year int not null,
			month int not null,
			created_at timestamp,
			updated_at timestamp,
			constraint expense_pkey primary key (expense_uuid)
		);
		create unique index if not exists idx_expense_year_month_user_nkey on expense.expenses using btree (user_uuid, year, month);
		comment on table expense.expenses is '月の収支情報を管理するテーブルです';
		comment on column expense.expenses.expense_uuid is 'レコードを一意に識別するIDです';
		comment on column expense.expenses.user_uuid is 'userUUIDです';
		comment on column expense.expenses.year is '収支の対象年です';
		comment on column expense.expenses.month is '収支の対象月です';
		comment on column expense.expenses.created_at is 'レコードを初めに登録した日です';
		comment on column expense.expenses.updated_at is 'レコードを更新した日です';
		comment on index expense.idx_expense_year_month_user_nkey is 'レコードを一意に特定する自然キーです。userとyearとmonthで一意に特定します';
		-- incomesテーブル作成
		create table if not exists expense.incomes (
			income_uuid uuid not null,
			expense_uuid uuid not null,
			amount int not null,
			created_at timestamp,
			updated_at timestamp,
			constraint income_pkey primary key (income_uuid),
			constraint incomes_expenses_fkey foreign key (expense_uuid) references expense.expenses(expense_uuid) on delete cascade
		);
		create unique index if not exists idx_income_expense_uuid_nkey on expense.incomes using btree (expense_uuid);
		comment on table expense.incomes is '月の収入情報を管理するテーブルです';
		comment on column expense.incomes.income_uuid is 'レコードを一意に識別するIDです';
		comment on column expense.incomes.expense_uuid is 'expenseUUIDです';
		comment on column expense.incomes.amount is '収入量です';
		comment on column expense.incomes.created_at is 'レコードを初めに登録した日です';
		comment on column expense.incomes.updated_at is 'レコードを更新した日です';
		comment on index expense.idx_income_expense_uuid_nkey is 'レコードを一意に特定する自然キーです。expenses1レコードにつきincomesのレコードも1つです';
		-- outcomesテーブル作成
		create table if not exists expense.outcomes (
			outcome_uuid uuid not null,
			expense_uuid uuid not null,
			amount int not null,
			title text,
			day int,
			created_at timestamp,
			updated_at timestamp,
			constraint outcome_pkey primary key (outcome_uuid),
			constraint outcomes_expenses_fkey foreign key (expense_uuid) references expense.expenses(expense_uuid) on delete cascade
		);
		comment on table expense.outcomes is '月の支出情報を管理するテーブルです';
		comment on column expense.outcomes.outcome_uuid is 'レコードを一意に識別するIDです';
		comment on column expense.outcomes.expense_uuid is 'expenseUUIDです';
		comment on column expense.outcomes.amount is '収入量です';
		comment on column expense.outcomes.title is '収支タイトルです';
		comment on column expense.outcomes.day is '収支の日付です';
		comment on column expense.outcomes.created_at is 'レコードを初めに登録した日です';
		comment on column expense.outcomes.updated_at is 'レコードを更新した日です';
		`
	_, err = db.Exec(cmd)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return sqlHandler
}
