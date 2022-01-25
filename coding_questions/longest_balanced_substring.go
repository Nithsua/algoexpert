package main

/*
******Bruteforce******
use a two pointer approach and check the balance of all the possible
substrings

use a stack to check the balance of the substring
*/

func checkBalance(chars []rune) bool {
	stack := []rune{}
	for _, v := range chars {
		if v == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0
}

func LongestBalancedSubstring(str string) int {
	chars := []rune(str)
	maxCount := 0
	
	for i := 0; i < len(chars); i++ {
		for j := i; j < len(chars); j++ {
			tempStr := chars[i:j+1]
			if tempLen := len(tempStr); tempLen % 2 == 0 && tempLen > maxCount {
				isBalanced := checkBalance(tempStr)
				if isBalanced {
					maxCount = tempLen
				}
			}
		}
	}
	
	return maxCount
}

