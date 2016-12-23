// insertionsort
package sort

import (
	"reflect"
)

//Insertion sort works well for certain types of nonrandom arrays that often arise in practice
type InsertionSort struct {
	sortbase
}

func (is InsertionSort) Sort(s interface{}, fieldName string) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	N := slice.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := 1; i < N; i++ {
			for j := i; j >= h && is.less(slice.Index(j), slice.Index(j-h), fieldName); j -= h {
				is.swap(j-h, j, slice)
			}
		}
		h = h / 3
	}

}
