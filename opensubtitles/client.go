package opensubtitles

import (
	"errors"

	"github.com/kolo/xmlrpc"
)

var (
	ErrAuthenticationFailure    = errors.New("Authentication failure")
	ErrNotAuthenticated         = errors.New("Not authenticated")
	ErrAlreadyAuthenticated     = errors.New("Already authenticated")
	ErrNoLanguagesProvided      = errors.New("No languages were provided")
	ErrUnableToSearchSubtitles  = errors.New("Unable to search for subtitles")
	ErrUnableToDownloadSubtitle = errors.New("Unable to download subtitle")
)

const (
	StatusSuccess      = "200 OK"
	DefaultOSUserAgent = "OSTestUserAgent"
	DefaultOSLanguage  = "en"
	OpenSubtitlesAPI   = "http://api.opensubtitles.org/xml-rpc"
)

type Client struct {
	state ClientState
}

func NewClient() *Client {
	client, _ := xmlrpc.NewClient(OpenSubtitlesAPI, nil)

	return &Client{
		NewUnauthenticatedClientState(*client),
	}
}

func (c *Client) Authenticate(user, password string) error {
	state, err := c.state.Authenticate(user, password)

	if err == nil {
		c.state = state
	}

	return err
}

func (c Client) Download(subtitle Subtitle) (SubtitleFile, error) {
	return c.state.Download(subtitle)
}

func (c Client) Search(hash string, languages []string) (Subtitles, error) {
	return c.state.Search(hash, languages)
}
