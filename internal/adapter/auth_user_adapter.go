package adapter

import (
	"github.com/karan-khu/go-online-shop/internal/app/auth"
	"github.com/karan-khu/go-online-shop/internal/app/user"
)

type AuthUserAdapter interface {
	CreateUser(req *auth.UserLoginRequest) error
	UserExists(userID string) bool
}

type authUserAdapterImpl struct {
	userRepo user.UserRepository
}

func NewAuthUserAdapter(userRepo user.UserRepository) AuthUserAdapter {
	return &authUserAdapterImpl{userRepo: userRepo}
}

func (a *authUserAdapterImpl) CreateUser(req *auth.UserLoginRequest) error {
	entity := &user.UserEntity{
		UserId:   req.ID,
		FullName: req.FullName,
		Email:    req.Email,
		Picture:  req.Picture,
	}
	_, err := a.userRepo.Creating(entity)
	return err
}

func (a *authUserAdapterImpl) UserExists(userID string) bool {
	u, err := a.userRepo.FindById(userID)
	return err == nil && u != nil
}
