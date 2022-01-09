package main

// Brute Force
func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) < len(sequence) {
		return false
	} else {
		index := 0
		for _, v := range array {
			if sequence[index] == v {
				index++
			}
			if index == len(sequence) {
				return true
			}
		}
	}
	return false
}
