package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := (*float64)(nil)
	v := ValidateZeroValue(x)
	fmt.Printf("%v", v.Interface().(float64))
}

func ValidateZeroValue(a interface{}) reflect.Value {
	valoo := reflect.ValueOf(a)
	if valoo.Kind() == reflect.Ptr && valoo.IsNil() {
		elem := reflect.TypeOf(a).Elem()
		return reflect.Zero(elem)
	}

	if valoo.Kind() != reflect.Ptr && valoo.IsValid() {
		return valoo
	}

	return valoo.Elem()
}
