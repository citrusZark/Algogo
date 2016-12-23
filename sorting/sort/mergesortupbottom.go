// mergesort
package sort

import (
	"reflect"
)

type MergeSortUpBottom struct {
	basemergesort
}

//Best use for large data,time required for sorting = NlogN,where N is data length
func (m MergeSortUpBottom) Sort(s interface{}, fieldName string) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	N := slice.Len()
	aux := reflect.MakeSlice(reflect.SliceOf(slice.Index(0).Type()), N+1, N+1)

	m.innerSort(slice, fieldName, aux, 0, N-1)
}
