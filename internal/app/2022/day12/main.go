package day12

import (
	"fmt"
	"strings"

	"github.com/beefsack/go-astar"
	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "12"
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

func part1Func(heightmap utils.Map[HeightmapElem]) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	xS, yS := FindS(heightmap)
	xE, yE := FindE(heightmap)

	//p, _, found := astar.Path(heightmap[0][0], heightmap[5][2])
	p, _, found := astar.Path(heightmap[xS][yS], heightmap[xE][yE])
	if !found {
		partLogger.Error().Msg("Could not find path")
	}
	partLogger.Debug().Msgf("%v", p)

	fmt.Printf("total path length: %d\n", len(p)-1)
	puzzleAnswer = cast.ToUint64(len(p) - 1)
	return puzzleAnswer
}

func part2Func(heightmap utils.Map[HeightmapElem]) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	xE, yE := FindE(heightmap)

	fastestPath := int(1000)
	for x := range heightmap {
		for y := range heightmap[x] {
			if heightmap[x][y].Height == 'a' || heightmap[x][y].Height == 'S' {
				p, dist, found := astar.Path(heightmap[x][y], heightmap[xE][yE])
				if !found {
					partLogger.Error().Msg("Could not find path")
				}
				if dist > 1000 {
					continue
				}
				if len(p)-1 < fastestPath {
					partLogger.Debug().Msgf("%v", p)
					fastestPath = len(p) - 1
				}
			}
		}
	}

	fmt.Printf("fastest path length: %d\n", fastestPath)
	puzzleAnswer = cast.ToUint64(fastestPath)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) utils.Map[HeightmapElem] {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := utils.NewMap[HeightmapElem](len(input[0]), len(input))
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			converted[x][y] = HeightmapElem{
				Height: rune(input[y][x]),
				X:      x,
				Y:      y,
				Map:    &converted,
			}
		}
	}
	dayLogger.Trace().Msgf("converted input: %v", converted)

	return converted
}

func FindS(heightmap utils.Map[HeightmapElem]) (x, y int) {
	for x := range heightmap {
		for y := range heightmap[x] {
			if heightmap[x][y].Height == 'S' {
				return x, y
			}
		}
	}
	panic("")
}

func FindE(heightmap utils.Map[HeightmapElem]) (x, y int) {
	for x := range heightmap {
		for y := range heightmap[x] {
			if heightmap[x][y].Height == 'E' {
				return x, y
			}
		}
	}
	panic("")
}

// Ensure HeightmapElem implements astar.Pather
var _ astar.Pather = (*HeightmapElem)(nil)

type HeightmapElem struct {
	Height rune
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

func (from HeightmapElem) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(HeightmapElem)
	toheight := toT.Height
	fromheight := from.Height
	if toT.Height == 'E' {
		toheight = 'z'
	}
	if toT.Height == 'S' {
		toheight = 'a'
	}
	if from.Height == 'S' {
		fromheight = 'a'
	}
	if from.Height == 'E' {
		fromheight = 'z'
	}
	if toheight-fromheight <= 1 {
		return 1
	} else {
		return 1000
	}
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
	return string(t.Height)
}
