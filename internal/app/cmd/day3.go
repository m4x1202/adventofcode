package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Bit uint8

func ToUint16(in [12]Bit) uint16 {
	var res uint16
	multiplier := uint16(1)
	for i := len(in) - 1; i >= 0; i-- {
		res += uint16(in[i]) * multiplier
		multiplier *= 2
	}
	return res
}

func init() {
	rootCmd.AddCommand(day3Cmd)

	day3Cmd.AddCommand(day3part1Cmd)
	day3Cmd.AddCommand(day3part2Cmd)
}

var (
	day3logger = log.With().
			Int("day", 3).
			Logger()
	day3Cmd = &cobra.Command{
		Use:   "day3",
		Short: "Day 3 Challenge",
	}
	day3part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 3 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day3part1logger := day3logger.With().
				Int("part", 1).
				Logger()
			day3part1logger.Info().Msg("Start")
			converted := prepareDay3Input()

			var gamma [12]Bit
			for i := 0; i < 12; i++ {
				var count int
				for j := 0; j < len(converted); j++ {
					if converted[j][i] == 1 {
						count++
					} else {
						count--
					}
				}
				if count > 0 {
					gamma[i] = 1
				} else {
					gamma[i] = 0
				}
			}

			var epsilon [12]Bit
			for i := 0; i < len(epsilon); i++ {
				epsilon[i] = (gamma[i] + 1) % 2
			}

			day3part1logger.Info().Msgf("gamma rate: %d", ToUint16(gamma))
			day3part1logger.Info().Msgf("epsilon rate: %d", ToUint16(epsilon))

			log.Info().Msgf("power consumption: %d", uint64(ToUint16(gamma))*uint64(ToUint16(epsilon)))
			fmt.Println("power consumption:", uint64(ToUint16(gamma))*uint64(ToUint16(epsilon)))
		},
	}
	day3part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 3 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day3part2logger := day3logger.With().
				Int("part", 2).
				Logger()
			day3part2logger.Info().Msg("Start")
			converted := prepareDay3Input()

			oxygenGenRating := calcOxygenGenRating(converted, day3part2logger)
			co2ScrubRating := calcCO2ScrubRating(converted, day3part2logger)
			day3part2logger.Info().Msgf("oxygen generator rating: %d", oxygenGenRating)
			day3part2logger.Info().Msgf("co2 scrubber rating: %d", co2ScrubRating)

			log.Info().Msgf("life support rating: %d", uint64(oxygenGenRating)*uint64(co2ScrubRating))
			fmt.Println("life support rating:", uint64(oxygenGenRating)*uint64(co2ScrubRating))
		},
	}
)

func prepareDay3Input() [][12]Bit {
	content, err := os.ReadFile("resources/day3.txt")
	if err != nil {
		day3logger.Fatal().Err(err).Send()
	}

	input := bytes.Split(content, []byte("\n"))
	input = input[:len(input)-1]
	day3logger.Info().Msgf("length of input file: %d", len(input))

	converted := make([][12]Bit, len(input))
	for i := 0; i < len(input); i++ {
		var row [12]Bit
		for j := 0; j < 12; j++ {
			if input[i][j] == byte(48) { //Byte 48 equals to string '0'
				row[j] = 0
			} else {
				row[j] = 1
			}
		}
		converted[i] = row
		day3logger.Trace().Int("index", i).Msgf("converted: %v", converted[i])
	}
	return converted
}

func calcOxygenGenRating(converted [][12]Bit, logger zerolog.Logger) uint16 {
	filtered := make([][12]Bit, len(converted))
	copy(filtered, converted)
	for i := 0; i < 12; i++ {
		var count int
		for j := 0; j < len(filtered); j++ {
			if filtered[j][i] == 1 {
				count++
			} else {
				count--
			}
		}
		logger.Debug().Int("index", i).Msgf("count: %d", count)
		var selected Bit
		if count >= 0 {
			selected = 1
		} else {
			selected = 0
		}
		new := make([][12]Bit, 0, len(filtered))
		for j := 0; j < len(filtered); j++ {
			if filtered[j][i] == selected {
				new = append(new, filtered[j])
			}
		}
		filtered = new
		logger.Debug().Int("index", i).Int("selected", int(selected)).Msgf("%v", filtered)
		if len(filtered) == 1 {
			return ToUint16(filtered[0])
		}
	}
	return 0
}

func calcCO2ScrubRating(converted [][12]Bit, logger zerolog.Logger) uint16 {
	filtered := make([][12]Bit, len(converted))
	copy(filtered, converted)
	for i := 0; i < 12; i++ {
		var count int
		for j := 0; j < len(filtered); j++ {
			if filtered[j][i] == 1 {
				count++
			} else {
				count--
			}
		}
		logger.Debug().Int("index", i).Msgf("count: %d", count)
		var selected Bit
		if count >= 0 {
			selected = 1
		} else {
			selected = 0
		}
		new := make([][12]Bit, 0, len(filtered))
		for j := 0; j < len(filtered); j++ {
			if filtered[j][i] != selected {
				new = append(new, filtered[j])
			}
		}
		filtered = new
		logger.Debug().Int("index", i).Int("selected", int(selected)).Msgf("%v", filtered)
		if len(filtered) == 1 {
			return ToUint16(filtered[0])
		}
	}
	return 0
}
