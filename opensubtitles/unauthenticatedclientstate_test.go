package opensubtitles_test

import (
	"testing"

	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/stretchr/testify/assert"
)

func TestUnauthenticatedStateSearch(t *testing.T) {
	state := &opensubtitles.UnauthenticatedClientState{}

	_, err := state.Search("hash", []string{""})

	assert.Equal(t, opensubtitles.ErrNotAuthenticated, err)
}
