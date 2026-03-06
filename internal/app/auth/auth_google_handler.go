package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/res"
)

var (
	googleOAuth2Config *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "ACCESS_TOKEN"
	refreshTokenCookieName = "REFRESH_TOKEN"
	stateCookieName        = "STATE"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

type AuthGoogleHandlerImpl struct {
	logger            *slog.Logger
	env               *config.Env
	authGoogleUsecase AuthGoogleUsecase
}

func NewAuthGoogleHandler(logger *slog.Logger, conf *config.Config, authGoogleUsecase AuthGoogleUsecase) AuthGoogleHandler {
	once.Do(func() {
		scopes := strings.Split(conf.Env.GOOGLE_SCOPES, ",")
		googleOAuth2Config = &oauth2.Config{
			ClientID:     conf.Env.GOOGLE_CLIENT_ID,
			ClientSecret: conf.Env.GOOGLE_CLIENT_SECRET,
			RedirectURL:  conf.Env.GOOGLE_REDIRECT_URL,
			Scopes:       scopes,
			Endpoint:     google.Endpoint,
		}
	})
	return &AuthGoogleHandlerImpl{logger, conf.Env, authGoogleUsecase}
}

func (h *AuthGoogleHandlerImpl) GoogleLogin(pctx *echo.Context) error {
	state := randomState()

	h.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, googleOAuth2Config.AuthCodeURL(state))
}

func (h *AuthGoogleHandlerImpl) GoogleLoginCallBack(pctx *echo.Context) error {
	ctx := context.Background()

	errValidate := errors.New("")
	const MAX_ATTEMPTS = 3
	const DELAY = 1 * time.Second
	for attempt := 0; attempt < MAX_ATTEMPTS; attempt++ {
		errValidate = (h.callbackValidating(pctx))
		if errValidate == nil {
			break
		}
		time.Sleep(DELAY)
	}
	if errValidate != nil {
		h.logger.Error("Failed to validate callback", "error", errValidate)
		return res.Unauthorized(pctx, errValidate)
	}

	token, err := googleOAuth2Config.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		h.logger.Error("Failed to exchange token", "error", err)
		return res.Unauthorized(pctx, err)
	}

	client := googleOAuth2Config.Client(ctx, token)

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

	h.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	h.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)

	return res.Success(pctx, "Logged in successfully", nil)
}

func (h *AuthGoogleHandlerImpl) Logout(pctx *echo.Context) error {
	panic("unimplemented")
}

func (h *AuthGoogleHandlerImpl) UserAuthorizing(pctx *echo.Context, next echo.HandlerFunc) error {
	panic("unimplemented")
}

func randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (h *AuthGoogleHandlerImpl) setCookie(pctx *echo.Context, name, value string) {
	isProduction := h.env.GO_ENV == "production"
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   isProduction,
	}
	pctx.SetCookie(&cookie)
}

func (h *AuthGoogleHandlerImpl) removeCookie(pctx *echo.Context, name string) {
	cookie := http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	pctx.SetCookie(&cookie)
}

func (h *AuthGoogleHandlerImpl) setSameSiteCookie(pctx *echo.Context, name, value string) {
	isProduction := h.env.GO_ENV == "production"
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   isProduction,
	}
	pctx.SetCookie(&cookie)
}

func (h *AuthGoogleHandlerImpl) callbackValidating(pctx *echo.Context) error {
	state := pctx.QueryParam("state")

	stateFromCookie, err := pctx.Request().Cookie(stateCookieName)
	if err != nil {
		h.logger.Error("State cookie not found", "error", err)
		return errors.New("Error: State cookie not found")
	}

	if state == "" || state != stateFromCookie.Value {
		h.logger.Error("Error: State mismatch")
		return errors.New("Error: State mismatch")
	}

	h.removeCookie(pctx, stateCookieName)

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
