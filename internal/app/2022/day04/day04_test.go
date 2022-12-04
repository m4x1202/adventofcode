package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPuzzleInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
)

var (
	prepareduzzleInput = []Pair{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, prepareduzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(prepareduzzleInput)

	assert.EqualValues(t, 2, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(prepareduzzleInput)

	assert.EqualValues(t, 4, resPart2)
}
