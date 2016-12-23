// MergeSortBottomUp
package sort

import (
	"math"
	"reflect"
)

type MergeSortBottomUp struct {
	basemergesort
}

//Best use for large data,time required for sorting = NlogN,where N is data length
func (m MergeSortBottomUp) Sort(s interface{}, fieldName string) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	N := slice.Len()
	aux := reflect.MakeSlice(reflect.SliceOf(slice.Index(0).Type()), N+1, N+1)
	for sz := 1; sz < N; sz = sz + sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
			m.merge(slice, fieldName, aux, lo, lo+sz-1, int(math.Min(float64(lo+sz+sz-1), float64(N-1))))
		}
	}
}
