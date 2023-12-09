package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"slices"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

	fmt.Printf("Running Part 2\n")
	part2(*util.GetFileAsLines("input"))

}
func part1(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))

	childSum := 0

	for _, line := range lines {
		//fmt.Printf(".............................................\n")
		valuesList := util.StringsToInts(line)
		_, sum := processLine1(valuesList)
		childSum += sum
	}

	fmt.Printf("TOTAL: %d\n", childSum)
	return childSum

}

func part2(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))

	childSum := 0

	for _, line := range lines {
		//fmt.Printf(".............................................\n")
		valuesList := util.StringsToInts(line)
		newVal := processLine2(valuesList)
		childSum += newVal
	}

	fmt.Printf("TOTAL: %d\n", childSum)
	return childSum

}

func processLine1(valuesList []int) (diff int, sum int) {
	if !slices.ContainsFunc(valuesList, func(i int) bool {
		return i != 0
	}) {
		//fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", 0, 0, valuesList)
		return 0, 0
	}

	diffList := getDiffList(valuesList)

	childDiff, sum := processLine1(diffList)

	newVal := childDiff + valuesList[len(valuesList)-1]
	newSum := sum + newVal

	//fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", newVal, newSum, valuesList)
	return childDiff, newSum
}

func processLine2(valuesList []int) (newVal int) {
	if !slices.ContainsFunc(valuesList, func(i int) bool {
		return i != 0
	}) {
		//fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", 0, 0, valuesList)
		return 0
	}

	diffList := getDiffList(valuesList)

	childNewVal := processLine2(diffList)

	newVal = valuesList[0] - childNewVal

	//fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", newVal, newSum, valuesList)
	return newVal
}

func getDiffList(valuesList []int) []int {
	diffList := make([]int, len(valuesList)-1)
	//skip 0
	for i := 1; i < len(valuesList); i++ {
		diff := valuesList[i] - valuesList[i-1]
		diffList[i-1] = diff
	}
	return diffList
}
