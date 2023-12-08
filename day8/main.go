package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("./day8/input"))

}

type networkMap = map[string]map[string]string

var handRegex = regexp.MustCompile(`[A-Z]{3}`)

func part1(lines []string) int {
	instructions := strings.Split(lines[0], "")

	fmt.Printf("got instructions %v\n", instructions)

	myNetworkMap := networkMap{}

	for i := 2; i < len(lines); i++ {
		matches := handRegex.FindAllString(lines[i], -1)
		myNetworkMap[matches[0]] = map[string]string{"L": matches[1], "R": matches[2]}
	}

	steps := getStepsForElement("AAA", instructions, myNetworkMap)
	fmt.Printf("Step Count = %d\n", steps)
	return steps
}

func getStepsForElement(element string, instructionList []string, myNetworkMap networkMap) int {
	steps := 0
	currentElement := element
	for {
		for _, instruction := range instructionList {
			if currentElement == "ZZZ" {
				return steps
			}

			steps += 1
			currentElement = myNetworkMap[currentElement][instruction]
		}
	}
}
