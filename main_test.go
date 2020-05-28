package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ValidateZeroValue(t *testing.T) {
	var (
		a1 = 100.0
		a2 = "xxxx"
	)
	mock := []struct {
		flag   string
		input  interface{}
		result interface{}
	}{
		{"string", "xxx", "xxx"},
		{"string", (*string)(nil), ""},
		{"float64", (*float64)(nil), 0.0},
		{"float64", (*float64)(&a1), 100.0},
		{"string", (*string)(&a2), "xxxx"},
	}

	for i, v := range mock {
		ref := ValidateZeroValue(v.input)
		r := GetReflectData(v.flag, ref)
		if r != v.result {
			fmt.Printf("ERROR CASE (%d), actual = (%v), expected = (%v)", i+1, r, v.result)
		}
	}
}

func GetReflectData(flag string, v reflect.Value) interface{} {
	if flag == "string" {
		return v.Interface().(string)
	}

	if flag == "float64" {
		return v.Interface().(float64)
	}

	return nil
}
