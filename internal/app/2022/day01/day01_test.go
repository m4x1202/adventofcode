package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPuzzleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
)

var (
	prepareduzzleInput = []int{24000, 11000, 10000, 6000, 4000}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, prepareduzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(prepareduzzleInput)

	assert.EqualValues(t, 24000, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(prepareduzzleInput)

	assert.EqualValues(t, 45000, resPart2)
}
