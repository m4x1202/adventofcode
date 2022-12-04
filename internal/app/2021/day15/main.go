package day15

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	astar "github.com/beefsack/go-astar"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DAY = 15
)

var (
	dayLogger = log.With().
			Int("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func Part1() {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	converted := prepareInput()

	p, _, found := astar.Path(converted[0][0], converted[len(converted)-1][len(converted[0])-1])
	if !found {
		partLogger.Error().Msg("Could not find path")
	}
	var totalRisk uint
	// Skip first element since it is our starting node
	for _, step := range p[1:] {
		hElem := step.(*HeightmapElem)
		totalRisk += uint(hElem.Height)
	}

	fmt.Printf("%d\n", totalRisk)
}

func Part2() {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	converted := prepareInput()
	converted = postPrepareInput(converted)

	p, _, found := astar.Path(converted[0][0], converted[len(converted)-1][len(converted[0])-1])
	if !found {
		partLogger.Error().Msg("Could not find path")
	}
	var totalRisk uint
	// Skip first element since it is our starting node
	for _, step := range p[1:] {
		hElem := step.(*HeightmapElem)
		totalRisk += uint(hElem.Height)
	}

	fmt.Printf("%d\n", totalRisk)
}

func prepareInput() Heightmap {
	content, err := os.ReadFile(fmt.Sprintf("internal/app/day%d/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	converted := Heightmap{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			newElem := &HeightmapElem{Height: uint8(input[y][x] - '0')}
			converted.SetTile(newElem, x, y)
		}
	}

	partLogger.Debug().Msgf("converted input: %v", converted)

	return converted
}

func postPrepareInput(m Heightmap) Heightmap {
	res := Heightmap{}
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			elem := m.Tile(x, y)
			for xMul := 0; xMul < 5; xMul++ {
				for yMul := 0; yMul < 5; yMul++ {
					newElem := &HeightmapElem{Height: (elem.Height-1+uint8(xMul)+uint8(yMul))%9 + 1}
					res.SetTile(newElem, x+len(m)*xMul, y+len(m[x])*yMul)
				}
			}
		}
	}
	partLogger.Debug().Msgf("post prepared input: %v", res)
	return res
}

type HeightmapElem struct {
	Height uint8
	X, Y   int
	Map    Heightmap
}

func (t *HeightmapElem) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.Map.Tile(t.X+offset[0], t.Y+offset[1]); n != nil {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (t *HeightmapElem) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*HeightmapElem)
	return float64(toT.Height)
}

func (t *HeightmapElem) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*HeightmapElem)
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

type Heightmap map[int]map[int]*HeightmapElem

func (w Heightmap) Tile(x, y int) *HeightmapElem {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

func (w Heightmap) SetTile(t *HeightmapElem, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*HeightmapElem{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.Map = w
}

func (w Heightmap) String() string {
	var builder strings.Builder
	for y := 0; y < len(w[0]); y++ {
		for x := 0; x < len(w); x++ {
			builder.WriteString(strconv.Itoa(int(w[x][y].Height)))
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
