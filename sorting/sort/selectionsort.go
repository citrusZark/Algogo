// selectionsort
package sort

import (
	"reflect"
)

type SelectionSort struct {
	sortbase
}

func (ss SelectionSort) Sort(s interface{}, fieldName string) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	for i := 0; i < slice.Len(); i++ {
		min := i
		for j := i + 1; j < slice.Len(); j++ {

			if ss.less(slice.Index(j), slice.Index(min), fieldName) {
				ss.swap(i, j, slice)
			}
		}
	}
}
