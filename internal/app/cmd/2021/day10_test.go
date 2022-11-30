package cmd2021

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func Test_ParseChunk(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	length, _ := ParseChunk("[]")
	assert.Equal(t, 2, length, "Expected length of read chunk data is 2")
}

func Test_ParseIncompleteChunk(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	length, _ := ParseChunk("[")
	assert.Equal(t, 1, length, "Expected length of read chunk data is 1")
}
