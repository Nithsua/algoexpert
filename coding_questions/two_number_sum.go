package main

// 2 Traversal
func TwoNumberSum(array []int, target int) []int {
	result := []int{}
	valueMap := map[int]int{}

	for i, v := range array {
		valueMap[v] = i
	}

	for _, v := range array {
		temp := target - v
		if _, ok := valueMap[temp]; temp != v && ok {
			result = append(result, temp)
			result = append(result, v)
			break
		}
	}

	return result
}

// 1 Traversal
func BetterTwoNumberSum(array []int, target int) []int {
	result := []int{}
	valueMap := map[int]int{}

	for i, v := range array {
		temp := target - v
		if _, ok := valueMap[temp]; temp != v && ok {
			result = append(result, temp)
			result = append(result, v)
			break
		}
		valueMap[v] = i
	}

	return result
}
