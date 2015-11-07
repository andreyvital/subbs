package opensubtitles

import "github.com/kolo/xmlrpc"

type UnauthenticatedClientState struct {
	client xmlrpc.Client
}

func NewUnauthenticatedClientState(
	client xmlrpc.Client,
) *UnauthenticatedClientState {
	return &UnauthenticatedClientState{
		client,
	}
}

func (c *UnauthenticatedClientState) Authenticate(
	user,
	password string,
) (ClientState, error) {
	args := []interface{}{
		user,
		password,
		DefaultOSLanguage,
		DefaultOSUserAgent,
	}

	res := struct {
		Status string `xmlrpc:"status"`
		Token  string `xmlrpc:"token"`
	}{}

	if err := c.client.Call("LogIn", args, &res); err != nil {
		return nil, ErrAuthenticationFailure
	}

	if res.Status != StatusSuccess {
		return nil, ErrAuthenticationFailure
	}

	return NewAuthenticatedClientState(c.client, res.Token), nil
}

func (c *UnauthenticatedClientState) Download(
	subtitle Subtitle,
) (SubtitleFile, error) {
	return SubtitleFile{}, ErrNotAuthenticated
}

func (c *UnauthenticatedClientState) Search(
	hash string,
	languages []string,
) (Subtitles, error) {
	return nil, ErrNotAuthenticated
}
