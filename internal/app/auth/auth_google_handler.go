package auth

import (
	"log/slog"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/karan-khu/go-online-shop/config"
)

var (
	googleOAuth2Config *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "ACCESS_TOKEN"
	refreshTokenCookieName = "REFRESH_TOKEN"
	stateCookieName        = "STATE"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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
	panic("unimplemented")
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
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
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
