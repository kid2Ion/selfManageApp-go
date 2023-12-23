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
	// usersテーブル作成
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
		comment on constraint user_pkey on user_setting.users is 'PK制約です';
		`
	_, err = db.Exec(cmd)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return sqlHandler
}
