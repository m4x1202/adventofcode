package day13

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "13"
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

func part1Func(preparedInput []PacketPair) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for i, pp := range preparedInput {
		if pp.CheckRightOrder() {
			puzzleAnswer += uint64(i) + 1
		}
	}

	fmt.Printf("sum of indices of packet pairs in right order %d\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(preparedInput []PacketPair) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	packetList := make([]Packet, len(preparedInput)*2)
	for i, elem := range preparedInput {
		packetList[i*2] = elem.Left
		packetList[i*2+1] = elem.Right
	}
	dividerPacket1 := Packet{Packet{uint8(2)}}
	packetList = append(packetList, dividerPacket1)
	dividerPacket2 := Packet{Packet{uint8(6)}}
	packetList = append(packetList, dividerPacket2)
	sort.Slice(packetList, func(i, j int) bool {
		outcome, _ := packetList[i].CompareWith(packetList[j])
		if outcome {
			return true
		} else {
			return false
		}
	})
	var dividerPacket1Index, dividerPacket2Index int
	for i, elem := range packetList {
		if reflect.DeepEqual(elem, dividerPacket1) {
			dividerPacket1Index = i + 1
		} else if reflect.DeepEqual(elem, dividerPacket2) {
			dividerPacket2Index = i + 1
		}
	}

	puzzleAnswer = cast.ToUint64(dividerPacket1Index * dividerPacket2Index)
	fmt.Printf("mult of indices of divider packets %d\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []PacketPair {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	splitInput := utils.SplitSlice[[][]string](input, "")
	converted := make([]PacketPair, len(splitInput))
	for i, elem := range splitInput {
		converted[i] = ParsePacketPair(elem)
	}

	return converted
}

type PacketPair struct {
	Left  Packet
	Right Packet
}

func (pp PacketPair) CheckRightOrder() bool {
	outcome, err := pp.Left.CompareWith(pp.Right)
	if err != nil {
		panic("packets are equal")
	}
	return outcome
}

func ParsePacketPair(in []string) PacketPair {
	return PacketPair{
		Left:  ParsePacket(in[0]),
		Right: ParsePacket(in[1]),
	}
}

type Packet []any

func (p1 Packet) CompareWith(p2 Packet) (bool, error) {
	if len(p1) == 0 && 0 < len(p2) {
		return true, nil
	}
	for i, p1Elem := range p1 {
		if len(p2) <= i {
			return false, nil
		}
		switch e := p1Elem.(type) {
		case Packet:
			if p2Elem, ok := p2[i].(Packet); ok {
				if outcome, err := e.CompareWith(p2Elem); err == nil {
					return outcome, nil
				}
			} else if p2Elem, ok := p2[i].(uint8); ok {
				if outcome, err := e.CompareWith(Packet{p2Elem}); err == nil {
					return outcome, nil
				}
			}
		case uint8:
			if p2Elem, ok := p2[i].(Packet); ok {
				if outcome, err := (Packet{e}).CompareWith(p2Elem); err == nil {
					return outcome, nil
				}
			} else if p2Elem, ok := p2[i].(uint8); ok {
				if e != p2Elem {
					return e < p2Elem, nil
				}
			}
		}
	}
	if len(p1) < len(p2) {
		return true, nil
	}
	return false, errors.New("packets are equal")
}

func ParsePacket(in string) Packet {
	in = strings.TrimPrefix(in, "[")
	in = strings.TrimSuffix(in, "]")
	res := Packet{}
	if len(in) == 0 {
		return res
	}
	for len(in) > 0 {
		if strings.HasPrefix(in, "[") {
			indexClosingBracket := IndexOfMatchingClosingBracket(in)
			res = append(res, ParsePacket(in[:indexClosingBracket+1]))
			in = in[indexClosingBracket+1:]
		}
		newElem, newIn, exists := strings.Cut(in, ",")
		if len(newElem) > 0 {
			res = append(res, cast.ToUint8(newElem))
		}
		if !exists {
			return res
		}
		in = newIn
	}
	panic("somehow packet did not contain a final element")
}

func IndexOfMatchingClosingBracket(in string) int {
	openingBrackets := 0
	for i, char := range in {
		if char == '[' {
			openingBrackets++
		} else if char == ']' {
			openingBrackets--
		}
		if openingBrackets == 0 {
			return i
		}
	}
	return -1
}
