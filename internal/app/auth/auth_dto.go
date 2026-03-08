package auth

type UserLoginRequest struct {
	ID       string
	Email    string
	FullName string
	Picture  string
}
