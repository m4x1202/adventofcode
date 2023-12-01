package day15

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
	DAY = "15"
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

func part1Func(preparedInput any) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	// Logic here
	puzzleAnswer = cast.ToUint64(0)
	return puzzleAnswer
}

func part2Func(preparedInput any) uint64 {
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

func prepareInput(rawInput string) utils.CoordinateSystem[int, rune] {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := utils.CoordinateSystem[int, rune]{}
	for _, elem := range input {
		splitInput := strings.Split(elem, ": ")
		sensorString := strings.TrimPrefix(splitInput[0], "Sensor at ")
		sensorCoords := strings.Split(sensorString, ", ")
		sensorX := cast.ToInt(strings.TrimPrefix(sensorCoords[0], "x="))
		sensorY := cast.ToInt(strings.TrimPrefix(sensorCoords[1], "y="))
		converted.ModifyElemFunc(func(elem *rune) *rune {
			res := 'S'
			return &res
		}, sensorX, sensorY)
		beaconString := strings.TrimPrefix(splitInput[0], "closest beacon is at ")
		beaconCoords := strings.Split(beaconString, ", ")
		beaconX := cast.ToInt(strings.TrimPrefix(beaconCoords[0], "x="))
		beaconY := cast.ToInt(strings.TrimPrefix(beaconCoords[1], "y="))
		converted.ModifyElemFunc(func(elem *rune) *rune {
			res := 'B'
			return &res
		}, beaconX, beaconY)
	}

	return converted
}
