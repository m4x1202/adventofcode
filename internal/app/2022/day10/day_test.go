package day10

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTestInput string

//go:embed part1_test.txt
var part1TestInput string

//go:embed part2_test.txt
var part2TestInput string

var (
	expectedPreparedInput = []Command{
		{1, 0},
		{2, 3},
		{2, -5},
	}
	expectedResultPart1 = 13140
	expectedResultPart2 = "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n"
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(inputTestInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, expectedPreparedInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	preparedInput := prepareInput(part1TestInput)
	resPart1 := part1Func(preparedInput)

	assert.EqualValues(t, expectedResultPart1, resPart1)
}

func Test_Part2(t *testing.T) {
	preparedInput := prepareInput(part2TestInput)
	resPart2 := part2Func(preparedInput)

	assert.EqualValues(t, expectedResultPart2, resPart2)
}
