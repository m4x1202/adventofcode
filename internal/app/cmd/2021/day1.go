package cmd2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	cmd2021.AddCommand(day1Cmd)
}

var (
	day1logger = log.With().
			Int("day", 1).
			Logger()
	day1Cmd = &cobra.Command{
		Use:   "day1",
		Short: "Day 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day1logger.Info().Msg("Start")

			content, err := os.ReadFile("resources/day1.txt")
			if err != nil {
				day1logger.Fatal().Err(err).Send()
			}

			input := strings.Split(strings.TrimSpace(string(content)), "\n")
			day1logger.Info().Msgf("length of input file: %d", len(input))

			converted := make([]int, len(input))
			for i := 0; i < len(input); i++ {
				converted[i], _ = strconv.Atoi(input[i])
				day1logger.Trace().Int("index", i).Msgf("converted: %d", converted[i])
			}

			sum := make([]int, len(converted)-(len(converted)%3))
			for i := 0; i < len(sum); i++ {
				sum[i] += converted[i]
				sum[i] += converted[i+1]
				sum[i] += converted[i+2]
			}

			var inc_count, dec_count, eq_count int
			for i := 1; i < len(sum); i++ {
				day1logger.Trace().Int("index", i).Msgf("sum: %d", sum[i])
				switch {
				case sum[i-1] == sum[i]:
					day1logger.Trace().Int("index", i).Msg("equal")
					eq_count++
				case sum[i-1] > sum[i]:
					day1logger.Trace().Int("index", i).Msg("lower")
					dec_count++
				case sum[i-1] < sum[i]:
					day1logger.Trace().Int("index", i).Msgf("higher")
					inc_count++
				}
			}
			log.Log().Msgf("Equal Count: %d", eq_count)
			log.Log().Msgf("Decrease Count: %d", dec_count)
			log.Info().Msgf("Increase Count: %d", inc_count)
			fmt.Println("Increase Count:", inc_count)
		},
	}
)
