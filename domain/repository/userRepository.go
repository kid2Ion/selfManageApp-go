package repository

import "github.com/kid2Ion/selfManageApp-go/domain/model"

type (
	UserRepository interface {
		Get(u *model.User) (*model.User, error)
		Create(*model.User) error
	}
)
