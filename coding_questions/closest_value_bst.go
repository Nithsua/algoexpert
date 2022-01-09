package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func min(target int, value ...int) int {
	min := int(^uint(0) >> 1)
	index := 0
	fmt.Println(value)
	for i, v := range value {
		if temp := int(math.Abs(float64(target - v))); temp < min {
			min = temp
			index = i
		}
	}

	return value[index]
}

func (tree *BST) TraverseTree(target int) int {
	min := tree.Value
	p := tree
	for p != nil {
		if p.Value == target {
			min = p.Value
			break
		} else {
			temp1 := int(math.Abs(float64(p.Value - target)))
			temp2 := int(math.Abs(float64(min - target)))
			if temp1 < temp2 {
				min = p.Value
			}
			if p.Value < target {
				p = p.Right
			} else {
				p = p.Left
			}
		}
	}

	return min
}

func (tree *BST) FindClosestValue(target int) int {
	fmt.Println(tree.Value)
	min := tree.TraverseTree(target)

	return min
}
