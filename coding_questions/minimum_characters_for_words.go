package main

/* Brute Force */
func MinimumCharactersForWords(words []string) []string {
	charSet := map[rune]int{}
	
	for _, word := range words {
		tempSet := map[rune]int{}
		for _, char := range word {
			if _, ok := tempSet[char]; ok {
				tempSet[char]++
			} else {
				tempSet[char] = 1
			}
		}
		for k, v := range tempSet {
			if w, ok := charSet[k]; ok {
				if v > w {
					charSet[k] = v
				}
			} else {
				charSet[k] = v
			}
		}
	}
	
	arr := []string{}
	for k, v := range charSet {
		for i := 0; i < v; i++ {
			arr = append(arr, string(k))
		}
	}
	
	return arr
}

