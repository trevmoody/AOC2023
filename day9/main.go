package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

	part2(*util.GetFileAsLines("input"))

}
func part1(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))

	childsum := 0

	for _, line := range lines {
		fmt.Printf(".............................................\n")
		valuesList := util.StringsToInts(line)
		_, sum := processLine1(valuesList)
		childsum += sum
	}

	fmt.Printf("TOTAL: %d\n", childsum)
	return childsum

}

func part2(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))

	childsum := 0

	for _, line := range lines {
		fmt.Printf(".............................................\n")
		valuesList := util.StringsToInts(line)
		newVal, _ := processLine2(valuesList)
		childsum += newVal
	}

	fmt.Printf("TOTAL: %d\n", childsum)
	return childsum

}

func processLine1(valuesList []int) (diff int, sum int) {

	diffList := make([]int, len(valuesList)-1)

	foundNonZero := false
	for _, valToCheck := range valuesList {
		if valToCheck != 0 {
			foundNonZero = true
		}
	}

	if !foundNonZero {
		fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", 0, 0, valuesList)
		return 0, 0
	}
	//skip 0
	for i := 1; i < len(valuesList); i++ {
		diff := valuesList[i] - valuesList[i-1]
		diffList[i-1] = diff
	}

	childDiff, sum := processLine1(diffList)
	// so my new Val, is last val + diff

	newVal := childDiff + valuesList[len(valuesList)-1]
	newSum := sum + newVal

	fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", newVal, newSum, valuesList)
	return childDiff, newSum
}

func processLine2(valuesList []int) (diff int, sum int) {

	foundNonZero := false
	for _, valToCheck := range valuesList {
		if valToCheck != 0 {
			foundNonZero = true
		}
	}

	if !foundNonZero {
		fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", 0, 0, valuesList)
		return 0, 0
	}

	diffList := make([]int, len(valuesList)-1)
	//skip 0
	for i := 1; i < len(valuesList); i++ {
		diff := valuesList[i] - valuesList[i-1]
		diffList[i-1] = diff
	}

	childNewVal, sum := processLine2(diffList)
	// so my new Val, is last val + diff

	newVal := valuesList[0] - childNewVal
	newSum := sum + newVal

	fmt.Printf("Returning newVal: %d, sum: %d for list %v\n", newVal, newSum, valuesList)
	return newVal, newSum
}
