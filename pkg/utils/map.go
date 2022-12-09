package utils

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Map[T comparable] [][]T

func NewMap[T comparable](width, height int) (hm Map[T]) {
	hm = make(Map[T], width)
	for i := range hm {
		hm[i] = make([]T, height)
	}
	return
}

func (w Map[T]) Tile(x, y int) (T, bool) {
	if 0 <= x && x <= len(w)-1 && 0 <= y && y <= len(w[0])-1 {
		return w[x][y], true
	}
	return getZero[T](), false
}

var (
	ErrDoesNotExist = errors.New("element does not exist")
)

type MapElem[N constraints.Integer, T any] struct {
	X, Y N
	Data *T
}

type SingleSliceMap[N constraints.Integer, T any] []MapElem[N, T]

func (m SingleSliceMap[N, T]) String() string {
	var res strings.Builder
	height, negPortion := m.GetHeight()
	for i := N(height) - negPortion - 1; i >= -negPortion; i-- {
		row := m.GetRow(i)
		toString := make([]string, 0, len(row))
		for _, elem := range row {
			if elem.Data == nil {
				toString = append(toString, ".")
			} else {
				toString = append(toString, "#")
			}
		}
		res.WriteString(fmt.Sprintf("%s\n", strings.Join(toString, "")))
	}
	return res.String()
}

func (m *SingleSliceMap[N, T]) ModifyElem(mod func(elem *T) *T, x, y N) {
	elem, err := m.GetElem(x, y)
	if err == nil {
		elem.Data = mod(elem.Data)
		return
	}
	*m = append(*m, MapElem[N, T]{x, y, mod(nil)})
}

func (m *SingleSliceMap[N, T]) RemoveElem(x, y N) {
	for i := range *m {
		if x == (*m)[i].X && y == (*m)[i].Y {
			*m = append((*m)[:i], (*m)[i+1:]...)
			return
		}
	}
}

func (m SingleSliceMap[N, T]) GetElem(x, y N) (MapElem[N, T], error) {
	for _, elem := range m {
		if x == elem.X && y == elem.Y {
			return elem, nil
		}
	}
	return MapElem[N, T]{}, ErrDoesNotExist
}

func (m SingleSliceMap[N, T]) GetRow(index N) []MapElem[N, T] {
	width, negPortion := m.GetWidth()
	res := make([]MapElem[N, T], width)
	for _, elem := range m {
		if elem.Y == index {
			res[elem.X+negPortion] = elem
		}
	}
	return res
}

func (m SingleSliceMap[N, T]) GetCol(index N) []MapElem[N, T] {
	height, negPortion := m.GetHeight()
	res := make([]MapElem[N, T], height)
	for _, elem := range m {
		if elem.X == index {
			res[elem.Y+negPortion] = elem
		}
	}
	return res
}

func (m SingleSliceMap[N, T]) GetWidth() (uint, N) {
	var negPortion N
	var width uint
	for _, elem := range m {
		if elem.X < negPortion {
			negPortion = elem.X
		} else if elem.X > 0 && uint(elem.X) > width {
			width = uint(elem.X)
		}
	}
	return 1 + width + uint(-negPortion), -negPortion
}

func (m SingleSliceMap[N, T]) GetHeight() (uint, N) {
	var negPortion N
	var height uint
	for _, elem := range m {
		if elem.Y < negPortion {
			negPortion = elem.Y
		} else if elem.Y > 0 && uint(elem.Y) > height {
			height = uint(elem.Y)
		}
	}
	return 1 + height + uint(-negPortion), -negPortion
}
