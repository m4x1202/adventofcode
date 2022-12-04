package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPuzzleInput = `A Y
B X
C Z`
)

var (
	prepareduzzleInput = []string{"A Y", "B X", "C Z"}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, prepareduzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(prepareduzzleInput)

	assert.EqualValues(t, 15, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(prepareduzzleInput)

	assert.EqualValues(t, 12, resPart2)
}
