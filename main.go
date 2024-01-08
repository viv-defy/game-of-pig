package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func game(p1, p2 player) player {
	t := 1
	for {
		if t == 1 {
			p1.turn()
			if p1.score >= 100 {
				return p1
			}
		} else {
			p2.turn()
			if p2.score >= 100 {
				return p2
			}
		}
		t *= -1
	}
}

func simulateGames(p1Hold, p2Hold int) int {
	p1 := player{id: 1, hold: p1Hold}
	p2 := player{id: 2, hold: p2Hold}
	var wins int
	for i := 0; i < 10; i++ {
		if winner := game(p1, p2); winner.id == 1 {
			wins += 1
		}
		p1.reset()
		p2.reset()
	}
	return wins
}

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func fixStrategies(start1, end1, start2, end2 int) {
	var gamesPlayed, wins float64
	printPerHold := start1 == end1 || start2 == end2

	for i := start1; i <= end1; i++ {
		for j := start2; j <= end2; j++ {
			if i == j {
				continue
			}
			result := simulateGames(i, j)
			if printPerHold {
				fmt.Printf("Holding at %v vs Holding at %v: wins: %v/10 (%.1f%%), losses: %v/10 (%.1f%%)\n",
					i, j,
					result, float64(result)/10.0*100.0,
					10-result, float64(10-result)/10.0*100.0)
			}

			wins += float64(result)
			gamesPlayed += 10.0
		}

		if !printPerHold {
			fmt.Printf("Result: Wins, losses staying at k = %v: %v/%v (%.1f%%), %v/%v (%.1f%%)\n", i,
				wins, gamesPlayed, wins/gamesPlayed*100.0,
				gamesPlayed-wins, gamesPlayed, (gamesPlayed-wins)/gamesPlayed*100.0)
		}
		gamesPlayed = 0.0
		wins = 0.0
	}
}

func parseRange(input string) (int, int, error) {
	parts := strings.Split(input, "-")

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start of range: %v", err)
	}

	if len(parts) != 2 {
		return start, start, nil
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid end of range: %v", err)
	}

	return start, end, nil
}

func parseStrategies(strats []string) (int, int, int, int, error) {
	if len(strats) != 2 {
		return 0, 0, 0, 0, fmt.Errorf("invalid number of strategies")
	}

	var start, end []int
	for _, v := range strats {
		s, e, err := parseRange(v)
		if err != nil {
			return 0, 0, 0, 0, err
		}
		if s > e || s < 1 {
			return 0, 0, 0, 0, fmt.Errorf("invalid input strategy")
		}
		start = append(start, s)
		end = append(end, findMin(e, 100))
	}

	return start[0], end[0], start[1], end[1], nil
}

func main() {
	flag.Parse()
	strategies := flag.Args()
	start1, end1, start2, end2, err := parseStrategies(strategies)
	if err != nil {
		panic(err)
	}

	fixStrategies(start1, end1, start2, end2)
}
