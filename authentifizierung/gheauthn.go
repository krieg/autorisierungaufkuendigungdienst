package authentifizierung

import (
	"golang.org/x/oauth2"
	"net/http"
)

func Authn(token string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return tc
}