package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"golang.org/x/oauth2"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/app/auth"
	"github.com/karan-khu/go-online-shop/internal/response"
)

type AuthorizationMiddleware interface {
	Authorizing(next echo.HandlerFunc) echo.HandlerFunc
}

type AuthorizationMiddlewareImpl struct {
	logger      *slog.Logger
	oauth2      *config.Oauth2Config
	env         *config.Env
	authUsecase auth.AuthGoogleUsecase
}

func NewAuthorizationMiddleware(logger *slog.Logger, conf *config.Config, authUsecase auth.AuthGoogleUsecase) AuthorizationMiddleware {
	return &AuthorizationMiddlewareImpl{
		logger:      logger,
		oauth2:      conf.Oauth2Config,
		env:         conf.Env,
		authUsecase: authUsecase,
	}
}

func (m *AuthorizationMiddlewareImpl) Authorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx *echo.Context) error {
		ctx := context.Background()

		token, err := m.getToken(pctx)
		if err != nil {
			m.logger.Error("Failed to get token", "error", err)
			return response.Unauthorized(pctx, errors.New("Unauthorized"))
		}

		if !token.Valid() {
			if token, err = m.tokenRefreshing(pctx, token); err != nil {
				pctx.Logger().Error("Token is not valid", "error", err)
				return response.Unauthorized(pctx, errors.New("Unauthorized"))
			}
		}

		client := m.oauth2.GoogleOAuth2Config.Client(ctx, token)

		userInfo, err := m.getUserInfo(client)
		if err != nil {
			m.logger.Error("Failed to get user info", "error", err)
			return response.Unauthorized(pctx, errors.New("Unauthorized"))
		}

		if !m.authUsecase.UserExists(userInfo.ID) {
			m.logger.Error("User does not exist", "userID", userInfo.ID)
			return response.Unauthorized(pctx, errors.New("User does not exist"))
		}

		pctx.Set("userId", userInfo.ID)
		pctx.Set("userEmail", userInfo.Email)

		return next(pctx)
	}
}

func (m *AuthorizationMiddlewareImpl) tokenRefreshing(pctx *echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()
	updateToken, err := m.oauth2.GoogleOAuth2Config.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, err
	}

	pctx.SetCookie(&http.Cookie{
		Name:     m.oauth2.AccessTokenKey,
		Value:    updateToken.AccessToken,
		Path:     "/",
		HttpOnly: true,
	})
	pctx.SetCookie(&http.Cookie{
		Name:     m.oauth2.RefreshTokenKey,
		Value:    updateToken.RefreshToken,
		Path:     "/",
		HttpOnly: true,
	})
	return updateToken, nil
}

func (m *AuthorizationMiddlewareImpl) getToken(pctx *echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Request().Cookie(m.oauth2.AccessTokenKey)
	if err == nil {
		refreshToken, err := pctx.Request().Cookie(m.oauth2.RefreshTokenKey)
		if err == nil {
			return &oauth2.Token{
				AccessToken:  accessToken.Value,
				RefreshToken: refreshToken.Value,
			}, nil
		}
	}

	return nil, errors.New("Unauthorized")
}

func (m *AuthorizationMiddlewareImpl) getUserInfo(client *http.Client) (*auth.UserCredential, error) {
	resp, err := client.Get(m.env.GOOGLE_USER_INFO_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo := new(auth.UserCredential)
	if err := json.Unmarshal(userInfoBytes, userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
