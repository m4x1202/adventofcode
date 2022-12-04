package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type MapElem[T comparable] struct {
	X    uint
	Y    uint
	Data T
}
type Map[T comparable] struct {
	Width    uint
	Height   uint
	Elems    []*MapElem[T]
	Metadata any
}

func (m Map[T]) String() string {
	var res strings.Builder
	for i := uint(0); i < m.Height; i++ {
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

func (m *Map[T]) ModifyElem(mod func(elem T) T, x, y uint) {
	elem := m.GetElem(x, y)
	if elem != nil {
		elem.Data = mod(elem.Data)
		return
	}
	if x >= m.Width {
		m.Width = x + 1
	}
	if y >= m.Height {
		m.Height = y + 1
	}
	m.Elems = append(m.Elems, &MapElem[T]{x, y, mod(getZero[T]())})
}

func (m *Map[T]) RemoveElem(x, y uint) {
	if x > m.Width || y > m.Height {
		return
	}
	for i := 0; i < len(m.Elems); i++ {
		if x == m.Elems[i].X && y == m.Elems[i].Y {
			m.Elems = append(m.Elems[:i], m.Elems[i+1:]...)
			return
		}
	}
}

func (m *Map[T]) GetElem(x, y uint) *MapElem[T] {
	if x > m.Width || y > m.Height {
		return nil
	}
	for _, elem := range m.Elems {
		if x == elem.X && y == elem.Y {
			return elem
		}
	}
	return nil
}

func (m *Map[T]) GetRow(index uint) []*MapElem[T] {
	if index > m.Height {
		return nil
	}
	res := make([]*MapElem[T], m.Width)
	for _, elem := range m.Elems {
		if elem.Y == index {
			res[elem.X] = elem
		}
	}
	return res
}

func (m *Map[T]) GetCol(index uint) []*MapElem[T] {
	if index > m.Width {
		return nil
	}
	res := make([]*MapElem[T], m.Height)
	for _, elem := range m.Elems {
		if elem.X == index {
			res[elem.Y] = elem
		}
	}
	return res
}

func getZero[T any]() T {
	var result T
	return result
}
