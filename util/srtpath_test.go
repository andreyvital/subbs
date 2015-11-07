package util_test

import (
	"testing"

	"github.com/CentaurWarchief/subbs/util"
	"github.com/stretchr/testify/assert"
)

func TestSrtPath(t *testing.T) {
	for _, pair := range [][]string{
		[]string{"1.srt", "1.mkv"},
		[]string{"2.srt", "2.mp4"},
		[]string{"/path/to/video.srt", "/path/to/video.mkv"},
		[]string{"/path/video.mkv.srt", "/path/video.mkv.mkv"},
	} {
		assert.Equal(t, pair[0], util.SrtPath(pair[1]))
	}
}
