package cmd2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

func init() {
	cmd2021.AddCommand(day2Cmd)
}

var (
	day2logger = log.With().
			Int("day", 2).
			Logger()
	day2Cmd = &cobra.Command{
		Use:   "day2",
		Short: "Day 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day2logger.Info().Msg("Start")

			content, err := os.ReadFile("resources/day2.txt")
			if err != nil {
				day2logger.Fatal().Err(err).Send()
			}

			input := strings.Split(string(content), "\n")
			day2logger.Info().Msgf("length of input file: %d", len(input))

			type command struct {
				dir  Direction
				pace int
			}

			converted := make([]command, len(input))
			for i := 0; i < len(input); i++ {
				tmp := strings.Split(input[i], " ")
				dir := Direction(tmp[0])
				pace, _ := strconv.Atoi(tmp[1])
				converted[i] = command{dir, pace}
				day2logger.Trace().Int("index", i).Msgf("converted: %v", converted[i])
			}

			var horizontalPos, depth, aim int
			for _, command := range converted {
				switch {
				case command.dir == Forward:
					horizontalPos += command.pace
					depth += aim * command.pace
				case command.dir == Down:
					aim += command.pace
				case command.dir == Up:
					aim -= command.pace
				}
			}
			day2logger.Info().Msgf("horizontal position: %d", horizontalPos)
			day2logger.Info().Msgf("depth: %d", depth)

			log.Info().Msgf("distance: %d", depth*horizontalPos)
			fmt.Println("distance:", depth*horizontalPos)
		},
	}
)
