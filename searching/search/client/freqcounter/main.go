// main
package main

import (
	_ "algorityhm/searching/search/binarysearch"
	"algorityhm/searching/search/bst"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	length := flag.Int("len", 1, "length of word")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	//bst.New()
	bst.Init()
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) < *length {
			continue
		}
		//fmt.Println(word)
		if !bst.Contain(word) {
			bst.Put(word, 1)

			//fmt.Println(bst.GetKeys())
			//fmt.Println(bst.GetVals())
		} else {
			val, _ := bst.Get(word)
			bst.Put(word, val+1)
		}

	}

	max := ""
	bst.Put(max, 0)
	/*for _, word := range bst.GetKeys() {
		val, _ := bst.Get(word)
		maxval, _ := bst.Get(max)
		//fmt.Println("curval", val)
		//fmt.Println("maxval", maxval)
		if val > maxval {
			max = word
		}
	}*/
	nodes := bst.Iter()
	for _, node := range nodes {
		val, _ := bst.Get(node.Key)
		maxval, _ := bst.Get(max)
		if val > maxval {
			max = node.Key
		}
	}
	maxval, _ := bst.Get(max)
	fmt.Printf("%s : %v\n", max, maxval)
	elapsed := time.Since(start)
	log.Printf("Binary Search took %s", elapsed)
}
