package infra

import (
	"time"

	"github.com/kid2Ion/selfManageApp-go/domain/model"
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
	"golang.org/x/exp/slog"
)

type (
	UserRepository struct {
		SqlHandler
	}
)

// コンストラクタ
func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	return &UserRepository{SqlHandler: sqlHandler}
}

func (t *UserRepository) Get(u *model.User) (*model.User, error) {
	q := "select * from users u where u.firebase_uuid = $1;"
	rows, err := t.SqlHandler.DB.Query(q, u.FirebaseUUID)
	if err != nil {
		slog.Error("failed to fetch from db: %s", err.Error())
		return nil, err
	}

	defer rows.Close()

	var rslt struct {
		UserUUID     string    `db:"user_uuid"`
		Email        string    `db:"email"`
		Name         string    `db:"name"`
		FirebaseUUID string    `db:"firebase_uuid"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
	}
	var res model.User
	if rows.Next() {
		if err := rows.Scan(&rslt.UserUUID, &rslt.Email, &rslt.Name, &rslt.FirebaseUUID, &rslt.CreatedAt, &rslt.UpdatedAt); err != nil {
			slog.Error("failed to scan: %s", err.Error())
			return nil, err
		}
		res.UserUUID = rslt.UserUUID
		res.Email = rslt.Email
		res.Name = rslt.Name
		res.FirebaseUUID = rslt.FirebaseUUID
		res.CreatedAt = rslt.CreatedAt
		res.UpdatedAt = rslt.UpdatedAt
	}
	err = rows.Err()
	if err != nil {
		slog.Error("failed to fetch from db: %s", err.Error())
		return nil, err
	}

	return &res, nil
}

func (t *UserRepository) Create(u *model.User) error {
	cmd := "insert into users (user_uuid, mail, name, firebase_uuid, created_at, updated_at) values ($1, $2, $3, $4, $5, $6);"
	_, err := t.SqlHandler.DB.Exec(cmd, u.UserUUID, u.Email, u.Name, u.FirebaseUUID, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		slog.Error("failed to exec:\n %s", err.Error())
		return err
	}
	return nil
}
