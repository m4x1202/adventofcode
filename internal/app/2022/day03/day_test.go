package day03

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var testPuzzleInput string

var (
	preparedPuzzleInput = []Rucksack{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrVvPwwTWBwg", "PmmdzqPrV", "vPwwTWBwg"},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
		{"ttgJtRGJQctTZtZT", "ttgJtRGJ", "QctTZtZT"},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", "CrZsJsPPZsGz", "wwsLwLmpwMDw"},
	}
)

func Test_PrepareInput(t *testing.T) {
	preparedInput := prepareInput(testPuzzleInput)

	if assert.NotEmpty(t, preparedInput) {
		assert.Equal(t, preparedPuzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(preparedPuzzleInput)

	assert.EqualValues(t, 157, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(preparedPuzzleInput)

	assert.EqualValues(t, 70, resPart2)
}
