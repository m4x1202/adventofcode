package day12

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput = []string{""}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, preparedPuzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)
	resPart1 := part1Func(preparedInput)

	assert.EqualValues(t, 31, resPart1)
}

func Test_Part2(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)
	resPart2 := part2Func(preparedInput)

	assert.EqualValues(t, 29, resPart2)
}
