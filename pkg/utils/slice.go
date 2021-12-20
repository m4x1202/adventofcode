package utils

import (
	"reflect"
)

func Contains(list interface{}, elem interface{}) bool {
	listValue := reflect.ValueOf(list)
	if listValue.Kind() == reflect.Slice || listValue.Kind() == reflect.Array {
		for i := 0; i < listValue.Len(); i++ {
			if reflect.DeepEqual(listValue.Index(i).Interface(), elem) {
				return true
			}
		}
	}
	return false
}

func IndexOfEmpty(in []string) int {
	for index, elem := range in {
		if elem == "" {
			return index
		}
	}
	return -1
}

func SlidingWindowString(size int, input string) []string {
	// returns the input slice as the first element
	if len(input) <= size {
		return []string{input}
	}

	// allocate slice at the precise size we need
	r := make([]string, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
