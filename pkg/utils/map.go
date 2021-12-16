package utils

import (
	"fmt"
	"strings"
)

type MapElem struct {
	x    uint
	y    uint
	Data interface{}
}
type Map struct {
	Width    uint
	Height   uint
	elems    []*MapElem
	Metadata interface{}
}

func (m *Map) String() string {
	var res strings.Builder
	for _, elem := range m.elems {
		res.WriteString(fmt.Sprintf("%v", elem.Data))
	}
	return res.String()
}

func (m *Map) ModifyElem(mod func(elem interface{}) interface{}, x, y uint) {
	elem := m.GetElem(x, y)
	if elem != nil {
		elem.Data = mod(elem.Data)
		return
	}
	if x > m.Width {
		m.Width = x
	}
	if y > m.Height {
		m.Height = y
	}
	m.elems = append(m.elems, &MapElem{x, y, mod(nil)})
}

func (m *Map) GetElem(x, y uint) *MapElem {
	if x > m.Width || y > m.Height {
		return nil
	}
	for _, elem := range m.elems {
		if x == elem.x && y == elem.y {
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
	for _, elem := range m.elems {
		if elem.y == index {
			res[elem.x-1] = elem
		}
	}
	return res
}

func (m *Map) GetCol(index uint) []*MapElem {
	if index > m.Width {
		return nil
	}
	res := make([]*MapElem, m.Height)
	for _, elem := range m.elems {
		if elem.x == index {
			res[elem.y-1] = elem
		}
	}
	return res
}
