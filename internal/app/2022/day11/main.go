// day11 is the template package which can be copied, modified to its final location
package day11

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
	DAY = "11"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func ExecutePart(p uint8) {
	preparedInput := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		part1Func(preparedInput)
	case 2:
		part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
}

func part1Func(preparedInput []Monkey) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	// Logic here
	puzzleAnswer = cast.ToUint64(0)
	return puzzleAnswer
}

func part2Func(preparedInput []Monkey) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	// Logic here
	puzzleAnswer = cast.ToUint64(0)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []Monkey {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	splitInput := utils.SliceMap(utils.SplitSlice[[][]string](input, ""), func(in []string) []string {
		return utils.SliceMap(in, strings.TrimSpace)
	})
	converted := make([]Monkey, len(splitInput))
	for i := range splitInput {
		converted[i] = ParseMonkey(splitInput[i])
	}

	return converted
}

type Monkey struct {
	Name      uint8
	Items     []uint
	Operation string
	Test      uint8
	Decision  [2]uint8
}

func ParseMonkey(in []string) Monkey {
	res := Monkey{
		Items:    []uint{},
		Decision: [2]uint8{},
	}
	nameString := in[0]
	startingItemsString := in[1]
	operationString := in[2]
	testString := in[3]
	decision1String := in[4]
	decision2String := in[5]
	res.Name = uint8(nameString[len(nameString)-2] - '0')
	res.Test = cast.ToUint8(strings.TrimPrefix(testString, "Test: divisible by "))
	res.Decision[0] = cast.ToUint8(strings.TrimPrefix(decision1String, "If true: throw to monkey "))
	res.Decision[1] = cast.ToUint8(strings.TrimPrefix(decision2String, "If false: throw to monkey "))
	res.Operation = strings.TrimPrefix(operationString, "Operation: ")
	for _, item := range strings.Split(strings.TrimPrefix(startingItemsString, "Starting items: "), ", ") {
		res.Items = append(res.Items, cast.ToUint(item))
	}

	return res
}
