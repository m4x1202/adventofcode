package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPuzzleInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
)

var (
	prepareduzzleInput = []Rucksack{
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
		assert.Equal(t, prepareduzzleInput, preparedInput)
	}
}

func Test_Part1(t *testing.T) {
	resPart1 := part1Func(prepareduzzleInput)

	assert.EqualValues(t, 157, resPart1)
}

func Test_Part2(t *testing.T) {
	resPart2 := part2Func(prepareduzzleInput)

	assert.EqualValues(t, 70, resPart2)
}
