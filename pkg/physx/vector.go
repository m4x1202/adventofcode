package physx

import (
	"math"
	"strconv"
	"strings"
)

type Vector []float64

func ToVector(in string) Vector {
	v, _ := ToVectorE(in)
	return v
}

func ToVectorE(in string) (Vector, error) {
	vectorString := strings.Split(in, ",")
	res := make(Vector, len(vectorString))
	for i, scalar := range vectorString {
		parsed, err := strconv.ParseFloat(scalar, 64)
		if err != nil {
			return Zero(len(vectorString)), err
		}
		res[i] = parsed
	}

	return res, nil
}

func Zero(len int) Vector {
	res := make(Vector, len)
	for i := range res {
		res[i] = 0
	}
	return res
}

var (
	Up    = Vector{0, 1}
	Down  = Vector{0, -1}
	Left  = Vector{-1, 0}
	Right = Vector{1, 0}
)

func (a Vector) Copy() Vector {
	res := a
	return res
}

func (a Vector) Ceil() Vector {
	res := make(Vector, len(a))
	for i := range a {
		res[i] = math.Copysign(math.Ceil(math.Abs(a[i])), a[i])
	}
	return res
}

func (a Vector) Add(v Vector) Vector {
	if len(a) != len(v) {
		return a
	}
	res := make(Vector, len(a))
	for i := range a {
		res[i] = a[i] + v[i]
	}
	return res
}

func (a Vector) Sub(v Vector) Vector {
	if len(a) != len(v) {
		return a
	}
	res := make(Vector, len(a))
	for i := range a {
		res[i] = a[i] - v[i]
	}
	return res
}

func (a Vector) Mul(b float64) Vector {
	res := make(Vector, len(a))
	for i := range a {
		res[i] = a[i] * b
	}
	return res
}

func (a Vector) Magnitude() float64 {
	var res float64
	for _, scalar := range a {
		res += scalar * scalar
	}
	return math.Sqrt(res)
}

func (a Vector) Normalized() Vector {
	res := make(Vector, len(a))
	mag := a.Magnitude()

	for i, scalar := range a {
		res[i] = scalar / mag
	}
	return res
}
