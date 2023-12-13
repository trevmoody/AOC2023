package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
)

func main() {
	part1(*util.GetFileAsLines("input"))
}

func part1(lines []string) {
	patterns := getPatterns(lines)

	sum := 0
	for _, pattern := range patterns {
		sum += findHorizontal(pattern, 0) * 100
		sum += findVertical(pattern, 0) * 1
	}

	fmt.Printf("Part 1 Total: %d\n", sum)

	sum2 := 0
	for _, pattern := range patterns {
		sum2 += findHorizontal(pattern, 1) * 100
		sum2 += findVertical(pattern, 1) * 1
	}

	fmt.Printf("Part 2 Total: %d\n", sum2)

}

func findVertical(pattern []string, expectedDiffCount int) int {
	// pivot the pattern
	var pivot []string

	for index, _ := range pattern[0] {
		pivotRow := ""
		for _, inputRow := range pattern {
			pivotRow += string(inputRow[index])
		}
		pivot = append(pivot, pivotRow)
	}

	return find(pivot, "vertical", expectedDiffCount)
}

func findHorizontal(pattern []string, expectedDiffCount int) int {
	return find(pattern, "horizontal", expectedDiffCount)
}

func find(pattern []string, desc string, expectedDiffCount int) int {

	rowCount := len(pattern)

	for rowId := 0; rowId < rowCount; rowId++ {
		if checkRowDiffCount(rowId, pattern, expectedDiffCount) {
			//fmt.Printf("Found %s Mirror: %d\n", desc, rowId+1)
			return rowId + 1
		}
	}

	return 0
}

func checkRow(rowId int, pattern []string) bool {
	isMirror := false

	for offset := 0; ; offset++ {
		underRowId := rowId - offset
		overRowId := rowId + offset + 1

		if underRowId < 0 || overRowId >= len(pattern) {
			return isMirror
		}

		if pattern[underRowId] == pattern[overRowId] {
			isMirror = true
		} else {
			return false
		}
	}
}

func checkRowDiffCount(rowId int, pattern []string, expectedDiffCount int) bool {

	diffCount := 0
	for offset := 0; ; offset++ {
		underRowId := rowId - offset
		overRowId := rowId + offset + 1

		if underRowId < 0 || overRowId >= len(pattern) {
			return offset != 0 && diffCount == expectedDiffCount
		}

		for colId := 0; colId < len(pattern[underRowId]); colId++ {
			if pattern[underRowId][colId] != pattern[overRowId][colId] {
				diffCount += 1
			}
		}
	}
}

func getPatterns(lines []string) [][]string {
	var patterns [][]string
	var pattern []string
	for _, line := range lines {
		if len(line) == 0 {
			patterns = append(patterns, pattern)
			pattern = []string{}
			continue
		}
		pattern = append(pattern, line)
	}
	patterns = append(patterns, pattern)

	return patterns
}
