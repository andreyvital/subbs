package opensubtitles_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	data := make([]byte, opensubtitles.ChunkSize*2)

	copy(data, []byte("hahaha"))

	f, err := ioutil.TempFile(os.TempDir(), "opensubtitles")

	if err != nil {
		t.Fail()
		return
	}

	f.Write(data)
	f.Seek(0, 0)

	defer os.Remove(path.Join(os.TempDir(), f.Name()))

	var expected uint64 = 0x6168616a6168

	hash, err := opensubtitles.Hash(f)

	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, expected, hash)
}
