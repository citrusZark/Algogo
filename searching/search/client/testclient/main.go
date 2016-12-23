// main
package main

import (
	_ "algorityhm/searching/search/binarysearch"
	"algorityhm/searching/search/bst"
	"fmt"
)

func main() {
	var inputkey string
	inputkey = "B"
	input := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst.Init()
	i := 0
	for _, key := range input {
		bst.Put(key, i)
		i++
	}
	val, err := bst.Get(inputkey)
	min := bst.Min()
	max := bst.Max()
	floor, _ := bst.Floor(inputkey)
	ceil, _ := bst.Ceil(inputkey)
	nodes := bst.Iter()
	for _, node := range nodes {
		if node != nil {
			fmt.Print(node.Key)
		}

	}
	fmt.Println()
	if err != nil {
		fmt.Printf("val for %s is %s\n", inputkey, err)
	} else {
		fmt.Printf("val for %s is %v\n", inputkey, val)
	}

	fmt.Printf("min for is %s\n", min)
	fmt.Printf("max for is %s\n", max)
	fmt.Printf("floor for %s is %s\n", inputkey, floor)
	fmt.Printf("ceil for %s is %s\n", inputkey, ceil)

}
