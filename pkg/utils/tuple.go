package utils

import (
	"github.com/m4x1202/adventofcode/pkg/physx"
	"github.com/spf13/cast"
)

type Tuple [2]interface{}

func (t Tuple) GetInt() [2]int {
	return [2]int{cast.ToInt(t[0]), cast.ToInt(t[1])}
}

func (t Tuple) GetVector() [2]physx.Vector {
	return [2]physx.Vector{t[0].(physx.Vector), t[1].(physx.Vector)}
}
