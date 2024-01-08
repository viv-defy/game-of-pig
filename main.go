package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

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
		if s > e {
			return 0, 0, 0, 0, fmt.Errorf("invalid input strategy")
		}
		start = append(start, s)
		end = append(end, e)
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

	fmt.Printf("strat 1: %v-%v\nstrat 2: %v-%v\n", start1, end1, start2, end2)
}
