package day09

import (
	_ "embed"
	"testing"

	"github.com/m4x1202/adventofcode/pkg/physx"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput = []physx.Vector{
		{4, 0, 0},
		{0, 4, 0},
		{-3, 0, 0},
		{0, -1, 0},
		{4, 0, 0},
		{0, -1, 0},
		{-5, 0, 0},
		{2, 0, 0},
	}
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

	assert.EqualValues(t, 13, resPart1)
}

func Test_Part2(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)
	resPart2 := part2Func(preparedInput)

	assert.EqualValues(t, 36, resPart2)
}
