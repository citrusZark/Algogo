// basequicksort
package sort

import (
	"reflect"
)

type basequicksort struct {
	sortbase
}

func (m basequicksort) innerSort(s reflect.Value, fieldName string, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := m.partition(s, fieldName, lo, hi)
	m.innerSort(s, fieldName, lo, j-1)
	m.innerSort(s, fieldName, j+1, hi)
}
func (m basequicksort) partition(s reflect.Value, fieldName string, lo, hi int) int {
	a := lo
	b := hi + 1
	v := s.Index(lo)
	for {
		for {
			a++
			if a == hi {
				break
			}
			if m.less(v, s.Index(a), fieldName) {
				break
			}
		}
		for {
			b--
			if b == lo {
				break
			}
			if m.less(s.Index(b), v, fieldName) {
				break
			}

		}
		if a >= b {
			break
		}

		m.swap(a, b, s)

	}
	m.swap(lo, b, s)
	return b

}
