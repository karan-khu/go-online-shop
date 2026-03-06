package auth

type AuthGoogleUsecaseImpl struct {
	userCreator UserCreator
}

func NewAuthGoogleUsecase(userCreator UserCreator) AuthGoogleUsecase {
	return &AuthGoogleUsecaseImpl{
		userCreator: userCreator,
	}
}

func (u *AuthGoogleUsecaseImpl) UserLogin(userReq *UserLoginRequest) error {
	if !u.userCreator.UserExists(userReq.ID) {
		if err := u.userCreator.CreateUser(userReq); err != nil {
			return err
		}
	}
	return nil
}
