package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPuzzleInput = `199
200
208
210
200
207
240
269
260
263`
)

var (
	prepareduzzleInput = []uint16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, prepareduzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(prepareduzzleInput)

	assert.EqualValues(t, 7, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(prepareduzzleInput)

	assert.EqualValues(t, 5, resPart2)
}
