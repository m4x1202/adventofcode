package physx

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Vector [3]float64

func FromString(in string) Vector {
	v, _ := FromStringE(in)
	return v
}

func FromStringE(in string) (Vector, error) {
	vectorString := strings.Split(in, ",")
	if len(vectorString) != 3 {
		return Zero, errors.New("cannot parse string with invalid length to vector")
	}
	res := Vector{}
	for i, scalar := range vectorString {
		parsed, err := strconv.ParseFloat(scalar, 64)
		if err != nil {
			return Zero, err
		}
		res[i] = parsed
	}

	return res, nil
}

var (
	Zero  = Vector{0.0, 0.0, 0.0}
	Up    = Vector{0.0, 1.0, 0.0}
	Down  = Vector{0.0, -1.0, 0.0}
	Left  = Vector{-1.0, 0.0, 0.0}
	Right = Vector{1.0, 1.0, 0.0}
)

func (a Vector) Copy() Vector {
	res := a
	return res
}

func (a Vector) Ceil() Vector {
	res := Vector{}
	for i := range a {
		res[i] = math.Copysign(math.Ceil(math.Abs(a[i])), a[i])
	}
	return res
}

func (a Vector) Add(v Vector) Vector {
	if len(a) != len(v) {
		return a
	}
	for i := range a {
		a[i] += v[i]
	}
	return a
}

func (a Vector) Sub(v Vector) Vector {
	if len(a) != len(v) {
		return a
	}
	for i := range a {
		a[i] -= v[i]
	}
	return a
}

func (a Vector) Magnitude() float64 {
	var res float64
	for _, scalar := range a {
		res += math.Pow(scalar, 2)
	}
	return math.Sqrt(res)
}

func (a Vector) Normalized() Vector {
	res := Vector{}
	mag := a.Magnitude()

	for i, scalar := range a {
		res[i] = scalar / mag
	}
	return res
}
