package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// SevenSegmentNumber are by definition alphabet sorted
type SevenSegmentNumber string

const (
	ZERO  SevenSegmentNumber = "abcefg"
	ONE   SevenSegmentNumber = "cf"
	TWO   SevenSegmentNumber = "acdeg"
	THREE SevenSegmentNumber = "acdfg"
	FOUR  SevenSegmentNumber = "bcdf"
	FIVE  SevenSegmentNumber = "abdfg"
	SIX   SevenSegmentNumber = "abdefg"
	SEVEN SevenSegmentNumber = "acf"
	EIGHT SevenSegmentNumber = "abcdefg"
	NINE  SevenSegmentNumber = "abcdfg"
)

func (n SevenSegmentNumber) ToUint8() uint8 {
	switch n {
	case ZERO:
		return 0
	case ONE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	}
	return 10
}

func (a SevenSegmentNumber) Minus(b SevenSegmentNumber) SevenSegmentNumber {
	cp := string(a)
	for _, r := range b {
		cp = strings.Replace(cp, string(r), "", 1)
	}
	return SevenSegmentNumber(cp)
}

var (
	easyNumbersLengths = []int{2, 3, 4, 7}
	decoderCandidates  = map[rune][]rune{}
)

type SegmentDecoder map[rune]rune

func (d SegmentDecoder) String() string {
	var res strings.Builder
	res.WriteString("{ ")
	for k, v := range d {
		res.WriteRune(k)
		res.WriteRune(':')
		res.WriteRune(v)
		res.WriteRune(' ')
	}
	res.WriteString("}")
	return res.String()
}

func (d *SegmentDecoder) Train(n SevenSegmentNumber) *SegmentDecoder {
	switch len(n) {
	case 2:
		asRunes := []rune(n)
		(*d)['c'] = asRunes[0]
		(*d)['f'] = asRunes[1]
		if d.Encode(ONE) != n {
			day8logger.Error().Msg("Training ONE unsuccessful")
			return nil
		}
	case 3:
		remain := n.Minus(d.Encode(ONE))
		if len(remain) != 1 {
			day8logger.Warn().Msgf("not exactly 1 rune remaining after minus: %s", string(remain))
			return d
		}
		(*d)['a'] = rune(remain[0])
		if d.Encode(SEVEN) != n {
			day8logger.Error().Msgf("Training SEVEN unsuccessful: %s", d.Encode(SEVEN))
			return nil
		}
	case 4:
		remain := n.Minus(d.Encode(ONE))
		for _, r := range remain {
			decoderCandidates[r] = append(decoderCandidates[r], []rune{'b', 'd'}...)
		}
		day8logger.Debug().Msgf("decoder candidates: %v", decoderCandidates)
	case 5:
		remain := n.Minus(d.Encode(SEVEN))
		if len(remain) > 2 {
			return d
		}
		if len(remain) < 2 {
			day8logger.Warn().Msgf("not exactly 2 runes remaining after minus: %s", string(remain))
			return nil
		}
		for _, r := range remain {
			if _, ok := decoderCandidates[r]; ok {
				(*d)['d'] = r
				delete(decoderCandidates, r)
			} else {
				(*d)['g'] = r
			}
		}
		if len(decoderCandidates) > 1 {
			day8logger.Warn().Msg("Still too many candidates")
			return nil
		}
		for k := range decoderCandidates {
			(*d)['b'] = k
		}
		decoderCandidates = make(map[rune][]rune)
		if d.Encode(THREE) != n {
			day8logger.Error().Msgf("Training THREE unsuccessful: %s", d.Encode(THREE))
			return nil
		}
		for _, c := range EIGHT {
			found := false
			for _, v := range *d {
				if v == c {
					found = true
				}
			}
			if !found {
				(*d)['e'] = c
				break
			}
		}
	case 6:
		//Correct assumption from ONE
		decodedNum := d.Decode(n).ToUint8()
		if decodedNum == 10 {
			(*d)['c'], (*d)['f'] = (*d)['f'], (*d)['c']
		}
	case 7:
		if d.Encode(EIGHT) != n {
			day8logger.Error().Msgf("Training EIGHT unsuccessful: %s", d.Encode(EIGHT))
			return nil
		}
	default:
		return d
	}
	day8logger.Trace().Msgf("trained decoder: %v", *d)
	return d
}

func (d *SegmentDecoder) Encode(n SevenSegmentNumber) SevenSegmentNumber {
	var builder strings.Builder
	for _, r := range n {
		builder.WriteRune((*d)[r])
	}
	splitDigit := strings.Split(builder.String(), "")
	sort.Strings(splitDigit)
	return SevenSegmentNumber(strings.Join(splitDigit, ""))
}

