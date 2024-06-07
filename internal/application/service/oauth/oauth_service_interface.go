package oauth

import (
	"context"
	"golang.org/x/oauth2"
)

type OAuthProvider interface {
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
}
