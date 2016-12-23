// bst
package bst

import (
	"errors"
)

type Node struct {
	Key   string
	Value int
	Left  *Node
	Right *Node
	N     int
}

var root *Node
var nodes []*Node

func Init() {
	root = nil
}
func GetRoot() *Node {
	return root
}

func newNode(key string, val, size int) *Node {
	node := &Node{
		key, val, nil, nil, size,
	}
	return node
}
func Get(key string) (int, error) {
	return get(root, key)
}
func get(x *Node, key string) (int, error) {
	if x == nil {
		return 99, errors.New("Not Found")
	}
	if key < x.Key {
		return get(x.Left, key)
	} else if key > x.Key {
		return get(x.Right, key)
	} else {
		return x.Value, nil
	}
}
func Put(key string, val int) {
	root = put(root, key, val)
}
func put(x *Node, key string, val int) *Node {
	if x == nil {
		return newNode(key, val, 1)
	}
	if key < x.Key {
		x.Left = put(x.Left, key, val)
	} else if key > x.Key {
		x.Right = put(x.Right, key, val)
	} else {
		x.Value = val
	}
	x.N = size(x.Left) + size(x.Right) + 1
	return x
}
func Contain(key string) bool {
	if _, err := Get(key); err != nil {
		return false
	}
	return true
}
func Size() int {
	return size(root)
}
func size(x *Node) int {
	if x == nil {
		return 0
	} else {
		return x.N
	}
}
func Min() string {
	return min(root).Key
}
func min(x *Node) *Node {
	if x.Left == nil {
		return x
	}
	return min(x.Left)
}
func Max() string {
	return max(root).Key
}
func max(x *Node) *Node {
	if x.Right == nil {
		return x
	}
	return max(x.Right)
}
func Floor(key string) (string, error) {
	x := floor(root, key)
	if x != nil {
		return x.Key, nil
	}
	return "", errors.New("NotFound")
}
func floor(x *Node, key string) *Node {
	if x == nil {
		return nil
	}
	if x.Key == key {
		return x
	}
	if x.Key > key {
		return floor(x.Left, key)
	}
	t := floor(x.Right, key)
	if t != nil {
		return t
	} else {
		return x
	}
}
func Ceil(key string) (string, error) {
	x := ceil(root, key)
	if x != nil {
		return x.Key, nil
	}
	return "", errors.New("NotFound")
}
func ceil(x *Node, key string) *Node {
	if x == nil {
		return nil
	}
	if x.Key == key {
		return x
	}
	if x.Key < key {
		return ceil(x.Right, key)
	}
	t := ceil(x.Left, key)
	if t != nil {
		return t
	} else {
		return x
	}
}
func Select(k int) string {
	return selected(root, k).Key
}
func selected(x *Node, k int) *Node {
	//return node containing key of rank k.
	if x == nil {
		return nil
	}
	t := size(x.Left)
	if t > k {
		return selected(x.Left, k)
	} else if t < k {
		return selected(x.Right, k-t-1)
	} else {
		return x
	}
}
func Rank(key string) int {
	return rank(root, key)
}
func rank(x *Node, key string) int {
	// Return number of keys less than x.key in the subtree rooted at x.
	if x == nil {
		return 0
	}
	if key < x.Key {
		return rank(x.Left, key)
	} else if key > x.Key {
		return 1 + size(x.Left) + rank(x.Right, key)
	} else {
		return size(x.Left)
	}
}

func Iter() []*Node {
	nodes = make([]*Node, 0)
	iter(root)
	return nodes
}
func iter(x *Node) {
	if x == nil {
		return
	}
	nodes = append(nodes, x)
	iter(x.Left)
	iter(x.Right)
}
