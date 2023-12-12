package day02

import (
	"fmt"
	"strconv"
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

func ExecutePart(p uint8) uint64 {
	preparedInput := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		return part1Func(preparedInput)
	case 2:
		return part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
}

func part1Func(preparedInput []Game) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for _, game := range preparedInput {
		if game.Bag.ComparisonBagPossible(ComparisonBag) {
			puzzleAnswer += cast.ToUint64(game.ID)
		}
	}

	return puzzleAnswer
}

func part2Func(preparedInput []Game) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for _, game := range preparedInput {
		puzzleAnswer += uint64(game.Bag[Blue] * game.Bag[Green] * game.Bag[Red])
	}

	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2023/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []Game {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	games := make([]Game, len(input))
	for i := range input {
		gameIDStr := strings.Split(input[i], ": ")
		dayLogger.Debug().Msgf("%v", gameIDStr)
		newGame := Game{
			ID:  strings.TrimPrefix(gameIDStr[0], "Game "),
			Bag: NewBag(gameIDStr[1]),
		}
		games[i] = newGame
	}

	return games
}

var (
	ComparisonBag = Bag{
		Blue:  14,
		Green: 13,
		Red:   12,
	}
)

type Game struct {
	ID  string
	Bag Bag
}

type Color uint8

const (
	Blue Color = iota + 1
	Green
	Red
)

func ParseColor(in string) Color {
	switch in {
	case "blue":
		return Blue
	case "green":
		return Green
	case "red":
		return Red
	default:
		return 0
	}
}

type Bag map[Color]uint

func (b Bag) ComparisonBagPossible(comparison Bag) bool {
	if b[Blue] <= comparison[Blue] && b[Green] <= comparison[Green] && b[Red] <= comparison[Red] {
		return true
	}
	return false
}

func NewBag(in string) Bag {
	result := Bag{}
	result.ParseRandSelection(in)
	return result
}

func (b Bag) ParseRandSelection(in string) {
	selection := strings.Split(in, "; ")
	dayLogger.Debug().Msgf("%v", selection)
	for _, entry := range selection {
		b.ParseInformation(entry)
	}
}

func (b Bag) ParseInformation(in string) {
	selection := strings.Split(in, ", ")
	for _, entry := range selection {
		se := strings.Split(entry, " ")
		color := ParseColor(se[1])
		num, err := strconv.ParseUint(se[0], 10, 64)
		if err != nil {
			partLogger.Fatal().Err(err).Send()
		}
		if b[color] < uint(num) {
			b[color] = uint(num)
		}
	}
}
