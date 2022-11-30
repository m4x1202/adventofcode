package day01

import (
	"fmt"
	"os"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DAY = 01
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
	prepareInput()

	// Logic here
}

func Part2(args []string) {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	prepareInput()

	// Logic here
}

func prepareInput() any {
	// fsys, err := fs.Sub(inputs.InputFS_2021, "2021")
	// if err != nil {
	// 	day8logger.Fatal().Err(err).Send()
	// }
	// file, err := fsys.Open("day08/input.txt")
	// if err != nil {
	// 	day8logger.Fatal().Err(err).Send()
	// }
	// content, err := io.ReadAll(file)
	// if err != nil {
	// 	day8logger.Fatal().Err(err).Send()
	// }
	content, err := resources.InputFS.ReadFile("2021/day08/input.txt")
	if err != nil {
		day8logger.Fatal().Err(err).Send()
	}
	content, err := os.ReadFile(fmt.Sprintf("internal/app/day%d/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	// Required input conversion here

	return input
}
