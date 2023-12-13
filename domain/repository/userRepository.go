package repository

import "github.com/kid2Ion/selfManageApp-go/domain/model"

type (
	UserRepository interface {
		Create(*model.User) error
	}
)
