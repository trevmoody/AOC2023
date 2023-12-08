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

var regex = regexp.MustCompile(`[A-Z0-9]{3}`)

func part2(lines []string) int {
	instructions := lines[0]

	fmt.Printf("got instructions %v\n", instructions)

	networkMap := map[string]map[string]string{} // size 2
	var currentElements []string

	for i := 2; i < len(lines); i++ {
		matches := regex.FindAllString(lines[i], -1)
		networkMap[matches[0]] = map[string]string{"L": matches[1], "R": matches[2]}

		if strings.HasSuffix(matches[0], "A") {
			currentElements = append(currentElements, matches[0])
		}
	}

	//assume cyclical
	var results []int

	for _, currentElement := range currentElements {
		found := false
		steps := 0

		for found == false {
			for _, instruction := range strings.Split(instructions, "") {
				if strings.HasSuffix(currentElement, "Z") {
					results = append(results, steps)
					found = true
					break
				}
				steps += 1
				currentElement = networkMap[currentElement][instruction]
			}
		}
	}

	lcm := getLeastCommonMultiple(results)
	fmt.Printf("Step Count = %d\n", lcm)

	return lcm
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

// greatest common divisor (GCD) via Euclidean algorithm
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
