package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

type Polymer map[string]uint64

func (p Polymer) Step(insertionRuleMap map[string]rune) Polymer {
	res := Polymer{}
	for pair, amount := range p {
		day14logger.Trace().Msgf("polymer has %d of curr pair %s", amount, pair)
		if insert, ok := insertionRuleMap[pair]; ok {
			res[string([]rune{rune(pair[0]), insert})] += amount
			res[string([]rune{insert, rune(pair[1])})] += amount
		} else {
			day14logger.Warn().Msgf("No insertion rule found for pair: %s", pair)
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

func init() {
	rootCmd.AddCommand(day14Cmd)
}

var (
	day14logger = log.With().
			Int("day", 14).
			Logger()
	day14Cmd = &cobra.Command{
		Use:   "day14",
		Short: "Day 14 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				return
			}
			steps := cast.ToInt(args[0])

			day14logger.Info().Msg("Start")
			polymerTemplate, insertionRuleMap := prepareday14Input()

			for i := 1; i <= steps; i++ {
				polymerTemplate = polymerTemplate.Step(insertionRuleMap)
				day14logger.Debug().Msgf("after step %d polymer looks like %v", i, polymerTemplate)
			}
			elemCount := polymerTemplate.ElemCount()
			day14logger.Debug().Msgf("polymer elem count %v", elemCount)
			tupleSlice := []utils.Tuple{}
			for k, v := range elemCount {
				tupleSlice = append(tupleSlice, utils.Tuple{k, v})
			}
			sort.Slice(tupleSlice, func(i, j int) bool { return tupleSlice[i][1].(uint64) > tupleSlice[j][1].(uint64) })

			fmt.Printf("%d\n", tupleSlice[0][1].(uint64)-tupleSlice[len(tupleSlice)-1][1].(uint64))
		},
	}
)

func prepareday14Input() (Polymer, map[string]rune) {
	content, err := os.ReadFile("resources/day14.txt")
	if err != nil {
		day14logger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	day14logger.Info().Msgf("length of input file: %d", len(input))
	day14logger.Debug().Msgf("plain input: %v", input)

	polymerTemplate := input[0]
	day14logger.Debug().Msgf("polymer template: %s", polymerTemplate)

	polymer := Polymer{}
	polymerPairs := utils.SlidingWindowString(2, polymerTemplate)
	for _, pair := range polymerPairs {
		polymer[pair]++
	}
	day14logger.Debug().Msgf("converted polymer template: %v", polymer)

	input = input[2:]
	day14logger.Debug().Msgf("pair insertion rules: %v", input)

	converted := map[string]rune{}
	for _, rule := range input {
		ruleParts := strings.Split(rule, " -> ")
		converted[ruleParts[0]] = ([]rune(ruleParts[1]))[0]
	}
	day14logger.Debug().Msgf("converted pair insertion rules: %v", converted)

	return polymer, converted
}
