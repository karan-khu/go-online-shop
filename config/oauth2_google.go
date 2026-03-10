package config

import (
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

var (
	googleOAuth2Config *oauth2.Config
	once               sync.Once
)

type Oauth2Config struct {
	GoogleOAuth2Config   *oauth2.Config
	AccessTokenKey       string
	RefreshTokenKey      string
	AppClientRedirectKey string
	StateCookieName      string
}

func NewGoogleOAuth2Config(env *Env) *Oauth2Config {
	once.Do(func() {
		googleOAuth2Config = &oauth2.Config{
			ClientID:     env.GOOGLE_CLIENT_ID,
			ClientSecret: env.GOOGLE_CLIENT_SECRET,
			RedirectURL:  env.GOOGLE_REDIRECT_URL,
			Scopes:       strings.Split(env.GOOGLE_SCOPES, ","),
			Endpoint: oauth2.Endpoint{
				AuthURL:       env.GOOGLE_AUTH_URL,
				TokenURL:      env.GOOGLE_TOKEN_URL,
				DeviceAuthURL: env.GOOGLE_DEVICE_AUTH_URL,
				AuthStyle:     oauth2.AuthStyleInParams,
			},
		}
	})
	return &Oauth2Config{
		GoogleOAuth2Config:   googleOAuth2Config,
		AccessTokenKey:       "ACCESS_TOKEN",
		RefreshTokenKey:      "REFRESH_TOKEN",
		AppClientRedirectKey: "APP_CLIENT_REDIRECT",
		StateCookieName:      "STATE",
	}
}
