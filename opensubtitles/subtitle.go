package opensubtitles

import (
	"compress/gzip"
	"encoding/base64"
	"strings"
)

type Subtitle struct {
	IDSubtitleFile  string `xmlrpc:"IDSubtitleFile"`
	ISO639          string `xmlrpc:"ISO639"`
	LanguageName    string `xmlrpc:"LanguageName"`
	SubDownloadsCnt string `xmlrpc:"SubDownloadsCnt"`
	SubFeatured     string `xmlrpc:"SubFeatured"`
	SubFileName     string `xmlrpc:"SubFileName"`
	SubFormat       string `xmlrpc:"SubFormat"`
	SubHash         string `xmlrpc:"SubHash"`
}

type SubtitleFile struct {
	Id   string `xmlrpc:"idsubtitlefile"`
	Data string `xmlrpc:"data"`
}

func (s SubtitleFile) Reader() (*gzip.Reader, error) {
	return gzip.NewReader(base64.NewDecoder(
		base64.StdEncoding,
		strings.NewReader(s.Data),
	))
}
