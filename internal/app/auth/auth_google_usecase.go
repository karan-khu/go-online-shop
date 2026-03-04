package auth

import "github.com/karan-khu/go-online-shop/internal/app/user"

type AuthGoogleUsecaseImpl struct {
	userRepo user.UserRepository
}

func NewAuthGoogleUsecase(userRepo user.UserRepository) AuthGoogleUsecase {
	return &AuthGoogleUsecaseImpl{
		userRepo: userRepo,
	}
}

func (u *AuthGoogleUsecaseImpl) UserLogin(credential *UserCredential) error {
	panic("unimplemented")
}
