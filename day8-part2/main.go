package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Running Part 2\n")
	part2(*util.GetFileAsLines("input"))

}

type networkMap = map[string]map[string]string

var regex = regexp.MustCompile(`[A-Z0-9]{3}`)

func part2(lines []string) int {
	instructionList := strings.Split(lines[0], "")

	// parse map and find the start elements
	startingElements, myNetworkMap := parseMapAndStartElements(lines)
	//assume cyclical
	var results []int
	for _, currentElement := range startingElements {
		results = append(results, getStepsForElement(currentElement, instructionList, myNetworkMap))
	}

	lcm := getLeastCommonMultiple(results)
	fmt.Printf("Step Count = %d\n", lcm)

	return lcm
}

func parseMapAndStartElements(lines []string) ([]string, networkMap) {
	var startingElements []string
	myNetworkMap := make(networkMap, len(lines)-2)

	for i := 2; i < len(lines); i++ {
		matches := regex.FindAllString(lines[i], -1)
		myNetworkMap[matches[0]] = map[string]string{"L": matches[1], "R": matches[2]}

		if strings.HasSuffix(matches[0], "A") {
			startingElements = append(startingElements, matches[0])
		}
	}
	return startingElements, myNetworkMap
}

func getStepsForElement(element string, instructionList []string, myNetworkMap networkMap) int {
	steps := 0
	currentElement := element
	for {
		for _, instruction := range instructionList {
			if strings.HasSuffix(currentElement, "Z") {
				return steps
			}
			steps += 1
			currentElement = myNetworkMap[currentElement][instruction]
		}
	}
}

// https://en.wikipedia.org/wiki/Least_common_multiple
func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		lcm = lcm * (num2 / GCD(num1, num2))
	}
	return lcm
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
