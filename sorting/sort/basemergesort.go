// basemergesort
package sort

import (
	"reflect"
)

type basemergesort struct {
	sortbase
}

func (m basemergesort) Sort(s interface{}, fieldName string) {}
func (m basemergesort) innerSort(s reflect.Value, fieldName string, aux reflect.Value, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	m.innerSort(s, fieldName, aux, lo, mid)
	m.innerSort(s, fieldName, aux, mid+1, hi)
	if !m.lessequal(s.Index(mid), s.Index(mid+1), fieldName) {
		m.merge(s, fieldName, aux, lo, mid, hi)
	}
}
func (m basemergesort) merge(slice reflect.Value, fieldName string, aux reflect.Value, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		aux.Index(k).Set(slice.Index(k))
	}
	for k := lo; k <= hi; k++ {
		if i > mid {
			slice.Index(k).Set(aux.Index(j))
			j++
		} else if j > hi {
			slice.Index(k).Set(aux.Index(i))
			i++
		} else if m.less(aux.Index(j), aux.Index(i), fieldName) {
			slice.Index(k).Set(aux.Index(j))
			j++
		} else {
			slice.Index(k).Set(aux.Index(i))
			i++
		}
	}
}
func (ss basemergesort) lessequal(fi reflect.Value, fj reflect.Value, fieldName string) bool {
	var field string
	switch fi.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fi.Int() <= fj.Int()
	case reflect.String:
		return fi.String() <= fj.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fi.Uint() <= fj.Uint()
	case reflect.Float32, reflect.Float64:
		return fi.Float() <= fj.Float()
	case reflect.Struct:
		field = fieldName
		return ss.less(fi.FieldByName(field), fj.FieldByName(field), field)
	default:
		panic("Unsupported data or struct's field not found")
	}
}
