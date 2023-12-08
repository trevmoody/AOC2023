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

	for i := 0; i < len(currentElements); i++ {
		found := false
		steps := 0
		currentElement := currentElements[i]
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

func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}
	return lcm
}