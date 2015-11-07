package opensubtitles

import (
	"strconv"
	"strings"

	"github.com/kolo/xmlrpc"
)

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

func (c *AuthenticatedClientState) Search(
	hash string,
	languages []string,
) (Subtitles, error) {
	if len(languages) == 0 {
		return nil, ErrNoLanguagesProvided
	}

	args := []interface{}{
		c.token,
		[]struct {
			Hash  string `xmlrpc:"moviehash"`
			Langs string `xmlrpc:"sublanguageid"`
		}{{
			hash,
			strings.Join(languages, ","),
		}},
	}

	res := struct {
		Data Subtitles `xmlrpc:"data"`
	}{}

	if err := c.client.Call("SearchSubtitles", args, &res); err != nil {
		if !strings.Contains(err.Error(), "type mismatch") {
			return nil, ErrUnableToSearchSubtitles
		}
	}

	return res.Data, nil
}

func (c *AuthenticatedClientState) Download(subtitle Subtitle) (SubtitleFile, error) {
	id, err := strconv.Atoi(subtitle.IDSubtitleFile)

	if err != nil {
		return SubtitleFile{}, err
	}

	args := []interface{}{c.token, []int{id}}

	res := struct {
		Status string         `xmlrpc:"status"`
		Data   []SubtitleFile `xmlrpc:"data"`
	}{}

	if err := c.client.Call("DownloadSubtitles", args, &res); err != nil {
		return SubtitleFile{}, err
	}

	if res.Status != "200 OK" {
		return SubtitleFile{}, ErrUnableToDownloadSubtitle
	}

	return res.Data[0], nil
}

func (c *AuthenticatedClientState) Authenticate(
	user,
	password string,
) (ClientState, error) {
	return nil, ErrAlreadyAuthenticated
}
