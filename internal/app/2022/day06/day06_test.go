package day06

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput = []rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(strings.TrimSpace(testPuzzleInput))

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, preparedPuzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(preparedPuzzleInput)

	assert.EqualValues(t, 7, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(preparedPuzzleInput)

	assert.EqualValues(t, 19, resPart2)
}
