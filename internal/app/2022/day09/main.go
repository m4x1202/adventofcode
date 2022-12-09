package day09

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/physx"
	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "09"
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

func part1Func(preparedInput []physx.Vector) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	rope := make(Rope, 2)
	for i := range rope {
		rope[i] = physx.Zero(2)
	}
	rope.UpdateTrail()
	for _, step := range preparedInput {
		rope.Translate(step)
		partLogger.Info().Msgf("%v", rope)
	}

	puzzleAnswer = cast.ToUint64(len(Trail))
	fmt.Printf("total pos visited: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(preparedInput []physx.Vector) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	rope := make(Rope, 10)
	for i := range rope {
		rope[i] = physx.Zero(2)
	}
	rope.UpdateTrail()
	for _, step := range preparedInput {
		rope.Translate(step)
		partLogger.Info().Msgf("%v", rope)
	}

	fmt.Print(Trail.String())

	puzzleAnswer = cast.ToUint64(len(Trail))
	fmt.Printf("total pos visited: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []physx.Vector {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]physx.Vector, len(input))
	for i, mov := range input {
		converted[i] = ParseStep(mov)
	}

	return converted
}

func ParseStep(in string) physx.Vector {
	splitStep := strings.Split(in, " ")
	var res physx.Vector
	switch splitStep[0] {
	case "U":
		res = physx.Up
	case "D":
		res = physx.Down
	case "L":
		res = physx.Left
	case "R":
		res = physx.Right
	}
	res = res.Mul(cast.ToFloat64(splitStep[1]))
	return res
}

var (
	Trail = utils.SingleSliceMap[int, bool]{}
)

type Rope []physx.Vector

func (r Rope) Translate(vec physx.Vector) {
	magMoveVec := vec.Magnitude()
	normMoveVec := vec.Normalized().Ceil()
	for step := 0; step < int(magMoveVec); step++ {
		r[0] = r[0].Add(normMoveVec)
		for i := 1; i < len(r); i++ {
			distanceVec := r[i-1].Sub(r[i])
			if reflect.DeepEqual(distanceVec, physx.Zero(2)) {
				continue
			}
			for distanceVec.Magnitude() >= 2 {
				tailMoveVec := distanceVec.Normalized().Ceil()
				r[i] = r[i].Add(tailMoveVec)
				distanceVec = r[i-1].Sub(r[i])
				if i == len(r)-1 {
					r.UpdateTrail()
				}
			}
		}
	}
}

func (r Rope) UpdateTrail() {
	Trail.ModifyElem(func(elem *bool) *bool {
		if elem == nil {
			var res bool
			return &res
		}
		*elem = true
		return elem
	}, int(r[len(r)-1][0]), int(r[len(r)-1][1]))
}
