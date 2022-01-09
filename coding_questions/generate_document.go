package main

func GenerateDocument(characters string, document string) bool {
	characterMap := map[rune]int{}
	documentMap := map[rune]int{}

	for _, v := range characters {
		characterMap[v] += 1
	}
	for _, v := range document {
		documentMap[v] += 1
	}

	for k, v := range documentMap {
		if w, ok := characterMap[k]; ok {
			if v > w {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
