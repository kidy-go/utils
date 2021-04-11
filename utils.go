// utils.go kee > 2021/03/16

package utils

import (
	"net"
	"reflect"
)

// if the val is empty (IsEmpty) then ...
func HasOr(has interface{}, then ...interface{}) interface{} {
	var (
		t = has
		f interface{}
	)
	if len(then) > 1 {
		t, f = then[0], then[1]
	} else if len(then) > 0 {
		f = then[0]
	}

	if !IsEmpty(has) {
		return t
	}
	return f
}

var ipTyp = reflect.TypeOf(net.IP{})

// Returns true if has nil or false or 0 or ""
func IsEmpty(val interface{}) bool {
	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		// If Net.IP
		if v.Type() == ipTyp {
			return len(v.Interface().(net.IP)) == 0
		}
		return v.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		if v.CanInterface() {
			return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
		}
		return false
	case reflect.String:
		return v.Len() <= 0
	case reflect.Bool:
		return !val.(bool)
	default:
		return false
	}
}

func InArray(needle interface{}, haystack interface{}) (index int, exists bool) {
	index, exists = -1, false

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(haystack)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return index, exists
}
