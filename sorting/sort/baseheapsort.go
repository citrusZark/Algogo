package sort

import (
	"reflect"
)

type baseheapsort struct {
	sortbase
}

func (b baseheapsort) sink(s reflect.Value, fieldName string, k, n int) {
	for 2*k+1 <= n {
		j := 2*k + 1
		if j < n && b.less(s.Index(j), s.Index(j+1), fieldName) {
			j++
		}
		if !b.less(s.Index(k), s.Index(j), fieldName) {
			break
		}
		b.swap(k, j, s)
		k = j
	}
}
