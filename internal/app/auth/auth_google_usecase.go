package auth

type AuthGoogleUsecaseImpl struct {
	userCreator UserCreatorAdapter
}

func NewAuthGoogleUsecase(userCreator UserCreatorAdapter) AuthGoogleUsecase {
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

func (u *AuthGoogleUsecaseImpl) UserExists(userID string) bool {
	return u.userCreator.UserExists(userID)
}
