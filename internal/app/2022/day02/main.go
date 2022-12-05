package day02

import (
	"errors"
	"fmt"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
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

func part1Func(rounds []string) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	converted := make([]Round, len(rounds))
	for i := 0; i < len(rounds); i++ {
		splitRoundInput := strings.Split(rounds[i], " ")
		round := Round{
			OpponentShape: ParseShape(splitRoundInput[0]),
			MyShape:       ParseShape(splitRoundInput[1]),
			Outcome:       UNDEF_OUTCOME,
		}
		round.MatchWithPossible()
		converted[i] = round
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	var score uint16
	for _, round := range converted {
		score += uint16(round.MyShape)
		score += uint16(round.Outcome)
	}

	fmt.Printf("my total score: %d\n", score)
	puzzleAnswer = cast.ToUint64(score)
	return puzzleAnswer
}

func part2Func(rounds []string) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	converted := make([]Round, len(rounds))
	for i := 0; i < len(rounds); i++ {
		splitRoundInput := strings.Split(rounds[i], " ")
		round := Round{
			OpponentShape: ParseShape(splitRoundInput[0]),
			MyShape:       UNDEF_SHAPE,
			Outcome:       ParseOutcome(splitRoundInput[1]),
		}
		round.MatchWithPossible()
		converted[i] = round
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	var score uint16
	for _, round := range converted {
		score += uint16(round.MyShape)
		score += uint16(round.Outcome)
	}

	fmt.Printf("my total score: %d\n", score)
	puzzleAnswer = cast.ToUint64(score)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return strings.TrimSpace(string(content))
}

func prepareInput(rawInput string) []string {
	input := strings.Split(rawInput, "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	return input
}

// Helper data structures and functions

type Shape uint8

const (
	Rock Shape = iota + 1
	Paper
	Scissors
	UNDEF_SHAPE
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
		panic("Could not parse shape")
	}
}

type RoundOutcome uint8

const (
	Loss    RoundOutcome = iota
	Draw    RoundOutcome = iota + 2
	Victory RoundOutcome = iota + 4
	UNDEF_OUTCOME
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
		panic("Could not parse outcome")
	}
}

type Round struct {
	OpponentShape Shape
	MyShape       Shape
	Outcome       RoundOutcome
}

func (r *Round) MatchWithPossible() error {
	var allPossibleRounds = map[RoundOutcome][][2]Shape{
		Victory: {{Rock, Paper}, {Paper, Scissors}, {Scissors, Rock}},
		Draw:    {{Rock, Rock}, {Paper, Paper}, {Scissors, Scissors}},
		Loss:    {{Rock, Scissors}, {Paper, Rock}, {Scissors, Paper}},
	}

	switch {
	case r.OpponentShape == UNDEF_SHAPE:
		return errors.New("opponent shape required to validate")
	case r.MyShape == UNDEF_SHAPE:
		if r.Outcome == UNDEF_OUTCOME {
			return errors.New("both my shape and outcome not valid")
		} else {
			r.MyShape = func(opponent Shape, requiredOutcome RoundOutcome) Shape {
				for _, round := range allPossibleRounds[requiredOutcome] {
					if round[0] == opponent {
						return round[1]
					}
				}
				panic("opponent shape + round outcome do not produce a valid my shape")
			}(r.OpponentShape, r.Outcome)
		}
	case r.Outcome == UNDEF_OUTCOME:
		if r.MyShape == UNDEF_SHAPE {
			return errors.New("both my shape and outcome not valid")
		} else {
			r.Outcome = func(opponent, player Shape) RoundOutcome {
				for outcome, possibleRounds := range allPossibleRounds {
					for _, round := range possibleRounds {
						if round == [2]Shape{opponent, player} {
							return outcome
						}
					}
				}
				panic("opponent shape + my shape do not have a valid outcome")
			}(r.OpponentShape, r.MyShape)
		}
	}
	return nil
}
