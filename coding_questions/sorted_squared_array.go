package main

import (
	"math"
	"sort"
)

//Optimal
func BetterSortedSquaredArray(array []int) []int {
	result := []int{}
	small := 0
	large := len(array) - 1

	for len(result) < len(array) {
		s := int(math.Abs(float64(array[small])))
		l := int(math.Abs(float64(array[large])))
		if s > l {
			result = append([]int{s * s}, result...)
			small++
		} else {
			result = append([]int{l * l}, result...)
			large--
		}
	}

	return result
}

//Brute Force
func SortedSquaredArray(array []int) []int {
	for i, v := range array {
		array[i] = v * v
	}

	sort.SliceStable(array, func(i, j int) bool { return array[i] < array[j] })

	return array
}
