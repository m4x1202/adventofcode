package math

import (
	"strings"

	"github.com/spf13/cast"
)

type WorryLevel []uint8

func NewWorryLevel(in string) WorryLevel {
	res := make([]uint8, len(in))
	for i := len(in) - 1; i >= 0; i-- {
		res[i] = uint8(in[i] - '0')
	}
	return res
}

func (w WorryLevel) Add(b uint) WorryLevel {
	stringB := cast.ToString(b)
	return w.AddWL(NewWorryLevel(stringB))
}

func (w WorryLevel) AddWL(b WorryLevel) WorryLevel {
	res := make([]uint8, len(w))
	copy(res, w)
	for i, elem := range b {
		if i >= len(res) {
			res = append(res, elem)
		} else {
			res[i] += elem
		}
	}
	uebertrag := true
	for uebertrag {
		uebertrag = false
		for i, elem := range res {
			if elem/10 == 0 {
				continue
			}
			uebertrag = true
			curr := elem / 10
			res[i] = elem % 10
			if i == len(res)-1 {
				res = append(res, curr)
			} else {
				res[i+1] += curr
			}
		}
	}
	return res
}

func (w WorryLevel) Mul(b uint) WorryLevel {
	res := w
	for i := uint(1); i < b; i++ {
		res = res.AddWL(res)
	}
	return res
}

func (w WorryLevel) Dividable(b uint) bool {
	var crosssum uint
	for _, elem := range w {
		crosssum += uint(elem)
	}
	return crosssum%b == 0
}

func (w WorryLevel) String() string {
	var res strings.Builder
	for i := len(w) - 1; i >= 0; i-- {
		res.WriteString(cast.ToString(w[i]))
	}
	return res.String()
}
