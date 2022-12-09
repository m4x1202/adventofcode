package utils

import (
	"fmt"
	"reflect"
	"strings"
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

type MapElem[T comparable] struct {
	X    uint
	Y    uint
	Data T
}

type SingleSliceMap[T comparable] struct {
	Elems    []*MapElem[T]
	Metadata any
}

func (m SingleSliceMap[T]) String() string {
	var res strings.Builder
	for i := uint(0); i < m.GetHeight(); i++ {
		row := m.GetRow(i)
		toString := make([]string, 0, len(row))
		for _, elem := range row {
			if elem == nil {
				toString = append(toString, ".")
			} else {
				elemValue := reflect.ValueOf(elem.Data)
				if elemValue.Kind() == reflect.Ptr {
					toString = append(toString, fmt.Sprintf("%v", elemValue.Elem()))
				} else {
					if elem.Data == getZero[T]() {
						toString = append(toString, "#")
					} else {
						toString = append(toString, fmt.Sprintf("%v", elem.Data))
					}
				}
			}
		}
		res.WriteString(fmt.Sprintf("%s\n", strings.Join(toString, " ")))
	}
	return res.String()
}

func (m *SingleSliceMap[T]) ModifyElem(mod func(elem T) T, x, y uint) {
	elem := m.GetElem(x, y)
	if elem != nil {
		elem.Data = mod(elem.Data)
		return
	}
	m.Elems = append(m.Elems, &MapElem[T]{x, y, mod(getZero[T]())})
}

func (m *SingleSliceMap[T]) RemoveElem(x, y uint) {
	for i := 0; i < len(m.Elems); i++ {
		if x == m.Elems[i].X && y == m.Elems[i].Y {
			m.Elems = append(m.Elems[:i], m.Elems[i+1:]...)
			return
		}
	}
}

func (m *SingleSliceMap[T]) GetElem(x, y uint) *MapElem[T] {
	for _, elem := range m.Elems {
		if x == elem.X && y == elem.Y {
			return elem
		}
	}
	return nil
}

func (m *SingleSliceMap[T]) GetRow(index uint) []*MapElem[T] {
	res := make([]*MapElem[T], m.GetWidth())
	for _, elem := range m.Elems {
		if elem.Y == index {
			res[elem.X] = elem
		}
	}
	return res
}

func (m *SingleSliceMap[T]) GetCol(index uint) []*MapElem[T] {
	res := make([]*MapElem[T], m.GetHeight())
	for _, elem := range m.Elems {
		if elem.X == index {
			res[elem.Y] = elem
		}
	}
	return res
}

func (m SingleSliceMap[T]) GetWidth() uint {
	var width uint
	for _, elem := range m.Elems {
		if elem.X > width {
			width = elem.X
		}
	}
	return width
}

func (m SingleSliceMap[T]) GetHeight() uint {
	var height uint
	for _, elem := range m.Elems {
		if elem.Y > height {
			height = elem.Y
		}
	}
	return height
}

func getZero[T any]() T {
	var result T
	return result
}
