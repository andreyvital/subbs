package opensubtitles

import "github.com/kolo/xmlrpc"

type AuthenticatedClientState struct {
	client xmlrpc.Client
	token  string
}

func NewAuthenticatedClientState(
	client xmlrpc.Client,
	token string,
) *AuthenticatedClientState {
	return &AuthenticatedClientState{
		client,
		token,
	}
}

func (c *AuthenticatedClientState) Authenticate(
	user,
	password string,
) (ClientState, error) {
	return nil, ErrAlreadyAuthenticated
}

func (c *AuthenticatedClientState) Search(
	hash string,
	languages []string,
) (Subtitles, error) {
	if len(languages) == 0 {
		return nil, ErrNoLanguagesProvided
	}

	return nil, nil
}
