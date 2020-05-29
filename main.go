package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := (*float64)(nil)
	fmt.Printf("%v", Get(x).(float64))
}

func Get(a interface{}) interface{} {
	valoo := reflect.ValueOf(a)
	if valoo.Kind() == reflect.Ptr && valoo.IsNil() {
		elem := reflect.TypeOf(a).Elem()
		return reflect.Zero(elem).Interface()
	}

	if valoo.Kind() != reflect.Ptr && valoo.IsValid() {
		return valoo.Interface()
	}

	return valoo.Elem().Interface()
}
