package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part2(*util.GetFileAsLines("input"))

}

func part2(lines []string) int {

	sum := 0
	for _, line := range lines {
		sum += getPossibleMatchesCount(line)
	}

	fmt.Printf("TOTAL %d\n", sum)
	return sum
}

func getPossibleMatchesCount(line string) int {

	splitLine := strings.Split(line, " ")

	var conditionRecord string
	conditionSubRecord := splitLine[0]

	var check []int
	for i := 0; i < 5; i++ {
		conditionRecord = conditionRecord + conditionSubRecord + "?"

		for _, i := range strings.Split(splitLine[1], ",") {
			converted, _ := strconv.Atoi(i)
			check = append(check, converted)
		}
	}

	conditionRecord = strings.TrimSuffix(conditionRecord, "?")

	matches := count(conditionRecord, check)

	fmt.Printf("Condition: %s check %v matches: %d \n", conditionRecord, check, matches)
	return matches

}

type state struct {
	pattern string
	numbers string
}

var cache = make(map[state]int)

func store(pattern string, numbers []int, value int) int {
	cache[state{pattern, fmt.Sprint(numbers)}] = value
	return value
}

func count(conditionRecord string, restOfcheckList []int) int {

	//	fmt.Printf("proceesing substring %s, checkSize %d rest of list %v \n", conditionRecord, restOfcheckList)
	if len(conditionRecord) == 0 && len(restOfcheckList) == 0 {
		return 1
	}

	if len(conditionRecord) == 0 {
		return 0
	}

	//cache here.
	// test cache
	if value, ok := cache[state{conditionRecord, fmt.Sprint(restOfcheckList)}]; ok {
		return value
	}

	currentConditionRecord := conditionRecord[0]
	switch currentConditionRecord {
	case '.':
		return count(conditionRecord[1:], restOfcheckList)
	case '?':
		assumeDot := count(conditionRecord[1:], restOfcheckList) //  same as a dot.
		assumeHash := count("#"+conditionRecord[1:], restOfcheckList)
		return assumeDot + assumeHash // same as a hash
	case '#':

		if len(restOfcheckList) == 0 {
			return 0
		}

		currentCheck := restOfcheckList[0]
		//ok how many non dots..
		indexDot := strings.Index(conditionRecord, ".")
		if indexDot == -1 {
			indexDot = len(conditionRecord)
		}
		if indexDot < currentCheck {
			// not enough
			return 0
		}

		// so we have found something.

		remainingConditionRecord := conditionRecord[currentCheck:]
		if len(remainingConditionRecord) == 0 {
			res := count(remainingConditionRecord, restOfcheckList[1:])
			store(remainingConditionRecord, restOfcheckList[1:], res)
			return res
		}
		if remainingConditionRecord[0] == '#' {
			// fail
			return 0
		} else {
			// remove the first as is a dot, or a ? that we treat as a dot.
			res := count(remainingConditionRecord[1:], restOfcheckList[1:])
			store(remainingConditionRecord[1:], restOfcheckList[1:], res)
			return res
		}

	default:
		panic("aaaargh")
	}

	return 0
}
