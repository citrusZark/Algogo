// main
package main

import (
	"algorityhm/sorting/sort"
	"fmt"
)

type person struct {
	Name string
	Age  int
	job
}
type job struct {
	Title  string
	Salary float32
}

func main() {
	var sort sort.HeapSort
	//ii := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	ii := []string{"Q", "U", "I", "C", "K", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	//ii := []string{"K", "R", "A", "T", "E", "L", "E", "P", "U", "I", "M", "Q", "C", "X", "O", "S"}

	//ii := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	//ii := []float32{12.2, 34.5, 67.8, 1.2}
	//ii := 3
	//ii := []bool{true, false, true}
	fmt.Println(ii)
	sort.Sort(ii, "")
	sort.Show(ii)
	fmt.Println(sort.IsSorted(ii, ""))
}
