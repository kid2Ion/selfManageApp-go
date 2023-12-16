package infra

import (
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

func (t *UserRepository) Create(u *model.User) error {
	cmd := "insert into users (user_uuid, mail, name, firebase_uuid, created_at, updated_at) values ($1, $2, $3, $4, $5, $6);"
	_, err := t.SqlHandler.DB.Exec(cmd, u.UserUUID, u.Email, u.Name, u.FirebaseUUID, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		slog.Error("failed to exec:\n %s", err.Error())
		return err
	}
	return nil
}
