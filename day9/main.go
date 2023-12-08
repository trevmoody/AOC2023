package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

}

func part1(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))
	return 0
}
