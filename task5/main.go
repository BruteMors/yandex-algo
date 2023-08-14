package main

import (
	"fmt"
)

type node struct {
	val   int
	right *node
	left  *node
}

type queue struct {
	arr  []*node
	size int
	pos  int
}

func newQueue(size int) queue {
	arr := make([]*node, 0, size)
	return queue{arr: arr,
		size: 0,
		pos:  0}
}

func (q *queue) insert(n *node) {
	q.arr = append(q.arr, n)
	q.size++
}

func (q *queue) get() *node {
	elem := q.arr[q.pos]
	q.pos++
	q.size--
	return elem
}

func (q *queue) isEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func main() {
	root := node{
		val: 1,
		right: &node{
			val: 2,
			right: &node{
				val:   3,
				right: nil,
				left: &node{
					val:   4,
					right: nil,
					left:  nil,
				},
			},
			left: nil,
		},
		left: &node{
			val: 5,
			right: &node{
				val:   6,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	}

	fmt.Println(printTree(&root))

}

// Обход дерева в ширину
func printTree(root *node) []int {
	arr := make([]int, 0, 10)
	q := newQueue(10)
	q.insert(root)
	for !q.isEmpty() {
		n := q.get()
		arr = append(arr, n.val)
		if n.left != nil {
			q.insert(n.left)
		}
		if n.right != nil {
			q.insert(n.right)
		}
	}
	return arr
}
