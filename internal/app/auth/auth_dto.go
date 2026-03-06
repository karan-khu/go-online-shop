package auth

type UserLoginRequest struct {
	ID       string
	Email    string
	FullName string
	Picture  string
}

func (u *UserLoginRequest) ToUserEntity() *UserEntity {
	return &UserEntity{
		UserId:   u.ID,
		Email:    u.Email,
		FullName: u.FullName,
		Picture:  u.Picture,
	}
}
