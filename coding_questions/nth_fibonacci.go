package main

func GetNthFib(n int) int {
	p1 := 0
	p2 := 1
	if n == 1 {
		return p1
	} else if n == 2 {
		return p2
	} else {
		for i := 0; i < (n - 3); i++ {
			temp := p1
			p1 = p2
			p2 = temp + p1
		}
		return p2 + p1
	}
}