func (d *SegmentDecoder) Decode(n SevenSegmentNumber) SevenSegmentNumber {
	var builder strings.Builder
	for _, r := range n {
		for k, v := range *d {
			if v == r {
				builder.WriteRune(k)
			}
		}
	}
	splitDigit := strings.Split(builder.String(), "")
	sort.Strings(splitDigit)
	return SevenSegmentNumber(strings.Join(splitDigit, ""))
}

func init() {
	rootCmd.AddCommand(day8Cmd)

	day8Cmd.AddCommand(day8part1Cmd)
	day8Cmd.AddCommand(day8part2Cmd)
}

var (
	day8logger = log.With().
			Int("day", 8).
			Logger()
	day8Cmd = &cobra.Command{
		Use:   "day8",
		Short: "Day 8 Challenge",
	}
	day8part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 8 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day8part1logger := day8logger.With().
				Int("part", 1).
				Logger()
			day8part1logger.Info().Msg("Start")
			converted := prepareday8Input()

			var count uint16
			for _, line := range converted {
				outputNumberSegments := line[1].([]SevenSegmentNumber)
				day8part1logger.Trace().Msg("-------------------------------------------")
				for _, output := range outputNumberSegments {
					day8part1logger.Trace().Msgf("output number: %s", output)
					if utils.Contains(easyNumbersLengths, len(output)) {
						count++
					}
				}
				day8part1logger.Trace().Msgf("curr count: %d", count)
			}

			fmt.Printf("%d\n", count)
		},
	}
	day8part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 8 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day8part2logger := day8logger.With().
				Int("part", 2).
				Logger()
			day8part2logger.Info().Msg("Start")
			converted := prepareday8Input()

			var addedOutputs uint32
			for _, line := range converted {
				decoder := &SegmentDecoder{}
				signalNumberSegments := line[0].([]SevenSegmentNumber)
				outputNumberSegments := line[1].([]SevenSegmentNumber)

				for _, signal := range signalNumberSegments {
					log.Trace().Msgf("training on %s", signal)
					decoder.Train(signal)
				}

				decodedOutput := make([]SevenSegmentNumber, 0, len(outputNumberSegments))
				for _, output := range outputNumberSegments {
					decodedOutput = append(decodedOutput, decoder.Decode(output))
				}
				log.Debug().Msgf("decoded output: %v", decodedOutput)

				decodedOutputAsNumbers := make([]uint8, 0, len(outputNumberSegments))
				for _, output := range decodedOutput {
					decodedOutputAsNumbers = append(decodedOutputAsNumbers, output.ToUint8())
				}
				log.Debug().Msgf("decoded output as numbers: %v", decodedOutputAsNumbers)

				finalNumber := SliceToNumber(decodedOutputAsNumbers)
				log.Debug().Msgf("decoded output as number: %d", finalNumber)

				addedOutputs += uint32(finalNumber)
			}
			fmt.Printf("%d\n", addedOutputs)
		},
	}
)

func SliceToNumber(in []uint8) uint16 {
	decimal := uint16(1)
	var res uint16
	for i := len(in) - 1; i >= 0; i-- {
		res += uint16(in[i]) * decimal
		decimal *= 10
	}
	return res
}

func prepareday8Input() []utils.Tuple {
	content, err := os.ReadFile("resources/day8.txt")
	if err != nil {
		day8logger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	day8logger.Info().Msgf("length of input file: %d", len(input))
	day8logger.Debug().Msgf("plain input: %v", input)

	converted := make([]utils.Tuple, len(input))
	for i := 0; i < len(input); i++ {
		vectorString := strings.Split(input[i], " | ")
		signal := strings.Split(vectorString[0], " ")
		var signalConverted []SevenSegmentNumber
		for _, digit := range signal {
			splitDigit := strings.Split(digit, "")
			sort.Strings(splitDigit)
			signalConverted = append(signalConverted, SevenSegmentNumber(strings.Join(splitDigit, "")))
		}
		sort.Slice(signalConverted, func(i, j int) bool { return len(signalConverted[i]) < len(signalConverted[j]) })
		output := strings.Split(vectorString[1], " ")
		var outputConverted []SevenSegmentNumber
		for _, digit := range output {
			splitDigit := strings.Split(digit, "")
			sort.Strings(splitDigit)
			outputConverted = append(outputConverted, SevenSegmentNumber(strings.Join(splitDigit, "")))
		}
		converted[i] = utils.Tuple{signalConverted, outputConverted}
	}

	day8logger.Debug().Msgf("converted input: %v", converted)

	return converted
}
