package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for i, v := range args {
		fmt.Printf("arg %v - %v\n", i, v)
	}
}
