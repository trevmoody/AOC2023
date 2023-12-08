package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

}

type element string

var handRegex = regexp.MustCompile(`[A-Z]{3}`)

func part1(lines []string) int {
	instructions := lines[0]

	fmt.Printf("got instructions %v\n", instructions)

	networkMap := map[string]map[string]string{} // size 2

	for i := 2; i < len(lines); i++ {
		matches := handRegex.FindAllString(lines[i], -1)
		networkMap[matches[0]] = map[string]string{"L": matches[1], "R": matches[2]}
	}

	steps := 0
	currentElement := "AAA"
	destElement := "ZZZ"

	for true {
		for _, instruction := range strings.Split(instructions, "") {
			currentElement = networkMap[currentElement][instruction]
			steps += 1
			if currentElement == destElement {
				break
			}
		}
		if currentElement == destElement {
			break
		}
	}

	fmt.Printf("Step Count = %d\n", steps)

	return steps
}
