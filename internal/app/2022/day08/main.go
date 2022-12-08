package day08

import (
	"fmt"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "08"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func ExecutePart(p uint8) {
	preparedInput := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		part1Func(preparedInput)
	case 2:
		part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
}

func part1Func(preparedInput Forest) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	treeVisibleMap := preparedInput.CheckTreesVisible2()
	for _, col := range treeVisibleMap {
		for _, elem := range col {
			if elem {
				puzzleAnswer++
			}
		}
	}

	fmt.Printf("trees visible from outside: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(preparedInput Forest) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for x := range preparedInput {
		for y := range preparedInput[x] {
			var n, e, s, w uint64
			heightTree := preparedInput[x][y]
			for xW := x - 1; xW >= 0; xW-- {
				w++
				if preparedInput[xW][y] >= heightTree {
					break
				}
			}
			for xE := x + 1; xE < len(preparedInput[0]); xE++ {
				e++
				if preparedInput[xE][y] >= heightTree {
					break
				}
			}
			for yS := y + 1; yS < len(preparedInput[x]); yS++ {
				s++
				if preparedInput[x][yS] >= heightTree {
					break
				}
			}
			for yN := y - 1; yN >= 0; yN-- {
				n++
				if preparedInput[x][yN] >= heightTree {
					break
				}
			}
			scenicScore := n * e * s * w
			partLogger.Debug().Int("x", x).Int("y", y).Msgf("tree scenic score: %d", scenicScore)
			if scenicScore > puzzleAnswer {
				puzzleAnswer = scenicScore
			}
		}
	}

	fmt.Printf("max scenic score: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) Forest {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	height := len(input)
	width := len(input[0])
	converted := utils.NewMap[uint8](width, height)
	for y := range converted {
		row := input[y]
		for x := range converted {
			converted[x][y] = uint8(row[x] - '0')
		}
	}

	return Forest(converted)
}

type Forest utils.Map[uint8]

func (f Forest) String() string {
	var builder strings.Builder
	for y := range f[0] {
		for x := range f {
			builder.WriteString(cast.ToString(f[x][y]))
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (f Forest) CheckTreesVisible2() utils.Map[bool] {
	treeVisibleMap := utils.NewMap[bool](len(f), len(f[0]))
	for x := range f {
		treeVisibleMap[x][0] = true
		treeVisibleMap[x][len(f)-1] = true
	}
	for y := range f[0] {
		treeVisibleMap[0][y] = true
		treeVisibleMap[len(f[0])-1][y] = true
	}

	for x := 0; x < len(f)-1; x++ {
		highestN := uint8(0)
		for y := 0; y < len(f[x])-1; y++ {
			if f[x][y] > highestN {
				treeVisibleMap[x][y] = true
				highestN = f[x][y]
			}
		}
	}
	for y := 0; y < len(f[0])-1; y++ {
		highestW := uint8(0)
		for x := 0; x < len(f)-1; x++ {
			if f[x][y] > highestW {
				treeVisibleMap[x][y] = true
				highestW = f[x][y]
			}
		}
	}

	for x := 0; x < len(f)-1; x++ {
		highestS := uint8(0)
		for y := len(f[x]) - 1; y > 0; y-- {
			if f[x][y] > highestS {
				treeVisibleMap[x][y] = true
				highestS = f[x][y]
			}
		}
	}
	for y := 0; y < len(f[0])-1; y++ {
		highestE := uint8(0)
		for x := len(f) - 1; x > 0; x-- {
			if f[x][y] > highestE {
				treeVisibleMap[x][y] = true
				highestE = f[x][y]
			}
		}
	}
	return treeVisibleMap
}

func (f Forest) CheckTreesVisible() utils.Map[uint8] {
	treeVisibleMap := utils.NewMap[uint8](len(f), len(f[0]))
	for i := uint8(9); i > 0; i-- {
		tmpMap := f.CheckTreesWithHeight(i)
		partLogger.Trace().Msgf("%v", tmpMap)
		tmpHeightMap := ConnectHeightMap(tmpMap, i)
		for x := range treeVisibleMap {
			for y := range treeVisibleMap[x] {
				if tmpHeightMap[x][y] > treeVisibleMap[x][y] {
					treeVisibleMap[x][y] = tmpHeightMap[x][y]
				}
			}
		}
	}
	// for i := range f[0] {
	// 	treeVisibleMap[0][i] = true
	// 	treeVisibleMap[len(f[0])-1][i] = true
	// }
	// for j := range f {
	// 	treeVisibleMap[j][0] = true
	// 	treeVisibleMap[j][len(f)-1] = true
	// }
	// for i, j := 1, len(f)-2; i < j; i, j = i+1, j-1 {
	// 	for k, l := 1, len(f[0])-2; k < l; k, l = k+1, l-1 {
	// 	}
	// }
	return treeVisibleMap
}

func (f Forest) CheckTreesWithHeight(height uint8) utils.Map[bool] {
	partLogger.Debug().Msgf("Checking trees with height %d", height)
	treeMap := utils.NewMap[bool](len(f), len(f[0]))
	for x := range f {
		for y := range f[x] {
			if f[x][y] == height {
				treeMap[x][y] = true
			}
		}
	}
	return treeMap
}

func ConnectHeightMap(in utils.Map[bool], height uint8) utils.Map[uint8] {
	treeMap := utils.NewMap[uint8](len(in), len(in[0]))
	for x, col := range in {
		min_y := -1
		max_y := -1
		for y, elem := range col {
			if elem {
				if min_y == -1 {
					min_y = y
					treeMap[x][y] = height
				} else {
					max_y = y
				}
			}
		}
		if max_y != -1 {
			for y := min_y; y < max_y; y++ {
				treeMap[x][y] = height
			}
		}
	}
	for y := range in[0] {
		min_x := -1
		max_x := -1
		for x := range in {
			if in[x][y] {
				if min_x == -1 {
					min_x = x
					treeMap[x][y] = height
				} else {
					max_x = x
				}
			}
		}
		if max_x != -1 {
			for x := min_x; x < max_x; x++ {
				treeMap[x][y] = height
			}
		}
	}
	return treeMap
}
