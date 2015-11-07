package opensubtitles_test

import (
	"testing"

	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/stretchr/testify/assert"
)

func TestGetFirst(t *testing.T) {
	subtitles := opensubtitles.Subtitles{}

	assert.Len(t, subtitles, 0)
	assert.Nil(t, subtitles.First())
}

func TestSubtitlesFirst(t *testing.T) {
	subtitles := opensubtitles.Subtitles{
		opensubtitles.Subtitle{},
		opensubtitles.Subtitle{},
	}

	assert.Len(t, subtitles, 2)
	assert.NotNil(t, subtitles.First(), 0)
}
