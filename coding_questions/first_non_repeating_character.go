package main

type Temp struct {
	index    int
	isUnique bool
}

func FirstNonRepeatingCharacter(str string) int {
	characterMap := map[rune]Temp{}
	for i, v := range str {
		if _, ok := characterMap[v]; ok {
			characterMap[v] = Temp{i, false}
		} else {
			characterMap[v] = Temp{i, true}
		}
	}

	intMax := int(^uint(0) >> 1)
	min := intMax
	for _, v := range characterMap {
		if v.isUnique && v.index < min {
			min = v.index
		}
	}

	if min == intMax {
		return -1
	}
	return min
}
