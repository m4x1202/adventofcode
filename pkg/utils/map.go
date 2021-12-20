package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type MapElem struct {
	X    uint
	Y    uint
	Data interface{}
}
type Map struct {
	Width    uint
	Height   uint
	Elems    []*MapElem
	Metadata interface{}
}

func (m Map) String() string {
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
					if elem.Data == nil {
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

func (m *Map) ModifyElem(mod func(elem interface{}) interface{}, x, y uint) {
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
	m.Elems = append(m.Elems, &MapElem{x, y, mod(nil)})
}

func (m *Map) RemoveElem(x, y uint) {
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

func (m *Map) GetElem(x, y uint) *MapElem {
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

func (m *Map) GetRow(index uint) []*MapElem {
	if index > m.Height {
		return nil
	}
	res := make([]*MapElem, m.Width)
	for _, elem := range m.Elems {
		if elem.Y == index {
			res[elem.X] = elem
		}
	}
	return res
}

func (m *Map) GetCol(index uint) []*MapElem {
	if index > m.Width {
		return nil
	}
	res := make([]*MapElem, m.Height)
	for _, elem := range m.Elems {
		if elem.X == index {
			res[elem.Y] = elem
		}
	}
	return res
}
