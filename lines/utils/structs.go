package utils

import (
	"fmt"
	"reflect"
)

func GetField(obj interface{}, name string) (interface{}, bool) {
	r := reflect.ValueOf(obj)
	field := reflect.Indirect(r).FieldByName(name)

	// Make sure field is valid and can be interfaced
	if field.IsValid() && field.CanInterface() {
		return field.Interface(), true
	}

	// Field does not exist or cannot be interfaced
	fmt.Printf("utils.GetField: Invalid field: %s\n", name)
	return "", false
}
