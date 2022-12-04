package day01

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
	DAY = "01"
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
	depths := prepareInput()

	var increaseCount int
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			increaseCount++
		}
	}

	fmt.Printf("num of increases: %d\n", increaseCount)
}

func Part2(args []string) {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	depths := prepareInput()

	sliding := utils.SlidingWindow[[][]uint16](depths, 3)
	slidingSums := utils.CombineFunc(sliding, func(in []uint16) uint16 {
		var res uint16
		for _, v := range in {
			res += v
		}
		return res
	})

	var increaseCount int
	for i := 1; i < len(slidingSums); i++ {
		if slidingSums[i-1] < slidingSums[i] {
			increaseCount++
		}
	}

	fmt.Printf("num of sliding increases: %d\n", increaseCount)
}

func prepareInput() []uint16 {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2021/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]uint16, len(input))
	for i := 0; i < len(input); i++ {
		converted[i] = cast.ToUint16(input[i])
		partLogger.Trace().Int("index", i).Msgf("converted: %d", converted[i])
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	return converted
}
