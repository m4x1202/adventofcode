package day05

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
	DAY = "05"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func ExecutePart(p uint8) {
	switch p {
	case 1:
		partLogger = dayLogger.With().
			Int("part", 1).
			Logger()
		preparedInput1, preparedInput2 := prepareInput(readPuzzleInput())
		part1Func(preparedInput1, preparedInput2)
	case 2:
		partLogger = dayLogger.With().
			Int("part", 2).
			Logger()
		preparedInput1, preparedInput2 := prepareInput(readPuzzleInput())
		part2Func(preparedInput1, preparedInput2)
	default:
		panic("part does not exist")
	}
}

func part1Func(bay CargoBay, steps []ProcedureStep) string {
	partLogger.Info().Msg("Start")
	var puzzleAnswer string

	for _, step := range steps {
		bay.ApplyProcedureStep(step)
	}
	for _, topCrate := range bay {
		if topCrateRune, ok := topCrate.Peek(); ok {
			puzzleAnswer += string(topCrateRune)
		}
	}
	fmt.Printf("top crates: %s\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(bay CargoBay, steps []ProcedureStep) string {
	partLogger.Info().Msg("Start")
	var puzzleAnswer string

	for _, step := range steps {
		bay.ApplyProcedureStepDirectly(step)
	}
	for _, topCrate := range bay {
		if topCrateRune, ok := topCrate.Peek(); ok {
			puzzleAnswer += string(topCrateRune)
		}
	}
	fmt.Printf("top crates: %s\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) (CargoBay, []ProcedureStep) {
	input := strings.Split(rawInput, "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	var cargoBayStrings []string
	var procedureSteps []ProcedureStep
	var secondHalf bool
	for _, s := range input {
		if s == "" {
			secondHalf = true
			continue
		}
		if !secondHalf {
			cargoBayStrings = append(cargoBayStrings, s)
		} else {
			procedureSteps = append(procedureSteps, ParseProcedureStep(s))
		}
	}

	return ParseCargoBay(cargoBayStrings), procedureSteps
}

type CargoBay []utils.Stack[rune]

func ParseCargoBay(in []string) CargoBay {
	lastLine := in[len(in)-1]
	lastLine = strings.TrimSpace(lastLine)
	splitLastLine := strings.Split(lastLine, "   ")
	numCrateStacks := cast.ToUint8(splitLastLine[len(splitLastLine)-1])
	cargoBay := make([]utils.Stack[rune], numCrateStacks)

	for _, line := range in[:len(in)-1] {
		cargoBaySlice := utils.ChunkSlice[[][]rune]([]rune(line), 4)
		for i, crateLayer := range cargoBaySlice {
			possibleCrate := strings.TrimSpace(string(crateLayer))
			if possibleCrate == "" {
				continue
			}
			cargoBay[i].Push(crateLayer[1])
		}
	}
	for _, crateStack := range cargoBay {
		utils.Reverse(crateStack)
	}

	return cargoBay
}

func (b *CargoBay) ApplyProcedureStep(step ProcedureStep) {
	for i := 0; i < int(step.Amount); i++ {
		if crate, ok := (*b)[step.FromStack-1].Pop(); ok {
			(*b)[step.ToStack-1].Push(crate)
		}
	}
}

func (b *CargoBay) ApplyProcedureStepDirectly(step ProcedureStep) {
	tmp := make(utils.Stack[rune], step.Amount)
	for i := 0; i < int(step.Amount); i++ {
		if crate, ok := (*b)[step.FromStack-1].Pop(); ok {
			tmp.Push(crate)
		}
	}
	for i := 0; i < int(step.Amount); i++ {
		if crate, ok := tmp.Pop(); ok {
			(*b)[step.ToStack-1].Push(crate)
		}
	}
}

type ProcedureStep struct {
	Amount    uint8
	FromStack uint8
	ToStack   uint8
}

func ParseProcedureStep(in string) ProcedureStep {
	// in looks like this "move 1 from 2 to 1"
	// we are only interested in the numbers
	splitInput := strings.Split(in, " ")
	return ProcedureStep{
		Amount:    cast.ToUint8(splitInput[1]),
		FromStack: cast.ToUint8(splitInput[3]),
		ToStack:   cast.ToUint8(splitInput[5]),
	}
}
