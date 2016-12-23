package binarysearch

import (
	"errors"
	"fmt"
)

var (
	keys []string
	vals []int
)

func New() {
	keys = make([]string, 0)
	vals = make([]int, 0)
}
func GetKeys() []string {
	return keys
}
func GetVals() []int {
	return vals
}

func Get(key string) (int, error) {
	if IsEmpty() {
		return 99, errors.New("Data is empty.")
	}
	i := rank(key)
	if i < len(keys) && keys[i] == key {
		return vals[i], nil
	} else {
		return 99, errors.New("Not Found")
	}
}
func Contain(key string) bool {
	if _, err := Get(key); err != nil {
		return false
	}
	return true
}
func Put(key string, value int) {
	n := len(keys)
	if n == 0 {
		keys = append(keys, key)
		vals = append(vals, value)
		return
	}
	i := rank(key)
	if i < n && keys[i] == key {
		vals[i] = value
		return
	}
	if i == n {
		keys = append(keys, key)
		vals = append(vals, value)
		return
	}
	keys = append(keys, "")
	vals = append(vals, 0)
	for j := n; j > i; j-- {
		keys[j], vals[j] = keys[j-1], vals[j-1]
	}
	keys[i] = key
	vals[i] = value
}
func rank(key string) int {
	lo, hi := 0, len(keys)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < keys[mid] {
			hi = mid - 1
		} else if key > keys[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}
func Min() string {
	return keys[0]
}
func Max() string {
	return keys[len(keys)-1]
}
func Select(k int) string {
	return keys[k]
}
func Ceiling(key string) string {
	i := rank(key)
	return keys[i]
}
func IsEmpty() bool {
	if len(vals) == 0 {
		return true
	}
	return false
}
func Print() {
	fmt.Println("keys", keys)
	fmt.Println("vals", vals)
}
