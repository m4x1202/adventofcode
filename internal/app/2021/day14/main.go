package day14

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func Part1() {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	pairInsertion(10)
}

func Part2() {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	pairInsertion(40)
}

func pairInsertion(steps uint8) {
	polymerTemplate, insertionRuleMap := prepareInput()

	for i := 1; i <= int(steps); i++ {
		polymerTemplate = polymerTemplate.Step(insertionRuleMap)
		partLogger.Debug().Msgf("after step %d polymer looks like %v", i, polymerTemplate)
	}
	elemCount := polymerTemplate.ElemCount()
	partLogger.Debug().Msgf("polymer elem count %v", elemCount)
	tupleSlice := []utils.Tuple[rune, uint64]{}
	for k, v := range elemCount {
		tupleSlice = append(tupleSlice, utils.Tuple[rune, uint64]{V1: k, V2: v})
	}
	sort.Slice(tupleSlice, func(i, j int) bool { return tupleSlice[i].V2 > tupleSlice[j].V2 })

	fmt.Printf("%d\n", tupleSlice[0].V2-tupleSlice[len(tupleSlice)-1].V2)
}

func prepareInput() (Polymer, map[string]rune) {
	content, err := os.ReadFile(fmt.Sprintf("internal/app/day%d/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	polymerTemplate := input[0]
	partLogger.Debug().Msgf("polymer template: %s", polymerTemplate)

	polymer := Polymer{}
	polymerPairs := utils.SlidingWindow[[][]rune]([]rune(polymerTemplate), 2)
	for _, pair := range polymerPairs {
		polymer[string(pair)]++
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
