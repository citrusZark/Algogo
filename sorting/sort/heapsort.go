// heapsort
package sort

import (
	"reflect"
)

type HeapSort struct {
	baseheapsort
}

func (h HeapSort) Sort(s interface{}, fieldName string) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	//slice.Index(0).Set(reflect.ValueOf(""))
	n := slice.Len() - 1
	for k := n / 2; k >= 0; k-- {
		h.sink(slice, fieldName, k, n)
	}
	for n >= 1 {
		h.swap(0, n, slice)
		n--
		h.sink(slice, fieldName, 0, n)

	}
}
