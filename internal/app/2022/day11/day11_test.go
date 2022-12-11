package day11

import (
	_ "embed"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTestInput string

//go:embed part1_test.txt
var part1TestInput string

//go:embed part2_test.txt
var part2TestInput string

var (
	preparedPuzzleInput = []Monkey{
		{0, 0, []uint{79, 98}, func(in uint) uint { return in * 19 }, 23, [2]uint8{2, 3}},
		{1, 0, []uint{54, 65, 75, 74}, func(in uint) uint { return in + 6 }, 19, [2]uint8{2, 0}},
	}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(inputTestInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, preparedPuzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	preparedInput := prepareInput(part1TestInput)
	resPart1 := part1Func(preparedInput)

	assert.EqualValues(t, 10605, resPart1)
}

func Test_Part2(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	preparedInput := prepareInput(part2TestInput)
	resPart2 := part2Func(preparedInput)

	assert.EqualValues(t, 2713310158, resPart2)
}
