package day02

import (
	"fmt"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Shape int

const (
	_ Shape = iota
	Rock
	Paper
	Scissors
)

func ParseShape(in string) Shape {
	switch in {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		return -1
	}
}

type RoundOutcome int

const (
	Loss    RoundOutcome = 0
	Draw    RoundOutcome = 3
	Victory RoundOutcome = 6
)

func ParseOutcome(in string) RoundOutcome {
	switch in {
	case "X":
		return Loss
	case "Y":
		return Draw
	case "Z":
		return Victory
	default:
		return -1
	}
}

type Round struct {
	OpponentShape Shape
	MyShape       Shape
	Outcome       RoundOutcome
}

func ParseRound1(in string) Round {
	splitRoundInput := strings.Split(in, " ")
	myShape := ParseShape(splitRoundInput[1])
	opponentShape := ParseShape(splitRoundInput[0])

	return Round{
		OpponentShape: opponentShape,
		MyShape:       myShape,
		Outcome:       CalculateRoundOutcome(opponentShape, myShape),
	}
}

func ParseRound2(in string) Round {
	splitRoundInput := strings.Split(in, " ")
	opponentShape := ParseShape(splitRoundInput[0])
	roundOutcome := ParseOutcome(splitRoundInput[1])

	return Round{
		OpponentShape: opponentShape,
		MyShape:       CalculateMyShape(opponentShape, roundOutcome),
		Outcome:       roundOutcome,
	}
}

var (
	CalculateRoundOutcome = func(opponent, me Shape) RoundOutcome {
		switch {
		case opponent == Rock && me == Paper:
			return Victory
		case opponent == Paper && me == Scissors:
			return Victory
		case opponent == Scissors && me == Rock:
			return Victory
		case opponent == me:
			return Draw
		default:
			return Loss
		}
	}
	CalculateMyShape = func(opponent Shape, requiredOutcome RoundOutcome) Shape {
		switch {
		case requiredOutcome == Draw:
			return opponent
		case requiredOutcome == Loss:
			switch opponent {
			case Rock:
				return Scissors
			case Paper:
				return Rock
			case Scissors:
				return Paper
			}
		case requiredOutcome == Victory:
			switch opponent {
			case Rock:
				return Paper
			case Paper:
				return Scissors
			case Scissors:
				return Rock
			}
		}
		return 0
	}
)

const (
	DAY = "02"
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
	rounds := prepareInput()

	converted := make([]Round, len(rounds))
	for i := 0; i < len(rounds); i++ {
		converted[i] = ParseRound1(rounds[i])
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	var score int
	for _, round := range converted {
		score += int(round.MyShape)
		score += int(round.Outcome)
	}

	fmt.Printf("my total score: %d\n", score)
}

func Part2(args []string) {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	rounds := prepareInput()

	converted := make([]Round, len(rounds))
	for i := 0; i < len(rounds); i++ {
		converted[i] = ParseRound2(rounds[i])
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	var score int
	for _, round := range converted {
		score += int(round.MyShape)
		score += int(round.Outcome)
	}

	fmt.Printf("my total score: %d\n", score)
}

func prepareInput() []string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	return input
}
