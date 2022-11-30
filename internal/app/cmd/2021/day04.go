package cmd2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type BingoCardElem struct {
	string
	bool
}
type BingoCardRow [5]*BingoCardElem

func (r *BingoCardRow) CheckBingo() bool {
	for _, elem := range r {
		if !elem.bool {
			return false
		}
	}
	return true
}

type BingoCardCol [5]*BingoCardElem

func (c *BingoCardCol) CheckBingo() bool {
	for _, elem := range c {
		if !elem.bool {
			return false
		}
	}
	return true
}

type BingoCard struct {
	rows      [5]BingoCardRow
	cols      [5]BingoCardCol
	completed bool
	score     int
}

func ToBingoCard(in string) BingoCard {
	day4logger.Debug().Msgf("bingo card conversion input: %v", in)
	res := BingoCard{}
	rows := strings.Split(in, "\n")
	day4logger.Trace().Msgf("bingo card num rows: %d", len(rows))
	for i := 0; i < len(rows); i++ {
		elems := strings.Split(rows[i], " ")
		day4logger.Trace().Msgf("bingo card num cols: %d", len(elems))
		var colNum int
		for j := 0; j < len(elems); j++ {
			if elems[j] == "" {
				continue
			}
			elem := &BingoCardElem{elems[j], false}
			res.cols[colNum][i] = elem
			res.rows[i][colNum] = elem
			colNum++
		}
	}
	return res
}

func (c *BingoCard) NewNumber(num string) bool {
	for _, row := range c.rows {
		for _, elem := range row {
			if elem.string == num {
				elem.bool = true
			}
		}
	}
	if c.CheckBingo() {
		c.completed = true
		var sumUnmarked int
		for _, row := range c.rows {
			for _, elem := range row {
				if !elem.bool {
					amount, _ := strconv.Atoi(elem.string)
					sumUnmarked += amount
				}
			}
		}
		finishingNum, _ := strconv.Atoi(num)
		c.score = finishingNum * sumUnmarked
		return true
	}
	return false
}

func (c *BingoCard) CheckBingo() bool {
	for _, col := range c.cols {
		if col.CheckBingo() {
			return true
		}
	}
	for _, row := range c.rows {
		if row.CheckBingo() {
			return true
		}
	}
	return false
}

func init() {
	cmd2021.AddCommand(day4Cmd)
}

var (
	day4logger = log.With().
			Int("day", 4).
			Logger()
	day4Cmd = &cobra.Command{
		Use:   "day4",
		Short: "Day 4 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day4logger.Info().Msg("Start")

			input, _ := os.ReadFile("resources/day4.txt")

			split := strings.Split(string(input), "\n\n")
			randomNumbers := strings.Split(split[0], ",")
			bingoCards := parseBingoCard(split[1:])
			completedCards := make([]int, 0, len(bingoCards))

			for _, num := range randomNumbers {
				for i := 0; i < len(bingoCards); i++ {
					if bingoCards[i].completed {
						continue
					}
					if bingoCards[i].NewNumber(num) {
						day4logger.Debug().Int("card", i).Int("score", bingoCards[i].score).Msg("BINGO!")
						if len(completedCards) == 0 {
							fmt.Printf("first card %d completed with score: %d\n", i, bingoCards[i].score)
						}
						completedCards = append(completedCards, i)
						if len(completedCards) == len(bingoCards) {
							day4logger.Info().Msg("all cards completed")
							fmt.Printf("last card %d completed with score: %d\n", i, bingoCards[i].score)
							return
						}
					}
				}
			}
		},
	}
)

func parseBingoCard(in []string) []BingoCard {
	res := make([]BingoCard, len(in))
	for i := 0; i < len(in); i++ {
		res[i] = ToBingoCard(in[i])
	}
	return res
}
