package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type machinePart struct {
	x, m, a, s int
}

type workflowData struct {
	name  string
	rules []rule
}

type rule struct {
	rawRule          string
	destination      string
	machinePartField string
	operand          string
	amount           int
	funcToRun        func(part machinePart) string
}

type Range struct {
	min, max int
}

func (r1 Range) length() int {
	return r1.max + 1 - r1.min
}

func (r1 Range) intersect(r2 Range) Range {
	if r2.min > r1.max || r1.min > r2.max {
		return Range{-1, -1}
	} else {
		return Range{min: int(math.Max(float64(r1.min), float64(r2.min))), max: int(math.Min(float64(r1.max), float64(r2.max)))}
	}
}

type MachinePartRanges struct {
	x, m, a, s Range
}

func (mpr MachinePartRanges) permutations() int {
	return mpr.m.length() * mpr.x.length() * mpr.a.length() * mpr.s.length()
}

func main() {
	lines := *util.GetFileAsLines("input")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0

	var workflowLines []string
	var ratingLines []string
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			workflowLines = lines[0:i]
			ratingLines = lines[i+1:]
		}
	}

	workflows := parseWorkflowDatas(workflowLines)
	ratings := parseRatingLine(ratingLines)

	for _, rating := range ratings {
		total += processRating(workflows, rating)
	}

	fmt.Printf("Part 1 Ratings Total %d \n", total)
}

func part2(lines []string) {
	var workflowLines []string

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			workflowLines = lines[0:i]
		}
	}

	workflowDatas := parseWorkflowDatas(workflowLines)

	startingRange := MachinePartRanges{
		Range{1, 4000},
		Range{1, 4000},
		Range{1, 4000},
		Range{1, 4000},
	}

	total := 0

	results := processWorkflow(workflowDatas, startingRange, "in")

	for _, result := range results {
		total += result.permutations()
	}

	fmt.Printf("Part 2 Accepted Count %d\n", total)
}

func processWorkflow(workflowDatas map[string]workflowData, startingRange MachinePartRanges, workFlowName string) []MachinePartRanges {

	if startingRange.x.min == -1 || startingRange.m.min == -1 || startingRange.a.min == -1 || startingRange.s.min == -1 {
		return nil
	}

	if workFlowName == "R" {
		return nil
	}

	if workFlowName == "A" {
		return []MachinePartRanges{startingRange}
	}

	workflowToProcess := workflowDatas[workFlowName]

	accepted := processRule(workflowDatas, workflowToProcess.rules[0], startingRange, workflowToProcess.rules[1:])

	return accepted

}

func processRule(workflowDatas map[string]workflowData, ruleToProcess rule, inputRange MachinePartRanges, nextRules []rule) []MachinePartRanges {

	// this caters for the case where the default is A/R
	if ruleToProcess.machinePartField == "" {
		if ruleToProcess.destination == "A" {
			return []MachinePartRanges{inputRange}
		}
		if ruleToProcess.destination == "R" {
			return nil
		}
		//must be another workflow
		return processWorkflow(workflowDatas, inputRange, ruleToProcess.destination)

	} else {

		// some of the range gets passed to next workflow
		rangeThatPassesRule, rangeThatFailsRule := splitRange(inputRange, ruleToProcess)
		rangeFromFollowingWorkflow := processWorkflow(workflowDatas, rangeThatPassesRule, ruleToProcess.destination)

		// the rest goes to the next rule.
		rangeFromFollowingRule := processRule(workflowDatas, nextRules[0], rangeThatFailsRule, nextRules[1:])

		// merge them together
		return append(rangeFromFollowingWorkflow, rangeFromFollowingRule...)
	}

}

