package day15

import (
	"fmt"
	"strings"

	astar "github.com/beefsack/go-astar"
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

func part1Func(riskMap utils.Map[HeightmapElem]) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	p, _, found := astar.Path(riskMap[0][0], riskMap[len(riskMap)-1][len(riskMap[0])-1])
	if !found {
		partLogger.Error().Msg("Could not find path")
	}
	var totalRisk uint
	// Skip first element since it is our starting node
	for _, step := range p[:len(p)-1] {
		hElem := step.(HeightmapElem)
		totalRisk += uint(hElem.Height)
	}

	fmt.Printf("total risk of path is %d\n", totalRisk)
	puzzleAnswer = cast.ToUint64(totalRisk)
	return puzzleAnswer
}

func part2Func(riskMap utils.Map[HeightmapElem]) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	newRiskMap := part2InputModifications(riskMap)

	p, _, found := astar.Path(newRiskMap[0][0], newRiskMap[len(newRiskMap)-1][len(newRiskMap[0])-1])
	if !found {
		partLogger.Error().Msg("Could not find path")
	}
	var totalRisk uint
	// Skip first element since it is our starting node
	for _, step := range p[:len(p)-1] {
		hElem := step.(HeightmapElem)
		totalRisk += uint(hElem.Height)
	}

	fmt.Printf("total risk of path is %d\n", totalRisk)
	puzzleAnswer = cast.ToUint64(totalRisk)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2021/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) utils.Map[HeightmapElem] {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Trace().Msgf("plain input: %v", input)

	converted := utils.NewMap[HeightmapElem](len(input[0]), len(input))
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			converted[x][y] = HeightmapElem{
				Height: uint8(input[y][x] - '0'),
				X:      x,
				Y:      y,
				Map:    &converted,
			}
		}
	}
	dayLogger.Trace().Msgf("converted input: %v", converted)

	return converted
}

func part2InputModifications(m utils.Map[HeightmapElem]) utils.Map[HeightmapElem] {
	width := len(m)
	height := len(m[0])
	res := utils.NewMap[HeightmapElem](width*5, height*5)
	partLogger.Debug().Msgf("new width %d, new height %d", len(res), len(res[0]))
	for x, c := range res {
		for y := range c {
			n := m[x%width][y%height].Height
			res[x][y] = HeightmapElem{
				Height: (n+uint8(x/width+y/height)-1)%9 + 1,
				X:      x,
				Y:      y,
				Map:    &res,
			}
		}
	}
	partLogger.Trace().Msgf("post prepared input: %v", res)
	return res
}

// Ensure HeightmapElem implements astar.Pather
var _ astar.Pather = (*HeightmapElem)(nil)

type HeightmapElem struct {
	Height uint8
	X, Y   int
	Map    *utils.Map[HeightmapElem]
}

func (t HeightmapElem) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n, ok := t.Map.Tile(t.X+offset[0], t.Y+offset[1]); ok {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (_ HeightmapElem) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(HeightmapElem)
	return float64(toT.Height)
}

func (t HeightmapElem) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(HeightmapElem)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (t HeightmapElem) String() string {
	return cast.ToString(t.Height)
}
