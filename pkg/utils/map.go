package utils

type MapElem struct {
	x    int
	y    int
	data *interface{}
}
type Map struct {
	Width    int
	Height   int
	elems    []MapElem
	Metadata interface{}
}

func (m *Map) ModifyElem(mod func(elem interface{}) interface{}, x, y int) {
	for i := 0; i < len(m.elems); i++ {
		if m.elems[i].y == y && m.elems[i].x == x {
			newElem := mod(*m.elems[i].data)
			m.elems[i].data = &newElem
			return
		}
	}
	if x > m.Width {
		m.Width = x
	}
	if y > m.Height {
		m.Height = y
	}
	newElem := mod(nil)
	m.elems = append(m.elems, MapElem{x, y, &newElem})
}

func (m *Map) GetRow(index int) []*interface{} {
	if index > m.Height {
		return nil
	}
	res := make([]*interface{}, m.Width)
	for _, elem := range m.elems {
		if elem.y == index {
			res[elem.x-1] = elem.data
		}
	}
	return res
}

func (m *Map) GetCol(index int) []*interface{} {
	if index > m.Width {
		return nil
	}
	res := make([]*interface{}, m.Height)
	for _, elem := range m.elems {
		if elem.x == index {
			res[elem.y-1] = elem.data
		}
	}
	return res
}
