package main

import (
	"fmt"
)

func BinarySearch(array []int, target int) int {
	index := -1
	l := 0
	r := len(array) - 1
	fmt.Println(array)
	for l <= r {
		mid := l + (r-l)/2
		fmt.Println(mid)
		if array[mid] == target {
			index = mid
			break
		} else if array[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return index
}
