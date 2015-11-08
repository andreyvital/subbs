package opensubtitles_test

import (
	"testing"

	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	state := &opensubtitles.AuthenticatedClientState{}

	assert.NotNil(t, state)

	_, err := state.Authenticate("", "")

	assert.Equal(t, opensubtitles.ErrAlreadyAuthenticated, err)
}
