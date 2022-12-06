package day02

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput = []string{"A Y", "B X", "C Z"}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, preparedPuzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(preparedPuzzleInput)

	assert.EqualValues(t, 15, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(preparedPuzzleInput)

	assert.EqualValues(t, 12, resPart2)
}
