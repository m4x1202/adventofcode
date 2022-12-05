package day05

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput1 = CargoBay{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	preparedPuzzleInput2 = []ProcedureStep{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput1, preparedInput2 := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput1) {
		assert.Equal(t, preparedPuzzleInput1, preparedInput1)
	}

	if assert.NotEmpty(t, preparedInput2) {
		assert.Equal(t, preparedPuzzleInput2, preparedInput2)
	}
}

func Test_Part1(t *testing.T) {
	preparedPuzzleInput1p1 := make(CargoBay, len(preparedPuzzleInput1))
	copy(preparedPuzzleInput1p1, preparedPuzzleInput1)
	resPart1 := part1Func(preparedPuzzleInput1p1, preparedPuzzleInput2)

	assert.EqualValues(t, "CMZ", resPart1)
}

func Test_Part2(t *testing.T) {
	preparedPuzzleInput1p2 := make(CargoBay, len(preparedPuzzleInput1))
	copy(preparedPuzzleInput1p2, preparedPuzzleInput1)
	resPart2 := part2Func(preparedPuzzleInput1p2, preparedPuzzleInput2)

	assert.EqualValues(t, "MCD", resPart2)
}
