package math

// import (
// 	"errors"
// 	"math"
// 	"strings"

// 	"github.com/spf13/cast"
// 	"golang.org/x/exp/constraints"
// )

// type Number interface {
// 	constraints.Integer | constraints.Float
// }

// type Vector[T Number] [3]T

// func FromString[T Number](in string) Vector[T] {
// 	v, _ := FromStringE[T](in)
// 	return v
// }

// func FromStringE[T Number](in string) (Vector[T], error) {
// 	vectorString := strings.Split(in, ",")
// 	if len(vectorString) != 3 {
// 		return Zero[T](), errors.New("cannot parse string with invalid length to vector")
// 	}
// 	res := Vector[T]{}
// 	for i, scalar := range vectorString {
// 		parsed, err := ToNumberE[T](scalar)
// 		if err != nil {
// 			return Zero[T](), err
// 		}
// 		res[i] = parsed
// 	}

// 	return res, nil
// }

// func getZero[T Number]() T {
// 	var result T
// 	return result
// }

// func ToNumberE[T Number](in any) (T, error) {
// 	var zero T
// 	switch any(zero).(type) {
// 	case uint8:
// 		out, err := cast.ToUint8E(in)
// 		return T(out), err
// 	case float64:
// 		out, err := cast.ToFloat64E(in)
// 		return T(out), err
// 	default:
// 		panic("")
// 	}
// }

// func Zero[T Number]() Vector[T] {
// 	return Vector[T]{0, 0, 0}
// }

// func (a Vector[T]) Copy() Vector[T] {
// 	res := a
// 	return res
// }

// func (a Vector[T]) Add(v Vector[T]) Vector[T] {
// 	if len(a) != len(v) {
// 		return a
// 	}
// 	for i := range a {
// 		a[i] += v[i]
// 	}
// 	return a
// }

// func (a Vector[T]) Mul(m T) Vector[T] {
// 	res := Vector[T]{}
// 	for i := range a {
// 		res[i] = a[i] * m
// 	}
// 	return res
// }

// func (a Vector[T]) Sub(v Vector[T]) Vector[T] {
// 	if len(a) != len(v) {
// 		return a
// 	}
// 	for i := range a {
// 		a[i] -= v[i]
// 	}
// 	return a
// }

// func (a Vector[T]) Magnitude() T {
// 	var res T
// 	for _, scalar := range a {
// 		res += scalar * scalar
// 	}
// 	return math.Sqrt(res)
// }

// func (a Vector[T]) Normalized() Vector[T] {
// 	res := Vector[T]{}
// 	mag := a.Magnitude()

// 	for i, scalar := range a {
// 		res[i] = scalar / mag
// 	}
// 	return res
// }
