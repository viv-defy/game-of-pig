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
		return 0, 0, fmt.Errorf("error parsing start of range: %v", err)
	}

	if len(parts) != 2 {
		return start, start, nil
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing end of range: %v", err)
	}

	return start, end, nil
}

func main() {
	flag.Parse()
	strategies := flag.Args()

	for i, v := range strategies {
		start, end, err := parseRange(v)
		if err != nil {
			fmt.Printf("invalid input %v: %v\n", i+1, err)
			return
		}
		fmt.Printf("strategy %v: %v-%v\n", i+1, start, end)
	}
}
