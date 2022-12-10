package utils

import (
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

/* ==========================================
 * ============ CoordinateSystem ==============
 * ==========================================
 */

type CoordinateSystem[N constraints.Integer, T any] map[N]map[N]*T

// String function satisfies fmt.Stringer interface
func (m CoordinateSystem[N, T]) String() string {
	var res strings.Builder
	height, negPortion := m.GetHeight()
	for y := N(height) - negPortion; y > -negPortion; y-- {
		row := m.GetRow(y - 1)
		toString := make([]string, 0, len(row))
		for _, elem := range row {
			if elem == nil {
				toString = append(toString, ".")
			} else {
				toString = append(toString, "#")
			}
		}
		res.WriteString(fmt.Sprintf("%s\n", strings.Join(toString, "")))
	}
	return res.String()
}

// ModifyElem modifies the element at the given coordinates using a given function
// If unset passes 'nil' to the function -> expected behavior is to set new elements this way
func (m CoordinateSystem[N, T]) ModifyElemFunc(mod func(elem *T) *T, x, y N) {
	if _, exists := m[x]; !exists {
		m[x] = map[N]*T{}
	}
	m[x][y] = mod(m[x][y])
}

// GetRow returns the row of elements at the given index
// Unset elements will default to zero value MapElem
func (m CoordinateSystem[N, T]) GetRow(index N) []*T {
	width, negPortion := m.GetWidth()
	res := make([]*T, width)
	for x, col := range m {
		if elem, exists := col[index]; exists {
			res[x+negPortion] = elem
		}
	}
	return res
}

// GetCol returns the column of elements at the given index
// Unset elements will default to zero value MapElem
func (m CoordinateSystem[N, T]) GetCol(index N) []*T {
	height, negPortion := m.GetHeight()
	res := make([]*T, height)
	if col, exists := m[index]; exists {
		for y, elem := range col {
			res[y+negPortion] = elem
		}
	}
	return res
}

// GetWidth returns the total width of the SingleSliceMap and an indicator of how big the negative portion of the height (as a positive value)
func (m CoordinateSystem[N, T]) GetWidth() (uint, N) {
	var negPortion N
	var width uint
	for x := range m {
		if x < negPortion {
			negPortion = x
		} else if x > 0 && uint(x) > width {
			width = uint(x)
		}
	}
	return 1 + width + uint(-negPortion), -negPortion
}

// GetHeight returns the total height of the SingleSliceMap and an indicator of how big the negative portion of the height (as a positive value)
func (m CoordinateSystem[N, T]) GetHeight() (uint, N) {
	var negPortion N
	var height uint
	for _, col := range m {
		for y := range col {
			if y < negPortion {
				negPortion = y
			} else if y > 0 && uint(y) > height {
				height = uint(y)
			}
		}
	}
	return 1 + height + uint(-negPortion), -negPortion
}

func (m CoordinateSystem[N, T]) TotalSize() int {
	var totalSize int
	if len(m) == 0 {
		return 0
	}
	for _, col := range m {
		totalSize += len(col)
	}
	return totalSize
}

/* ==========================================
 * ============ SingleSliceMap ==============
 * ==========================================
 */

type MapElem[N constraints.Integer, T any] struct {
	X, Y N
	Data *T
}

// SingleSliceMap is a map with X, Y coordinates that only stores elements that are actually set in a single slice
// Unset elements do thus not consume any memory
type SingleSliceMap[N constraints.Integer, T any] []MapElem[N, T]

// String function satisfies fmt.Stringer interface
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

// ModifyElem modifies the element at the given coordinates using a given function
// If unset passes 'nil' to the function -> expected behavior is to set new elements this way
func (m *SingleSliceMap[N, T]) ModifyElemFunc(mod func(elem *T) *T, x, y N) {
	elem := m.GetElem(x, y)
	if elem != nil {
		*elem = *mod(elem)
		return
	}
	*m = append(*m, MapElem[N, T]{x, y, mod(nil)})
}

// RemoveElem removes the element at the given coordinates
// If unset the function is no-op
func (m *SingleSliceMap[N, T]) RemoveElem(x, y N) {
	for i := range *m {
		if x == (*m)[i].X && y == (*m)[i].Y {
			*m = append((*m)[:i], (*m)[i+1:]...)
			return
		}
	}
}

// GetElem returns the element at the given coordinates or nil if unset
func (m SingleSliceMap[N, T]) GetElem(x, y N) *T {
	for _, elem := range m {
		if x == elem.X && y == elem.Y {
			return elem.Data
		}
	}
	return nil
}

// GetRow returns the row of elements at the given index
// Unset elements will default to zero value MapElem
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

// GetCol returns the column of elements at the given index
// Unset elements will default to zero value MapElem
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

// GetWidth returns the total width of the SingleSliceMap and an indicator of how big the negative portion of the height (as a positive value)
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

// GetHeight returns the total height of the SingleSliceMap and an indicator of how big the negative portion of the height (as a positive value)
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
