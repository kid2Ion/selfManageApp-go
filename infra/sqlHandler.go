package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type SqlHandler struct {
	DB *sql.DB
}

func NewSqlHandler() *SqlHandler {
	// 環境変数適応
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// usersテーブル作成
	cmd := fmt.Sprintf(`
		create table if not exists %s (
			uuid uuid not null,
			mail text not null,
			name text not null,
			created_at timestamp,
			updated_at timestamp
		)`, "users")
	_, err = db.Exec(cmd)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return sqlHandler
}
