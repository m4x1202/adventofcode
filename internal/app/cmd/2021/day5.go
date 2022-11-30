package cmd2021

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/physx"
	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
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

			oceanFloor := utils.Map{}
			for _, dataTuple := range converted {
				start := (dataTuple[0]).(*physx.Vector)
				end := (dataTuple[1]).(*physx.Vector)
				dir := end.Copy().Sub(*start)
				dirNormalized := dir.Normalized()
				switch {
				case (*dirNormalized)[0] == 0.0:
					fallthrough
				case (*dirNormalized)[1] == 0.0:
					day5logger.Trace().Msg("horizontal or vertical vector")
				case math.Abs((*dirNormalized)[0]) == math.Abs((*dirNormalized)[1]):
					day5logger.Trace().Msg("diagonal vector")
					dirNormalized = dirNormalized.Ceil()
				default:
					day5logger.Warn().Msg("dir vector cannot be used")
					continue
				}
				curr := start.Copy()
				for i := 0; i <= int(dir.Magnitude()); i++ {
					if curr.Copy().Sub(*start).Magnitude() > dir.Magnitude() {
						break
					}
					oceanFloor.ModifyElem(func(elem interface{}) interface{} {
						if elem == nil {
							return 1
						}
						return cast.ToInt(elem) + 1
					}, uint((*curr)[0]), uint((*curr)[1]))
					curr.Add(*dirNormalized)
				}
			}
			day5logger.Debug().Msgf("ocean floor: %s", oceanFloor)
			var dangerousAreas int
			for i := uint(0); i < oceanFloor.Height; i++ {
				row := oceanFloor.GetRow(i)
				for _, elem := range row {
					if elem == nil {
						continue
					}
					vents := cast.ToInt(elem.Data)
					if vents >= 2 {
						dangerousAreas++
					}
				}
			}

			fmt.Printf("dangerous areas: %d\n", dangerousAreas)
		},
	}
)

func prepareday5Input() []utils.Tuple {
	content, err := os.ReadFile("resources/day5.txt")
	if err != nil {
		day5logger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]
	day5logger.Info().Msgf("length of input file: %d", len(input))

	converted := make([]utils.Tuple, len(input))
	for i := 0; i < len(input); i++ {
		vectorString := strings.Split(input[i], " -> ")
		a, _ := physx.FromString(vectorString[0])
		b, _ := physx.FromString(vectorString[1])
		converted[i] = utils.Tuple{a, b}
	}
	day5logger.Debug().Msgf("converted input: %v", converted)

	return converted
}
