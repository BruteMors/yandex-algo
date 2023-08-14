package main

import (
	"fmt"
)

type node struct {
	val   int
	right *node
	left  *node
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

	sum := 0
	sumTree(&root, &sum)
	fmt.Println(sum)

}

// Обход дерева
func sumTree(root *node, sum *int) {
	if root != nil {
		*sum += root.val
		sumTree(root.left, sum)
		sumTree(root.right, sum)
	}

}
