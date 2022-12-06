package day03

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
	DAY = "03"
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

func part1Func(rucksacks []Rucksack) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	var prioSum uint16
	for _, r := range rucksacks {
		prioSum += uint16(CheckPrio(r.FindOverlap()))
	}

	fmt.Printf("priority sum: %d\n", prioSum)
	puzzleAnswer = cast.ToUint64(prioSum)
	return puzzleAnswer
}

func part2Func(rucksacks []Rucksack) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	groups := utils.ChunkSlice[[]ElfGroup, ElfGroup](rucksacks, 3)
	partLogger.Debug().Msgf("elf groups: %v", groups)

	var prioSum uint16
	for _, g := range groups {
		prioSum += uint16(CheckPrio(g.FindBadgeType()))
	}

	fmt.Printf("priority sum: %d\n", prioSum)
	puzzleAnswer = cast.ToUint64(prioSum)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []Rucksack {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]Rucksack, len(input))
	for i := 0; i < len(input); i++ {
		converted[i] = ParseRucksack(input[i])
	}
	dayLogger.Debug().Msgf("converted input: %v", converted)

	return converted
}

const PriorityList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CheckPrio(r rune) uint8 {
	for i, c := range PriorityList {
		if r == c {
			return uint8(i) + 1
		}
	}
	panic("Rune not in Prio List")
}

type ElfGroup []Rucksack

func (g ElfGroup) FindBadgeType() rune {
	g12 := utils.Intersection([]rune(g[0].AllItems), []rune(g[1].AllItems))
	g123 := utils.Intersection(g12, []rune(g[2].AllItems))
	if len(g123) != 1 {
		panic("Could not find badge type")
	}
	return g123[0]
}

type Rucksack struct {
	AllItems     string
	Compartment1 string
	Compartment2 string
}

func ParseRucksack(in string) Rucksack {
	return Rucksack{
		AllItems:     in,
		Compartment1: in[:len(in)/2],
		Compartment2: in[len(in)/2:],
	}
}

func (r Rucksack) FindOverlap() rune {
	g12 := utils.Intersection([]rune(r.Compartment1), []rune(r.Compartment2))
	if len(g12) != 1 {
		panic("Could not find Overlap")
	}
	return g12[0]
}
