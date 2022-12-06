package day15

import (
	_ "embed"
)

//go:embed input_test.txt
var testPuzzleInput string

// func Test_PrepareInput(t *testing.T) {
// 	preparedInput := prepareInput(testPuzzleInput)

// 	if assert.NotEmpty(t, preparedInput) {
// 		assert.Equal(t, preparedPuzzleInput1, preparedInput)
// 	}
// }

// func Test_Part1(t *testing.T) {
// 	preparedPuzzleInput1p1 := make(CargoBay, len(preparedPuzzleInput1))
// 	copy(preparedPuzzleInput1p1, preparedPuzzleInput1)
// 	resPart1 := part1Func(preparedPuzzleInput1p1, preparedPuzzleInput2)

// 	assert.EqualValues(t, "CMZ", resPart1)
// }

// func Test_Part2(t *testing.T) {
// 	preparedPuzzleInput1p2 := make(CargoBay, len(preparedPuzzleInput1))
// 	copy(preparedPuzzleInput1p2, preparedPuzzleInput1)
// 	resPart2 := part2Func(preparedPuzzleInput1p2, preparedPuzzleInput2)

// 	assert.EqualValues(t, "MCD", resPart2)
// }
