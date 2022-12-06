package cmd2021

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

type OctomapElem struct {
	x       int
	y       int
	flashed bool
	energy  uint8
	m       *Octomap
}

func (e *OctomapElem) EnergyInc() uint16 {
	if e.flashed {
		return 0
	}
	e.energy++
	if e.energy > 9 {
		e.flashed = true
		e.energy = 0
		return 1 + e.m.FlashAt(int(e.x), int(e.y))
	}
	return 0
}

func (e *OctomapElem) Reset() {
	e.flashed = false
}

type Octomap [][]*OctomapElem

func (m *Octomap) FlashAt(a, b int) uint16 {
	var totalFlashes uint16
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			totalX := a - 1 + x
			totalY := b - 1 + y
			if totalY < 0 || totalY > len(*m)-1 {
				continue
			}
			if totalX < 0 || totalX > len((*m)[totalY])-1 {
				continue
			}
			totalFlashes += (*m)[totalY][totalX].EnergyInc()
		}
	}
	return totalFlashes
}

func (m *Octomap) Step() (uint16, bool) {
	var totalFlashes uint16
	for y := 0; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			totalFlashes += (*m)[y][x].EnergyInc()
		}
	}
	for y := 0; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			(*m)[y][x].Reset()
		}
	}
	if totalFlashes == uint16(len(*m)*len((*m)[0])) {
		return totalFlashes, true
	}
	return totalFlashes, false
}

func init() {
	cmd2021.AddCommand(day11Cmd)

	day11Cmd.AddCommand(day11part1Cmd)
	day11Cmd.AddCommand(day11part2Cmd)
}

var (
	day11logger = log.With().
			Int("day", 11).
			Logger()
	day11Cmd = &cobra.Command{
		Use:   "day11",
		Short: "Day 11 Challenge",
	}
	day11part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 11 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day11part1logger := day11logger.With().
				Int("part", 1).
				Logger()

			if len(args) != 1 {
				return
			}
			steps := cast.ToInt(args[0])

			day11part1logger.Info().Msg("Start")
			converted := prepareday11Input()

			var totalFlashes uint32
			for step := 0; step < steps; step++ {
				flashes, _ := converted.Step()
				totalFlashes += uint32(flashes)
			}

			fmt.Printf("%d\n", totalFlashes)
		},
	}
	day11part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 11 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day11part2logger := day11logger.With().
				Int("part", 2).
				Logger()
			day11part2logger.Info().Msg("Start")
			converted := prepareday11Input()

			steps := uint32(1)
			for {
				_, sync := converted.Step()
				if sync {
					break
				}
				steps++
			}

			fmt.Printf("%d\n", steps)
		},
	}
)

func prepareday11Input() *Octomap {
	content, err := os.ReadFile("resources/day11.txt")
	if err != nil {
		day11logger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
	day11logger.Info().Msgf("length of input file: %d", len(input))
	day11logger.Debug().Msgf("plain input: %v", input)

	converted := make(Octomap, len(input))
	for y := 0; y < len(input); y++ {
		converted[y] = make([]*OctomapElem, len(input[y]))
		for x := 0; x < len(input[y]); x++ {
			converted[y][x] = &OctomapElem{
				x:      x,
				y:      y,
				energy: uint8(input[y][x] - '0'),
				m:      &converted,
			}
		}
	}

	return &converted
}