func splitRange(inputRange MachinePartRanges, process rule) (MachinePartRanges, MachinePartRanges) {

	var rangeToMergePassRule Range
	var rangeToMergeNotPassRule Range

	switch process.operand {

	case "<":
		{
			rangeToMergePassRule = Range{1, process.amount - 1}
			rangeToMergeNotPassRule = Range{process.amount, 4000}
		}
	case ">":
		{
			rangeToMergePassRule = Range{process.amount + 1, 4000}
			rangeToMergeNotPassRule = Range{1, process.amount}
		}
	}

	switch process.machinePartField {
	case "x":
		{
			return MachinePartRanges{x: inputRange.x.intersect(rangeToMergePassRule), m: inputRange.m, a: inputRange.a, s: inputRange.s},
				MachinePartRanges{x: inputRange.x.intersect(rangeToMergeNotPassRule), m: inputRange.m, a: inputRange.a, s: inputRange.s}

		}
	case "m":
		{
			return MachinePartRanges{x: inputRange.x, m: inputRange.m.intersect(rangeToMergePassRule), a: inputRange.a, s: inputRange.s},
				MachinePartRanges{x: inputRange.x, m: inputRange.m.intersect(rangeToMergeNotPassRule), a: inputRange.a, s: inputRange.s}

		}
	case "a":
		{
			return MachinePartRanges{x: inputRange.x, m: inputRange.m, a: inputRange.a.intersect(rangeToMergePassRule), s: inputRange.s},
				MachinePartRanges{x: inputRange.x, m: inputRange.m, a: inputRange.a.intersect(rangeToMergeNotPassRule), s: inputRange.s}

		}
	case "s":
		{
			return MachinePartRanges{x: inputRange.x, m: inputRange.m, a: inputRange.a, s: inputRange.s.intersect(rangeToMergePassRule)},
				MachinePartRanges{x: inputRange.x, m: inputRange.m, a: inputRange.a, s: inputRange.s.intersect(rangeToMergeNotPassRule)}

		}
	default:
		panic("aaarhg")

	}

}

func parseWorkflowDatas(lines []string) map[string]workflowData {
	workFlowMap := map[string]workflowData{}

	for i := 0; i < len(lines); i++ {
		matches := wfre.FindStringSubmatch(lines[i])
		wfName := matches[1]
		wfDefinition := matches[2]

		var ruleList []rule

		for _, split := range strings.Split(wfDefinition, ",") {
			innerMatches := wfre2.FindStringSubmatch(split)
			if len(innerMatches) == 0 {
				ruleList = append(ruleList, rule{
					rawRule:     split,
					destination: split,
					funcToRun:   func(xmas machinePart) string { return split }, // this is not right.
				})
			} else {
				ruleList = append(ruleList, buildRule(split, innerMatches[1], innerMatches[2], innerMatches[3], innerMatches[4]))
			}
		}

		workFlowMap[wfName] = workflowData{name: wfName, rules: ruleList}

	}

	return workFlowMap
}

func processRating(workflows map[string]workflowData, rating machinePart) int {

	workflowToCheck := workflows["in"]
	for {
		for _, rule := range workflowToCheck.rules {
			result := rule.funcToRun(rating)
			if result == "R" {
				return 0
			}
			if result == "A" {
				return rating.x + rating.m + rating.a + rating.s
			}
			if result == "" {
				// next check in the workflow
				continue
			}

			workflowToCheck = workflows[result]
			break

		}
	}
}

var ratingRe = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

func parseRatingLine(lines []string) []machinePart {
	var xmasRatings []machinePart

	for _, line := range lines {
		matches := ratingRe.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		m, _ := strconv.Atoi(matches[2])
		a, _ := strconv.Atoi(matches[3])
		s, _ := strconv.Atoi(matches[4])

		xmasRatings = append(xmasRatings, machinePart{x, m, a, s})
	}

	return xmasRatings
}

var wfre = regexp.MustCompile(`(^\w*){(.*)}`)
var wfre2 = regexp.MustCompile(`(^\w)([<>])(\w*):(\w*)`)

func buildRule(rawRule string, xmasField string, operand string, amountStr string, destination string) rule {
	amount, _ := strconv.Atoi(amountStr)

	f := func(xmas machinePart) string {
		var xmasFieldVal int
		switch xmasField {
		case "x":
			xmasFieldVal = xmas.x
		case "m":
			xmasFieldVal = xmas.m
		case "a":
			xmasFieldVal = xmas.a
		case "s":
			xmasFieldVal = xmas.s
		default:
			panic("Unknown rating")
		}

		switch operand {
		case "<":
			if xmasFieldVal < amount {
				return destination
			}
		case ">":
			if xmasFieldVal > amount {
				return destination
			}
		default:
			panic("aaargh")
			return ""
		}

		return ""
	}

	return rule{rawRule: rawRule, funcToRun: f, destination: destination, amount: amount, machinePartField: xmasField, operand: operand}

}
