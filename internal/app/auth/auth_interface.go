package auth

import "github.com/labstack/echo/v5"

type AuthGoogleHandler interface {
	GoogleLogin(pctx *echo.Context) error
	GoogleLoginCallBack(pctx *echo.Context) error
	Logout(pctx *echo.Context) error

	UserAuthorizing(pctx *echo.Context, next echo.HandlerFunc) error
}

type AuthGoogleUsecase interface {
	UserLogin(credential *UserCredential) error
}
