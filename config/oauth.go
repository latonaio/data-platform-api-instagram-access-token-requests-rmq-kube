package config

import (
	"os"
)

type OAuth struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}

func newOAuth() *OAuth {
	return &OAuth{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectUrl:  os.Getenv("OAUTH_INSTAGRAM_REDIRECT_URL"),
	}
}
