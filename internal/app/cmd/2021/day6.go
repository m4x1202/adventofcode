package cmd2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

type Fishschool struct {
	school []uint64
	size   uint64
}

func NewFishschool() *Fishschool {
	school := Fishschool{}
	school.size = 0
	school.school = make([]uint64, 9)
	return &school
}

func (s *Fishschool) GetSize() string {
	return strconv.FormatUint(s.size, 10)
}

func (s *Fishschool) AddFish(timer uint8, amount uint64) {
	s.school[timer] += amount
	s.size += amount
}

func (s *Fishschool) Day() {
	var birthGroup uint64
	birthGroup, s.school = s.school[0], s.school[1:]
	s.school = append(s.school, 0)
	if len(s.school) != 9 {
		day6logger.Warn().Msg("Fischschool has incorrect length")
	}
	s.school[6] += birthGroup
	s.AddFish(8, birthGroup)
	day6logger.Trace().Msgf("fishschool: %v", s.school)
}

func init() {
	cmd2021.AddCommand(day6Cmd)
}

var (
	day6logger = log.With().
			Int("day", 6).
			Logger()
	day6Cmd = &cobra.Command{
		Use:   "day6",
		Short: "Day 6 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				return
			}
			days := cast.ToInt(args[0])

			day6logger.Info().Msg("Start")
			converted := prepareday6Input()

			for day := 1; day <= days; day++ {
				converted.Day()
				day6logger.Trace().Msgf("fish after %d days: %s", day, converted.GetSize())
			}

			fmt.Printf("fish after %d days: %s\n", days, converted.GetSize())
		},
	}
)

func prepareday6Input() *Fishschool {
	content, err := os.ReadFile("resources/day6.txt")
	if err != nil {
		day6logger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), ",")
	day6logger.Info().Msgf("length of input file: %d", len(input))
	day6logger.Debug().Msgf("plain input: %v", input)

	converted := NewFishschool()
	for i := 0; i < len(input); i++ {
		fish := cast.ToUint8(input[i])
		converted.AddFish(fish, 1)
	}
	day6logger.Debug().Msgf("converted input: %v", converted)

	return converted
}
