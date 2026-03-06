package config

import (
	"strings"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOAuth2Config *oauth2.Config
	once               sync.Once
)

type Oauth2Config struct {
	GoogleOAuth2Config *oauth2.Config
	AccessTokenKey     string
	RefreshTokenKey    string
	StateCookieName    string
}

func NewGoogleOAuth2Config(env *Env) *Oauth2Config {
	once.Do(func() {
		googleOAuth2Config = &oauth2.Config{
			ClientID:     env.GOOGLE_CLIENT_ID,
			ClientSecret: env.GOOGLE_CLIENT_SECRET,
			RedirectURL:  env.GOOGLE_REDIRECT_URL,
			Scopes:       strings.Split(env.GOOGLE_SCOPES, ","),
			Endpoint:     google.Endpoint,
		}
	})
	return &Oauth2Config{
		GoogleOAuth2Config: googleOAuth2Config,
		AccessTokenKey:     "ACCESS_TOKEN",
		RefreshTokenKey:    "REFRESH_TOKEN",
		StateCookieName:    "STATE",
	}
}
