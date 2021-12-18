package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type HeightmapElem struct {
	visited bool
	height  uint8
}

type Heightmap [][]*HeightmapElem

func (m *Heightmap) CheckLocalMinimum(x, y int) bool {
	mid := (*m)[y][x].height
	surrounding := make([]uint8, 0, 4)

	if y > 0 {
		surrounding = append(surrounding, (*m)[y-1][x].height)
	}
	if y < len(*m)-1 {
		surrounding = append(surrounding, (*m)[y+1][x].height)
	}
	if x > 0 {
		surrounding = append(surrounding, (*m)[y][x-1].height)
	}
	if x < len((*m)[0])-1 {
		surrounding = append(surrounding, (*m)[y][x+1].height)
	}
	for _, surr := range surrounding {
		if surr <= mid {
			return false
		}
	}
	return true
}

func (m *Heightmap) GetBasinSize(x, y int) uint16 {
	curr := (*m)[y][x]
	if curr.visited {
		return 0
	}
	curr.visited = true
	if curr.height == 9 {
		return 0
	}
	var size uint16
	if y-1 >= 0 {
		size += m.GetBasinSize(x, y-1)
	}
	if y+1 <= len(*m)-1 {
		size += m.GetBasinSize(x, y+1)
	}
	if x-1 >= 0 {
		size += m.GetBasinSize(x-1, y)
	}
	if x+1 <= len((*m)[0])-1 {
		size += m.GetBasinSize(x+1, y)
	}
	return size + 1
}

func init() {
	rootCmd.AddCommand(day9Cmd)

	day9Cmd.AddCommand(day9part1Cmd)
	day9Cmd.AddCommand(day9part2Cmd)
}

var (
	day9logger = log.With().
			Int("day", 9).
			Logger()
	day9Cmd = &cobra.Command{
		Use:   "day9",
		Short: "Day 9 Challenge",
	}
	day9part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 9 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day9part1logger := day9logger.With().
				Int("part", 1).
				Logger()
			day9part1logger.Info().Msg("Start")
			converted := prepareday9Input()

			var totalRisk uint32
			for y := 0; y < len(converted); y++ {
				for x := 0; x < len(converted[y]); x++ {
					if converted.CheckLocalMinimum(x, y) {
						totalRisk += 1 + uint32(converted[y][x].height)
					}
				}
			}

			fmt.Printf("%v\n", totalRisk)
		},
	}
	day9part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 9 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day9part2logger := day9logger.With().
				Int("part", 2).
				Logger()
			day9part2logger.Info().Msg("Start")
			converted := prepareday9Input()

			biggestBasins := make([]int, 0, 3)
			for y := 0; y < len(converted); y++ {
				for x := 0; x < len(converted[y]); x++ {
					if converted.CheckLocalMinimum(x, y) {
						basinSize := int(converted.GetBasinSize(x, y))
						if len(biggestBasins) == 0 || basinSize > biggestBasins[0] {
							if len(biggestBasins) == 3 {
								_, biggestBasins = biggestBasins[0], biggestBasins[1:]
							}
							biggestBasins = append(biggestBasins, basinSize)
							sort.Ints(biggestBasins)
						}
					}
				}
			}
			day9part2logger.Debug().Msgf("%v", biggestBasins)
			basinMultiple := uint32(1)
			for _, size := range biggestBasins {
				basinMultiple *= uint32(size)
			}
			fmt.Printf("%v\n", basinMultiple)
		},
	}
)

func prepareday9Input() Heightmap {
	content, err := os.ReadFile("resources/day9.txt")
	if err != nil {
		day9logger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	day9logger.Info().Msgf("length of input file: %d", len(input))
	day9logger.Debug().Msgf("plain input: %v", input)

	converted := make(Heightmap, len(input))
	for y := 0; y < len(input); y++ {
		converted[y] = make([]*HeightmapElem, len(input[y]))
		for x := 0; x < len(input[y]); x++ {
			converted[y][x] = &HeightmapElem{height: uint8(input[y][x] - '0')}
		}
	}

	day9logger.Debug().Msgf("converted input: %v", converted)

	return converted
}
