package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/kid2Ion/selfManageApp-go/domain/model"
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
)

type (
	UserUsecase interface {
		Create(*UserReq) (*UserRes, error)
		Get(*UserReq) (*UserRes, error)
	}
	userusecase struct {
		userRepo repository.UserRepository
	}
	UserReq struct {
		FUUID string
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	UserRes struct {
		UserUUID string
		Email    string
		Name     string
	}
)

// コンストラクタ
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userusecase{userRepo: userRepository}
}

func (t *userusecase) Get(r *UserReq) (*UserRes, error) {
	user := &model.User{
		FirebaseUUID: r.FUUID,
	}
	res, err := t.userRepo.Get(user)
	if err != nil {
		return nil, err
	}

	return &UserRes{
		UserUUID: res.UserUUID,
		Email:    res.Email,
		Name:     res.Name,
	}, nil
}

func (t *userusecase) Create(r *UserReq) (*UserRes, error) {
	user := &model.User{
		UserUUID:     uuid.New().String(),
		Email:        r.Email,
		Name:         r.Name,
		FirebaseUUID: r.FUUID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := t.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &UserRes{
		UserUUID: user.UserUUID,
		Email:    r.Email,
		Name:     r.Name,
	}, nil
}
