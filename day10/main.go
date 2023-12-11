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
		newLine := processLine(valuesList)
		childSum += newLine
	}
	fmt.Printf("TOTAL: %d\n", childSum)
	return childSum

}

func part2(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))

	childSum := 0
	for _, line := range lines {
		//fmt.Printf(".............................................\n")
		valuesList := util.ReverseInts(util.StringsToInts(line))
		newLine := processLine(valuesList)
		childSum += newLine
	}
	fmt.Printf("TOTAL: %d\n", childSum)
	return childSum

}

func processLine(valuesList []int) (newVal int) {
	if !slices.ContainsFunc(valuesList, func(i int) bool {
		return i != 0
	}) {
		//fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", 0, 0, valuesList)
		//if everything zero, so new val zero, and no need to process diffs.
		return 0
	}

	diffList := getDiffList(valuesList)

	childNewVal := processLine(diffList)

	newVal = valuesList[len(valuesList)-1] + childNewVal

	//fmt.Printf("Returning newVal: %d, for list %v\n", newVal, valuesList)
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
