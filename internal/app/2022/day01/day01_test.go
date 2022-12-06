package day01

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput = []int{24000, 11000, 10000, 6000, 4000}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, preparedPuzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(preparedPuzzleInput)

	assert.EqualValues(t, 24000, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(preparedPuzzleInput)

	assert.EqualValues(t, 45000, resPart2)
}
