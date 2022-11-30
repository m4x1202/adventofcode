package cmd2021

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func BracketsMatch(a, b rune) bool {
	return b == GetMatchingBracket(a)
}

func GetMatchingBracket(b rune) rune {
	switch b {
	case '{':
		return '}'
	case '[':
		return ']'
	case '(':
		return ')'
	case '<':
		return '>'
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	case '>':
		return '<'
	default:
		return ' '
	}
}

var (
	bracketPoints = map[rune]uint16{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	autoCompletePoints = map[rune]uint8{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

type Chunk struct {
	Open  rune
	Close rune
	Subs  []Chunk
}

func ParseChunk(in string) (n int, c *Chunk) {
	if len(in) == 0 {
		return 0, nil
	}
	peek := rune(in[0])
	day10logger.Trace().Msgf("chunk peeked rune: %v", peek)
	switch peek {
	case '{':
		fallthrough
	case '[':
		fallthrough
	case '(':
		fallthrough
	case '<':
		chunk := Chunk{}
		chunk.Open = peek
		length := 1
		for {
			if len(in[length:]) == 0 {
				return length, &chunk
			}
			subL, sub := ParseChunk(in[length:])
			if subL == 0 {
				chunk.Close = rune(in[length])
				return length + 1, &chunk
			}
			length += subL
			chunk.Subs = append(chunk.Subs, *sub)
		}
	default:
		return 0, nil
	}
}

func (c Chunk) Verify() uint16 {
	for _, sub := range c.Subs {
		points := sub.Verify()
		if points > 0 {
			return points
		}
	}
	if !BracketsMatch(c.Open, c.Close) {
		return bracketPoints[c.Close]
	}
	return 0
}

func (c Chunk) CompleteChunk() string {
	var res string
	for _, sub := range c.Subs {
		res += sub.CompleteChunk()
	}
	if !BracketsMatch(c.Open, c.Close) {
		res += string(GetMatchingBracket(c.Open))
	}
	return res
}

func CalcAutoCompleteScore(s string) uint64 {
	var totalScore uint64
	for _, c := range s {
		totalScore *= 5
		totalScore += uint64(autoCompletePoints[c])
	}
	return totalScore
}

func init() {
	cmd2021.AddCommand(day10Cmd)

	day10Cmd.AddCommand(day10part1Cmd)
	day10Cmd.AddCommand(day10part2Cmd)
}

var (
	day10logger = log.With().
			Int("day", 10).
			Logger()
	day10Cmd = &cobra.Command{
		Use:   "day10",
		Short: "Day 10 Challenge",
	}
	day10part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 10 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day10part1logger := day10logger.With().
				Int("part", 1).
				Logger()
			day10part1logger.Info().Msg("Start")
			converted := prepareday10Input()

			var totalPoints uint32
			for _, chunk := range converted {
				totalPoints += uint32(chunk.Verify())
			}

			fmt.Printf("%d\n", totalPoints)
		},
	}
	day10part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 10 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day10part2logger := day10logger.With().
				Int("part", 2).
				Logger()
			day10part2logger.Info().Msg("Start")
			converted := prepareday10Input()

			day10part2logger.Debug().Msgf("total num chunks %d", len(converted))
			filteredChunks := make([]*Chunk, 0, len(converted)/2)
			for _, chunk := range converted {
				if chunk.Verify() == 0 {
					filteredChunks = append(filteredChunks, chunk)
				}
			}
			day10part2logger.Debug().Msgf("total num filtered chunks %d", len(filteredChunks))
			autoCompleteScores := make([]uint64, 0, len(filteredChunks))
			for _, chunk := range filteredChunks {
				necessaryToComplete := chunk.CompleteChunk()
				day10part2logger.Debug().Msgf("chunk would get completed by adding '%s'", necessaryToComplete)
				autoCompleteScores = append(autoCompleteScores, CalcAutoCompleteScore(necessaryToComplete))
			}
			sort.Slice(autoCompleteScores, func(i, j int) bool { return autoCompleteScores[i] < autoCompleteScores[j] })
			day10part2logger.Debug().Msgf("auto complete scores %v", autoCompleteScores)
			fmt.Printf("%d\n", autoCompleteScores[len(autoCompleteScores)/2])
		},
	}
)

func prepareday10Input() []*Chunk {
	content, err := os.ReadFile("resources/day10.txt")
	if err != nil {
		day10logger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	day10logger.Info().Msgf("length of input file: %d", len(input))
	day10logger.Debug().Msgf("plain input: %v", input)

	converted := make([]*Chunk, 0, len(input))
	for _, rawChunk := range input {
		_, chunk := ParseChunk(rawChunk)
		converted = append(converted, chunk)
	}

	return converted
}
