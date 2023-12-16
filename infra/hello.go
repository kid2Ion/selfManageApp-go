package infra

import (
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
)

type (
	HelloRepository struct {
		SqlHandler
	}
)

// コンストラクタ
func NewHelloRepository(sqlHandler SqlHandler) repository.HelloRepository {
	return &HelloRepository{SqlHandler: sqlHandler}
}

func (t *HelloRepository) Hello() string {
	return "hello world"
}
