package opensubtitles

import (
	"compress/gzip"
	"encoding/base64"
	"strings"
)

type Subtitle struct {
	IDSubtitle       string `xmlrpc:"IDSubtitle"`
	IDSubtitleFile   string `xmlrpc:"IDSubtitleFile"`
	ISO639           string `xmlrpc:"ISO639"`
	LanguageName     string `xmlrpc:"LanguageName"`
	MovieName        string `xmlrpc:"MovieName"`
	MovieNameEng     string `xmlrpc:"MovieNameEng"`
	MovieReleaseName string `xmlrpc:"MovieReleaseName"`
	SubDownloadLink  string `xmlrpc:"SubDownloadLink"`
	SubDownloadsCnt  string `xmlrpc:"SubDownloadsCnt"`
	SubFeatured      string `xmlrpc:"SubFeatured"`
	SubFileName      string `xmlrpc:"SubFileName"`
	SubFormat        string `xmlrpc:"SubFormat"`
	SubHash          string `xmlrpc:"SubHash"`
}

type SubtitleFile struct {
	Id     string `xmlrpc:"idsubtitlefile"`
	Data   string `xmlrpc:"data"`
	reader *gzip.Reader
}

func (s *SubtitleFile) Reader() (*gzip.Reader, error) {
	if s.reader != nil {
		return s.reader, nil
	}

	decoder := base64.NewDecoder(
		base64.StdEncoding,
		strings.NewReader(s.Data),
	)

	reader, err := gzip.NewReader(decoder)

	s.reader = reader

	return reader, err
}
