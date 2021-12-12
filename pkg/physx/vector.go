package physx

import (
	"math"
	"strconv"
	"strings"
)

type Vector []float64

func FromString(in string) (*Vector, error) {
	vectorString := strings.Split(in, ",")
	res := make(Vector, 0, len(vectorString))
	for _, scalar := range vectorString {
		parsed, err := strconv.ParseFloat(scalar, 64)
		if err != nil {
			return &Zero, err
		}
		res = append(res, parsed)
	}

	return &res, nil
}

var (
	Zero = Vector{0.0, 0.0, 0.0}
)

func (a *Vector) Copy() *Vector {
	res := make(Vector, len(*a))
	copy(res, *a)
	return &res
}

func (a *Vector) Ceil() *Vector {
	res := make(Vector, 0, len(*a))
	for i := 0; i < len(*a); i++ {
		res = append(res, math.Copysign(math.Ceil(math.Abs((*a)[i])), (*a)[i]))
	}
	return &res
}

func (a *Vector) Add(v Vector) *Vector {
	if len(*a) != len(v) {
		return a
	}
	for i := 0; i < len(*a); i++ {
		(*a)[i] += v[i]
	}
	return a
}

func (a *Vector) Sub(v Vector) *Vector {
	if len(*a) != len(v) {
		return a
	}
	for i := 0; i < len(*a); i++ {
		(*a)[i] -= v[i]
	}
	return a
}

func (a *Vector) Magnitude() float64 {
	var res float64
	for _, scalar := range *a {
		res += math.Pow(scalar, 2)
	}
	return math.Sqrt(res)
}

func (a *Vector) Normalized() *Vector {
	res := make(Vector, 0, len(*a))
	mag := a.Magnitude()

	for _, scalar := range *a {
		res = append(res, scalar/mag)
	}
	return &res
}
