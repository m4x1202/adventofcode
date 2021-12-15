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
