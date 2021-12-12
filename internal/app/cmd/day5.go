package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/physx"
	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day5Cmd)
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
			converted := prepareDay5Input()

			oceanFloor := utils.Map{}
			for _, dataTuple := range converted {
				vectorTuple := dataTuple.GetVector()
				var horizontal, vertical int
				if vectorTuple[1][0] == 0.0 {
					day5logger.Trace().Msg("vertical vector")
					vertical++
				}
				if vectorTuple[1][1] == 0.0 {
					day5logger.Trace().Msg("horizontal vector")
					horizontal++
				}
				if vertical == 0 && horizontal == 0 {
					day5logger.Trace().Msg("diagonal vector")
					continue
				}

				origin := vectorTuple[0]
				dirNormalized := vectorTuple[1].Normalized()
				for i := 0; i <= int(vectorTuple[1].Magnitude()); i++ {
					oceanFloor.ModifyElem(func(elem interface{}) interface{} {
						if elem == nil {
							return 1
						}
						return cast.ToInt(elem) + 1
					}, int(origin[0]), int(origin[1]))
					origin.Add(dirNormalized)
				}
			}
			var dangerousAreas int
			for i := 0; i < oceanFloor.Height; i++ {
				row := oceanFloor.GetRow(i)
				for _, elem := range row {
					vents := cast.ToInt(elem)
					if vents >= 2 {
						dangerousAreas++
					}
				}
			}

			fmt.Printf("dangerous areas: %d\n", dangerousAreas)
		},
	}
)

func prepareDay5Input() []utils.Tuple {
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
		b.Sub(a)
		converted[i] = utils.Tuple{a, b}
	}
	day5logger.Debug().Msgf("converted input: %v", converted)

	return converted
}
