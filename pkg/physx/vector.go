package physx

import (
	"math"
	"strconv"
	"strings"
)

type Scalar float64

func (a *Scalar) Add(b Scalar) {
	*a += b
}

func (a *Scalar) Sub(b Scalar) {
	*a -= b
}

type Vector []Scalar

func FromString(in string) (Vector, error) {
	vectorString := strings.Split(in, ",")
	res := make(Vector, 0, len(vectorString))
	for _, scalar := range vectorString {
		parsed, err := strconv.ParseFloat(scalar, 64)
		if err != nil {
			return Zero, err
		}
		res = append(res, Scalar(parsed))
	}

	return res, nil
}

var (
	Zero    = Vector{0.0, 0.0, 0.0}
	Gravity = Vector{0.0, -9.81, 0.0}
)

func (a Vector) Copy() (res Vector) {
	copy(res, a)
	return
}

func (a *Vector) Add(v Vector) {
	if len(*a) != len(v) {
		return
	}
	for i := 0; i < len(*a); i++ {
		(*a)[i].Add(v[i])
	}
}

func (a *Vector) Sub(v Vector) {
	if len(*a) != len(v) {
		return
	}
	for i := 0; i < len(*a); i++ {
		(*a)[i].Sub(v[i])
	}
}

func (a Vector) Magnitude() float64 {
	var res float64
	for _, scalar := range a {
		res += math.Pow(float64(scalar), 2)
	}
	return math.Sqrt(res)
}

func (a Vector) Normalized() Vector {
	res := make(Vector, 0, len(a))
	mag := a.Magnitude()

	for _, scalar := range a {
		res = append(res, scalar/Scalar(mag))
	}
	return res
}
