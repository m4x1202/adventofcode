package day14

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = 14
)

var (
	dayLogger = log.With().
			Int("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func Part1(args []string) {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	if len(args) != 1 {
		partLogger.Error().Msg("Incorrect number of arguments")
		return
	}
	steps := cast.ToInt(args[0])

	partLogger.Info().Msg("Start")
	polymerTemplate, insertionRuleMap := prepareInput()

	for i := 1; i <= steps; i++ {
		polymerTemplate = polymerTemplate.Step(insertionRuleMap)
		partLogger.Debug().Msgf("after step %d polymer looks like %v", i, polymerTemplate)
	}
	elemCount := polymerTemplate.ElemCount()
	partLogger.Debug().Msgf("polymer elem count %v", elemCount)
	tupleSlice := []utils.Tuple{}
	for k, v := range elemCount {
		tupleSlice = append(tupleSlice, utils.Tuple{k, v})
	}
	sort.Slice(tupleSlice, func(i, j int) bool { return tupleSlice[i][1].(uint64) > tupleSlice[j][1].(uint64) })

	fmt.Printf("%d\n", tupleSlice[0][1].(uint64)-tupleSlice[len(tupleSlice)-1][1].(uint64))
}

func prepareInput() (Polymer, map[string]rune) {
	content, err := os.ReadFile(fmt.Sprintf("internal/app/day%d/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	polymerTemplate := input[0]
	partLogger.Debug().Msgf("polymer template: %s", polymerTemplate)

	polymer := Polymer{}
	polymerPairs := utils.SlidingWindowString(2, polymerTemplate)
	for _, pair := range polymerPairs {
		polymer[pair]++
	}
	partLogger.Debug().Msgf("converted polymer template: %v", polymer)

	input = input[2:]
	partLogger.Debug().Msgf("pair insertion rules: %v", input)

	converted := map[string]rune{}
	for _, rule := range input {
		ruleParts := strings.Split(rule, " -> ")
		converted[ruleParts[0]] = ([]rune(ruleParts[1]))[0]
	}
	partLogger.Debug().Msgf("converted pair insertion rules: %v", converted)

	return polymer, converted
}

type Polymer map[string]uint64

func (p Polymer) Step(insertionRuleMap map[string]rune) Polymer {
	res := Polymer{}
	for pair, amount := range p {
		partLogger.Trace().Msgf("polymer has %d of curr pair %s", amount, pair)
		if insert, ok := insertionRuleMap[pair]; ok {
			res[string([]rune{rune(pair[0]), insert})] += amount
			res[string([]rune{insert, rune(pair[1])})] += amount
		} else {
			partLogger.Warn().Msgf("No insertion rule found for pair: %s", pair)
			res[pair] = amount
		}
	}
	return res
}

func (p Polymer) ElemCount() map[rune]uint64 {
	res := map[rune]uint64{}
	for pair, amount := range p {
		res[rune(pair[1])] += amount
	}
	return res
}