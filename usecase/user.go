package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/kid2Ion/selfManageApp-go/domain/model"
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
)

type (
	UserUsecase interface {
		Get(*UserReq) (*UserRes, error)
		GetByUserId(*UserReq) (*UserRes, error)
		Create(*UserReq) (*UserRes, error)
		Update(*UserReq) (*UserRes, error)
		Delete(*UserReq) error
	}
	userusecase struct {
		userRepo repository.UserRepository
	}
	UserReq struct {
		UserUUID string `json:"user_uuid"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		FUUID    string
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

func (t *userusecase) GetByUserId(r *UserReq) (*UserRes, error) {
	// TODO: 本当はRequestごとに構造体分けた方が良いかも
	user := &model.User{
		UserUUID: r.UserUUID,
	}
	res, err := t.userRepo.GetByUserId(user)
	if err != nil {
		return nil, err
	}

	return &UserRes{
		UserUUID: res.UserUUID,
		Email:    res.Email,
		Name:     res.Email,
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

func (t *userusecase) Update(r *UserReq) (*UserRes, error) {
	user := &model.User{
		UserUUID:  r.UserUUID,
		Email:     r.Email,
		Name:      r.Name,
		UpdatedAt: time.Now(),
	}
	err := t.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return &UserRes{
		UserUUID: user.UserUUID,
		Email:    user.Email,
		Name:     user.Name,
	}, nil
}

func (t *userusecase) Delete(r *UserReq) error {
	user := &model.User{
		UserUUID: r.UserUUID,
	}
	err := t.userRepo.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
