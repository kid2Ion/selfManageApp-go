package usecase

import "github.com/kid2Ion/selfManageApp-go/domain/repository"

type (
	HelloUsecase interface {
		Hello() string
	}
	hellousecase struct {
		helloRepo repository.HelloRepository
	}
)

// コンストラクタ
func NewHelloUsecase(helloRepository repository.HelloRepository) HelloUsecase {
	return &hellousecase{helloRepo: helloRepository}
}

func (t *hellousecase) Hello() string {
	return t.helloRepo.Hello()
}
