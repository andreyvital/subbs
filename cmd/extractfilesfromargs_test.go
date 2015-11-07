package cmd_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/CentaurWarchief/subbs/cmd"
	"github.com/stretchr/testify/assert"
)

func TestExtractFilesFromArgs(t *testing.T) {
	assert.Len(t, cmd.ExtractFilesFromArgs([]string{}), 0)
	assert.Len(t, cmd.ExtractFilesFromArgs([]string{"/foo", "/bar", "/tmp/foo/bar"}), 0)

	var dir string
	var err error

	if dir, err = ioutil.TempDir(os.TempDir(), "ExtractFilesFromArgs"); err != nil {
		t.Fail()
		return
	}

	ioutil.TempFile(dir, "/1.mkv")
	ioutil.TempFile(dir, "/2.mkv")
	ioutil.TempFile(dir, "/3.mkv")

	defer os.RemoveAll(dir)

	assert.Len(t, cmd.ExtractFilesFromArgs([]string{dir}), 3)
}
