package opensubtitles_test

import (
	"io/ioutil"
	"testing"

	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/stretchr/testify/assert"
)

func TestSubtitleFileReader(t *testing.T) {
	file := opensubtitles.SubtitleFile{
		// `hahaha`
		Data: "H4sICBGQPlYAA3Rlc3QAy0jMAEIuAImEshYHAAAA",
	}

	reader, err := file.Reader()

	assert.Nil(t, err)
	assert.NotNil(t, reader)

	all, err := ioutil.ReadAll(reader)

	assert.Nil(t, err)
	assert.Equal(t, "hahaha\n", string(all))
}
