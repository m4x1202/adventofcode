package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DAY = "10"
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

func part1Func(preparedInput []Command) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	signalStrengths := map[uint]int64{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}

	registerV := int64(1)
	cycle := uint(1)

	for _, command := range preparedInput {
		for i := uint8(0); i < uint8(command.Instruction); i++ {
			if _, exists := signalStrengths[cycle]; exists {
				signalStrengths[cycle] = registerV * int64(cycle)
			}
			cycle++
		}
		registerV += command.Data
	}

	for _, signal := range signalStrengths {
		puzzleAnswer += uint64(signal)
	}

	fmt.Printf("sum of signal strengths: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(preparedInput []Command) string {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")

	crt := utils.CoordinateSystem[int, bool]{}
	registerV := int64(1)
	cycle := uint(0)
	crtY := int(0)

	for _, command := range preparedInput {
		for i := uint8(0); i < uint8(command.Instruction); i++ {
			if cycle != 0 && cycle%40 == 0 {
				crtY--
			}
			partLogger.Trace().Msgf("crtY: %d, cycle: %d, registerV: %d", crtY, cycle%40, registerV)
			if registerV-1 <= int64(cycle%40) && int64(cycle%40) <= registerV+1 {
				crt.ModifyElemFunc(func(elem *bool) *bool {
					var res bool
					return &res
				}, int(cycle%40), crtY)
			}
			partLogger.Trace().Msgf("total length crt: %d", crt.TotalSize())
			cycle++
		}
		registerV += command.Data
	}

	res := crt.String()
	fmt.Print(res)
	return res
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []Command {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]Command, len(input))
	for i := range input {
		converted[i] = ParseCommand(input[i])
	}

	return converted
}

type Instruction uint8

const (
	_ Instruction = iota
	Noop
	Addx
)

type Command struct {
	Instruction Instruction
	Data        int64
}

func ParseCommand(in string) Command {
	splitCommand := strings.Split(in, " ")
	switch splitCommand[0] {
	case "noop":
		return Command{Noop, 0}
	case "addx":
		data, err := strconv.ParseInt(splitCommand[1], 10, 64)
		if err != nil {
			partLogger.Error().Msgf("command %s could not be parsed", in)
		}
		return Command{Addx, data}
	}
	return Command{}
}
