package fs_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/CentaurWarchief/subbs/fs"
	"github.com/stretchr/testify/assert"
)

func TestReadEmptyDirectory(t *testing.T) {
	var dir string
	var err error

	if dir, err = ioutil.TempDir(os.TempDir(), "TestReadEmptyDirectory"); err != nil {
		t.Fail()
		return
	}

	files := fs.ReadDir(dir)

	assert.Len(t, files, 0)
}

func TestReadDir(t *testing.T) {
	var dir string
	var err error

	if dir, err = ioutil.TempDir(os.TempDir(), "TestReadDir"); err != nil {
		t.Fail()
		return
	}

	ioutil.TempFile(dir, "1.mkv")
	ioutil.TempFile(dir, "2.mp4")
	ioutil.TempFile(dir, "3.avi")

	files := fs.ReadDir(dir)

	assert.Len(t, files, 3)
}
