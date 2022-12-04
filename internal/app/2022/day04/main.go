package day04

import (
	"fmt"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "04"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func Part1(args []string) {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	pairs := prepareInput()

	var doubleAssignmentCounter uint
	for _, pair := range pairs {
		if pair.DoubleAssignment() {
			doubleAssignmentCounter++
		}
	}

	fmt.Printf("pairs with double assignments: %d\n", doubleAssignmentCounter)
}

func Part2(args []string) {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	pairs := prepareInput()

	var overlapAssignmentCounter uint
	for _, pair := range pairs {
		if pair.OverlapAssignment() {
			overlapAssignmentCounter++
		}
	}

	fmt.Printf("pairs with overlap assignments: %d\n", overlapAssignmentCounter)

	var overlapAssignmentCounter2 uint
	for _, pair := range pairs {
		if pair.OverlapAssignment2() {
			overlapAssignmentCounter2++
		}
	}

	fmt.Printf("second solution: pairs with overlap assignments: %d\n", overlapAssignmentCounter2)
}

func prepareInput() []Pair {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]Pair, len(input))
	for i := 0; i < len(input); i++ {
		converted[i] = ParsePair(input[i])
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	return converted
}

type Section [2]uint8

func ParseSection(in string) Section {
	sectionStrings := strings.Split(in, "-")
	return Section{cast.ToUint8(sectionStrings[0]), cast.ToUint8(sectionStrings[1])}
}

func (s Section) ToSlice() []uint8 {
	res := make([]uint8, s[1]-s[0]+1)
	for i := 0; i < len(res); i++ {
		res[i] = uint8(i) + s[0]
	}
	return res
}

func (s1 Section) FullyContains(s2 Section) bool {
	return s1[0] <= s2[0] && s1[1] >= s2[1]
}

func (s1 Section) Overlaps(s2 Section) bool {
	return (s1[0] <= s2[0] && s2[0] <= s1[1]) || (s2[0] <= s1[0] && s1[0] <= s2[1])
}

type Pair [2]Section

func ParsePair(in string) Pair {
	pairStrings := strings.Split(in, ",")
	return Pair{ParseSection(pairStrings[0]), ParseSection(pairStrings[1])}
}

func (p Pair) DoubleAssignment() bool {
	return p[0].FullyContains(p[1]) || p[1].FullyContains(p[0])
}

func (p Pair) OverlapAssignment() bool {
	intersection := utils.Intersection(p[0].ToSlice(), p[1].ToSlice())
	return len(intersection) > 0
}

func (p Pair) OverlapAssignment2() bool {
	return p[0].Overlaps(p[1])
}
