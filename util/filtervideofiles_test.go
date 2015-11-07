package util_test

import (
	"testing"

	"github.com/CentaurWarchief/subbs/util"
	"github.com/stretchr/testify/assert"
)

func TestFilterEmptyList(t *testing.T) {
	filtered := util.FilterVideoFiles([]string{})

	assert.Len(t, filtered, 0)
}

func TestFilterVideoFiles(t *testing.T) {
	filtered := util.FilterVideoFiles([]string{
		"1.mkv",
		"1.nfo",
		"1.avi",
		"2.srt",
		"",
		"",
		"",
		"3.mp4",
	})

	assert.Len(t, filtered, 3)
	assert.Contains(t, filtered, "1.mkv")
	assert.Contains(t, filtered, "1.avi")
	assert.Contains(t, filtered, "3.mp4")
}
