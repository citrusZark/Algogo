// sort
package sort

import (
	"fmt"
	"reflect"
)

type Sort interface {
	Sort(s []interface{}, fieldName string)
	less(fi interface{}, fj interface{}, fieldName string) bool
	swap(fi interface{}, fj interface{})
	IsSorted(s []interface{}) bool
	Show(s []interface{})
}

type sortbase struct{}

//Sort method only take slice
func (ss sortbase) Sort(s interface{}, fieldName string) {}
func (ss sortbase) less(fi reflect.Value, fj reflect.Value, fieldName string) bool {
	var field string
	switch fi.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fi.Int() < fj.Int()
	case reflect.String:
		return fi.String() < fj.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fi.Uint() < fj.Uint()
	case reflect.Float32, reflect.Float64:
		return fi.Float() < fj.Float()
	case reflect.Struct:
		field = fieldName
		return ss.less(fi.FieldByName(field), fj.FieldByName(field), field)
	default:
		panic("Unsupported data or struct's field not found")
	}
}
func (_ sortbase) swap(i, j int, s reflect.Value) {
	x, y := s.Index(i).Interface(), s.Index(j).Interface()
	s.Index(i).Set(reflect.ValueOf(y))
	s.Index(j).Set(reflect.ValueOf(x))
}

func (ss sortbase) IsSorted(s interface{}, fieldName string) bool {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	for i := 1; i < slice.Len(); i++ {
		if ss.less(slice.Index(i), slice.Index(i-1), fieldName) {
			return false
		}
	}
	return true
}
func (_ sortbase) Show(s interface{}) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	for i := 0; i < slice.Len(); i++ {
		fmt.Printf("%v ", slice.Index(i))
	}
	fmt.Println()
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
