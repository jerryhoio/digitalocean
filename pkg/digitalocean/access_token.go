package digitalocean

import (
	"golang.org/x/oauth2"
)

type AccessToken struct {
	AccessToken string
}

func (t *AccessToken) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
