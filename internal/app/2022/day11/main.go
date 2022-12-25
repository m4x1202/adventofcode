package day11

import (
	"fmt"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "11"
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

func part1Func(preparedInput []Monkey) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for round := 0; round < 20; round++ {
		for i := range preparedInput {
			preparedInput[i].Turn(preparedInput)
		}
		roundLogger := partLogger.With().Int("round", round).Logger()
		for i, monkey := range preparedInput {
			roundLogger.Debug().Int("monkey", i).Msgf("%v", monkey.Items)
		}
	}
	for i, monkey := range preparedInput {
		partLogger.Debug().Int("monkey", i).Msgf("%v", monkey.InspectedItems)
	}
	inspectedItems := make([]int, len(preparedInput))
	for i, m := range preparedInput {
		inspectedItems[i] = int(m.InspectedItems)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectedItems)))
	monkeyBusiness := inspectedItems[0] * inspectedItems[1]
	puzzleAnswer = cast.ToUint64(monkeyBusiness)
	fmt.Printf("monkey business: %d\n", puzzleAnswer)
	return puzzleAnswer
}

var commonDivisor = uint(1)

func part2Func(preparedInput []Monkey) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64
	for _, monkey := range preparedInput {
		commonDivisor *= monkey.Test
	}
	partLogger.Info().Msgf("common divisor is %d", commonDivisor)

	for round := 0; round < 10000; round++ {
		for i := range preparedInput {
			preparedInput[i].Turn(preparedInput)
		}
		if round%1000 == 0 {
			roundLogger := partLogger.With().Int("round", round).Logger()
			for i, monkey := range preparedInput {
				roundLogger.Debug().Int("monkey", i).Msgf("%v", monkey.Items)
			}
		}
	}
	for i, monkey := range preparedInput {
		partLogger.Debug().Int("monkey", i).Msgf("%v", monkey.InspectedItems)
	}
	inspectedItems := make([]int, len(preparedInput))
	for i, m := range preparedInput {
		inspectedItems[i] = int(m.InspectedItems)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectedItems)))
	monkeyBusiness := inspectedItems[0] * inspectedItems[1]
	puzzleAnswer = cast.ToUint64(monkeyBusiness)
	fmt.Printf("monkey business: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []Monkey {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	splitInput := utils.SliceMap(utils.SplitSlice[[][]string](input, ""), func(in []string) []string {
		return utils.SliceMap(in, strings.TrimSpace)
	})
	converted := make([]Monkey, len(splitInput))
	for i := range splitInput {
		converted[i] = ParseMonkey(splitInput[i])
	}

	return converted
}

type Monkey struct {
	Name           uint8
	InspectedItems uint
	Items          []uint
	Operation      func(in uint) uint
	Test           uint
	Decision       [2]uint8
}

func ParseMonkey(in []string) Monkey {
	res := Monkey{
		Items:    []uint{},
		Decision: [2]uint8{},
	}
	nameString := in[0]
	startingItemsString := in[1]
	operationString := in[2]
	testString := in[3]
	decision1String := in[4]
	decision2String := in[5]
	res.Name = uint8(nameString[len(nameString)-2] - '0')
	res.Test = cast.ToUint(strings.TrimPrefix(testString, "Test: divisible by "))
	res.Decision[0] = cast.ToUint8(strings.TrimPrefix(decision1String, "If true: throw to monkey "))
	res.Decision[1] = cast.ToUint8(strings.TrimPrefix(decision2String, "If false: throw to monkey "))
	for _, item := range strings.Split(strings.TrimPrefix(startingItemsString, "Starting items: "), ", ") {
		res.Items = append(res.Items, cast.ToUint(item))
	}
	res.Operation = ParseMonkeyOperation(strings.TrimPrefix(operationString, "Operation: "))

	return res
}

func ParseMonkeyOperation(in string) func(in uint) uint {
	cutOperationString := strings.TrimPrefix(in, "new = old ")
	actualOperation, cutOperationString := rune(cutOperationString[0]), cutOperationString[2:]
	switch cutOperationString {
	case "old":
		switch actualOperation {
		case '+':
			return func(in uint) uint {
				return in + in
			}
		case '*':
			return func(in uint) uint {
				return in * in
			}
		}
	default:
		modifier := cast.ToUint(cutOperationString)
		switch actualOperation {
		case '+':
			return func(in uint) uint {
				return in + modifier
			}
		case '*':
			return func(in uint) uint {
				return in * modifier
			}
		}
	}
	panic("monkey operation could not be parsed")
}

func (m *Monkey) Turn(monkeyList []Monkey) {
	if len(m.Items) == 0 {
		return
	}
	monkeyLogger := partLogger.With().Uint8("monkey", m.Name).Logger()
	for len(m.Items) > 0 {
		m.InspectedItems++
		item := m.Items[0]
		m.Items = m.Items[1:]
		monkeyLogger.Trace().Msgf("Monkey inspects an item with a worry level of %d.", item)
		item = item % commonDivisor
		item = m.Operation(item)
		monkeyLogger.Trace().Msgf("Item worry level after operation %d.", item)
		//item = item.Relief()
		//monkeyLogger.Trace().Msgf("Monkey gets bored with item. Worry level is divided by 3 to %d.", item)
		if uint(item)%m.Test == 0 {
			monkeyLogger.Trace().Msgf("Current worry level is divisible by %d.", m.Test)
			monkeyList[m.Decision[0]].Items = append(monkeyList[m.Decision[0]].Items, item)
			monkeyLogger.Trace().Msgf("Item with worry level %d is thrown to monkey %d.", item, m.Decision[0])
		} else {
			monkeyLogger.Trace().Msgf("Current worry level is not divisible by %d.", m.Test)
			monkeyList[m.Decision[1]].Items = append(monkeyList[m.Decision[1]].Items, item)
			monkeyLogger.Trace().Msgf("Item with worry level %d is thrown to monkey %d.", item, m.Decision[1])
		}
	}
}
