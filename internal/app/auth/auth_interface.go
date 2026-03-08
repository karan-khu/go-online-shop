package auth

import "github.com/labstack/echo/v5"

type AuthGoogleHandler interface {
	GoogleLogin(pctx *echo.Context) error
	GoogleLoginCallBack(pctx *echo.Context) error
	Logout(pctx *echo.Context) error
}

type AuthGoogleUsecase interface {
	UserLogin(userReq *UserLoginRequest) error
	UserExists(userID string) bool
}

type UserCreatorAdapter interface {
	CreateUser(req *UserLoginRequest) error
	UserExists(userID string) bool
}
