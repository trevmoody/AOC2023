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

	matches := count(conditionRecord, check, make(map[state]int))

	fmt.Printf("Condition: %s check %v matches: %d \n", conditionRecord, check, matches)
	return matches

}

type state struct {
	pattern string
	numbers string
}

func store(stateStore map[state]int, pattern string, numbers []int, value int) int {
	stateStore[state{pattern, fmt.Sprint(numbers)}] = value
	return value
}

func count(conditionRecord string, restOfcheckList []int, stateStore map[state]int) int {
	//	fmt.Printf("proceesing substring %s, checkSize %d rest of list %v \n", conditionRecord, restOfcheckList)
	if len(conditionRecord) == 0 && len(restOfcheckList) == 0 {
		return 1
	}

	if len(conditionRecord) == 0 {
		return 0
	}

	if value, ok := stateStore[state{conditionRecord, fmt.Sprint(restOfcheckList)}]; ok {
		return value
	}

	currentConditionRecord := conditionRecord[0]
	switch currentConditionRecord {
	case '.':
		return count(conditionRecord[1:], restOfcheckList, stateStore)
	case '?':
		assumeDot := count(conditionRecord[1:], restOfcheckList, stateStore) //  same as a dot.
		assumeHash := count("#"+conditionRecord[1:], restOfcheckList, stateStore)
		return assumeDot + assumeHash // same as a hash
	case '#':

		if len(restOfcheckList) == 0 {
			return 0
		}

		currentCheck := restOfcheckList[0]
		//ok how many non dots.., treating ? as #
		indexDot := strings.Index(conditionRecord, ".")
		if indexDot == -1 {
			//not found asssume end of line
			indexDot = len(conditionRecord)
		}
		if indexDot < currentCheck {
			// not enough # to fill the group
			return 0
		}

		remainingConditionRecord := conditionRecord[currentCheck:]
		if len(remainingConditionRecord) == 0 {
			res := count(remainingConditionRecord, restOfcheckList[1:], stateStore)
			store(stateStore, remainingConditionRecord, restOfcheckList[1:], res)
			return res
		}
		if remainingConditionRecord[0] == '#' {
			// fail this means that out group too big now....
			return 0
		} else {
			// remove the first as is a dot, or a ? that we treat as a dot.
			res := count(remainingConditionRecord[1:], restOfcheckList[1:], stateStore)
			store(stateStore, remainingConditionRecord[1:], restOfcheckList[1:], res)
			return res
		}

	default:
		panic("aaaargh")
	}

	return 0
}
