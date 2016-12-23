// quicksort
package sort

import (
	"math/rand"
	"reflect"
	"time"
)

type QuickSort struct {
	basequicksort
}

func (m QuickSort) Sort(s interface{}, fieldName string) {
	slice, ok := takeArg(s, reflect.Slice)
	if !ok {
		panic("error: only take slice as argument")
	}
	n := slice.Len()
	rand.Seed(time.Now().UnixNano())
	pivotIndex := randInt(1, n)
	m.swap(pivotIndex, 0, slice)
	m.innerSort(slice, fieldName, 0, n-1)
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
