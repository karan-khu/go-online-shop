package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/res"
)

var (
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type AuthGoogleHandlerImpl struct {
	logger            *slog.Logger
	env               *config.Env
	oauth2Config      *config.Oauth2Config
	authGoogleUsecase AuthGoogleUsecase
}

func NewAuthGoogleHandler(logger *slog.Logger, conf *config.Config, authGoogleUsecase AuthGoogleUsecase) AuthGoogleHandler {
	return &AuthGoogleHandlerImpl{
		logger:            logger,
		env:               conf.Env,
		authGoogleUsecase: authGoogleUsecase,
		oauth2Config:      conf.Oauth2Config,
	}
}

func (h *AuthGoogleHandlerImpl) GoogleLogin(pctx *echo.Context) error {
	state := randomState()

	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.StateCookieName,
		Value:    state,
		Path:     "/",
		HttpOnly: true,
	})

	return pctx.Redirect(http.StatusFound, h.oauth2Config.GoogleOAuth2Config.AuthCodeURL(state))
}

func (h *AuthGoogleHandlerImpl) GoogleLoginCallBack(pctx *echo.Context) error {
	ctx := context.Background()

	var errValidate error
	for range 3 {
		errValidate = h.callbackValidating(pctx)
		if errValidate == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if errValidate != nil {
		h.logger.Error("Failed to validate callback", "error", errValidate)
		return res.Unauthorized(pctx, errValidate)
	}

	token, err := h.oauth2Config.GoogleOAuth2Config.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		h.logger.Error("Failed to exchange token", "error", err)
		return res.Unauthorized(pctx, err)
	}

	client := h.oauth2Config.GoogleOAuth2Config.Client(ctx, token)

	userInfo, err := h.getUserInfo(client)
	if err != nil {
		h.logger.Error("Failed to get user info", "error", err)
		return res.Unauthorized(pctx, err)
	}
	h.logger.Info("User info", "user", userInfo)

	userReq := userInfo.ToUserLoginRequest()
	if err := h.authGoogleUsecase.UserLogin(userReq); err != nil {
		h.logger.Error("Failed to login user", "error", err)
		return res.Unauthorized(pctx, err)
	}

	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.AccessTokenKey,
		Value:    token.AccessToken,
		Path:     "/",
		HttpOnly: true,
	})
	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.RefreshTokenKey,
		Value:    token.RefreshToken,
		Path:     "/",
		HttpOnly: true,
	})
	return res.Success(pctx, "Logged in successfully", userInfo)
}

func (h *AuthGoogleHandlerImpl) Logout(pctx *echo.Context) error {
	accessToken, err := pctx.Request().Cookie(h.oauth2Config.AccessTokenKey)
	if err != nil {
		h.logger.Error("Access token cookie not found", "error", err)
		return res.BadRequest(pctx, err)
	}

	if err := h.revokeToken(accessToken.Value); err != nil {
		h.logger.Error("Failed to revoke token", "error", err)
		return res.InternalError(pctx, err)
	}

	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.AccessTokenKey,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.RefreshTokenKey,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	return res.Success(pctx, "Logged out successfully", nil)
}

func randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (h *AuthGoogleHandlerImpl) callbackValidating(pctx *echo.Context) error {
	state := pctx.QueryParam("state")

	stateFromCookie, err := pctx.Request().Cookie(h.oauth2Config.StateCookieName)
	if err != nil {
		h.logger.Error("State cookie not found", "error", err)
		return errors.New("Error: State cookie not found")
	}

	if state == "" || state != stateFromCookie.Value {
		h.logger.Error("Error: State mismatch")
		return errors.New("Error: State mismatch")
	}

	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.StateCookieName,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	return nil
}

func (h *AuthGoogleHandlerImpl) getUserInfo(client *http.Client) (*UserCredential, error) {
	resp, err := client.Get(h.env.GOOGLE_USER_INFO_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo := new(UserCredential)
	if err := json.Unmarshal(userInfoBytes, userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (h *AuthGoogleHandlerImpl) revokeToken(token string) error {
	revokeUrl := fmt.Sprintf("%s?token=%s", h.env.GOOGLE_REVOKE_TOKEN_URL, token)

	resp, err := http.Post(revokeUrl, "application/x-www-form-urlencoded", nil)
	if err != nil {
		h.logger.Error("Failed to revoke token", "error", err)
		return err
	}

	defer resp.Body.Close()

	return nil

}
