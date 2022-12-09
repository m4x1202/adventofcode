package cmd2021

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/physx"
	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	cmd2021.AddCommand(day5Cmd)
}

var (
	day5logger = log.With().
			Int("day", 5).
			Logger()
	day5Cmd = &cobra.Command{
		Use:   "day5",
		Short: "Day 5 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day5logger.Info().Msg("Start")
			converted := prepareday5Input()

			oceanFloor := utils.SingleSliceMap[uint, int]{}
			for _, dataTuple := range converted {
				start := dataTuple.V1
				end := dataTuple.V2
				dir := end.Copy().Sub(start)
				dirNormalized := dir.Normalized()
				switch {
				case dirNormalized[0] == 0.0:
					fallthrough
				case dirNormalized[1] == 0.0:
					day5logger.Trace().Msg("horizontal or vertical vector")
				case math.Abs(dirNormalized[0]) == math.Abs(dirNormalized[1]):
					day5logger.Trace().Msg("diagonal vector")
					dirNormalized = dirNormalized.Ceil()
				default:
					day5logger.Warn().Msg("dir vector cannot be used")
					continue
				}
				curr := start.Copy()
				for i := 0; i <= int(dir.Magnitude()); i++ {
					if curr.Copy().Sub(start).Magnitude() > dir.Magnitude() {
						break
					}
					oceanFloor.ModifyElem(func(elem *int) *int {
						if elem == nil {
							res := 1
							return &res
						}
						*elem += 1
						return elem
					}, uint(curr[0]), uint(curr[1]))
					curr.Add(dirNormalized)
				}
			}
			day5logger.Debug().Msgf("ocean floor: %s", oceanFloor)
			var dangerousAreas int
			height, _ := oceanFloor.GetHeight()
			for i := uint(0); i < height; i++ {
				row := oceanFloor.GetRow(i)
				for _, elem := range row {
					if elem.Data == nil {
						continue
					}
					if *elem.Data >= 2 {
						dangerousAreas++
					}
				}
			}

			fmt.Printf("dangerous areas: %d\n", dangerousAreas)
		},
	}
)

func prepareday5Input() []utils.Tuple[physx.Vector, physx.Vector] {
	content, err := os.ReadFile("resources/day5.txt")
	if err != nil {
		day5logger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]
	day5logger.Info().Msgf("length of input file: %d", len(input))

	converted := make([]utils.Tuple[physx.Vector, physx.Vector], len(input))
	for i := 0; i < len(input); i++ {
		vectorString := strings.Split(input[i], " -> ")
		converted[i] = utils.Tuple[physx.Vector, physx.Vector]{physx.ToVector(vectorString[0]), physx.ToVector(vectorString[1])}
	}
	day5logger.Debug().Msgf("converted input: %v", converted)

	return converted
}
