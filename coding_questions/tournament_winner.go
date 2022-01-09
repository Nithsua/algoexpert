package main

import (
	"math"
)

func TournamentWinner(competitions [][]string, results []int) string {
	teamMap := make(map[string]int)
	for i, v := range results {
		teamMap[competitions[i][int(math.Abs(float64(v-1)))]] += 1
	}

	winningTeam := ""
	max := 0
	for k, v := range teamMap {
		if v > max {
			max = v
			winningTeam = k
		}
	}
	return winningTeam
}
