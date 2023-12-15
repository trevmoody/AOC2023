package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"strconv"
	"strings"
)

func main() {
	f := *util.GetFileAsLines("input")
	part1(f[0])
	part2(f[0])
}

type lensData struct {
	label       string
	focalLength int
	position    int
}

type box struct {
	lensLabelMap map[string]lensData
}

func part1(line string) int {

	result := 0
	for _, step := range strings.Split(line, ",") {
		result += labelToBoxId(step)
	}

	fmt.Printf("Part 1 Sum: %d\n", result)
	return result
}

func part2(line string) int {

	boxes := make([]box, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = box{map[string]lensData{}}
	}

	for _, step := range strings.Split(line, ",") {
		operation := step[strings.IndexAny(step, "=-")]
		splitStep := strings.Split(step, string(operation))
		lensLabel := splitStep[0]

		boxId := labelToBoxId(lensLabel)

		switch operation {
		case '-':
			lensToRemove, ok := boxes[boxId].lensLabelMap[lensLabel]
			if ok {
				delete(boxes[boxId].lensLabelMap, lensLabel)
				for key, lensToCheck := range boxes[boxId].lensLabelMap {
					if lensToCheck.position > lensToRemove.position {
						boxes[boxId].lensLabelMap[key] = lensData{lensToCheck.label, lensToCheck.focalLength, lensToCheck.position - 1}
					}
				}
			}

		case '=':
			focalLengthToAdd, err := strconv.Atoi(splitStep[1])
			if err != nil {
				panic("cant parse int")
			}
			currentLens, ok := boxes[boxId].lensLabelMap[lensLabel]
			if ok { // need to replace it in the list with same position
				boxes[boxId].lensLabelMap[lensLabel] = lensData{lensLabel, focalLengthToAdd, currentLens.position}
			} else { //need to add to the end of the list
				boxes[boxId].lensLabelMap[lensLabel] = lensData{lensLabel, focalLengthToAdd, len(boxes[boxId].lensLabelMap)}
			}
		}
	}

	focusPower := calculateFocusPower(boxes)
	fmt.Printf("Part 2 Power: %d\n", focusPower)

	return focusPower
}

func calculateFocusPower(boxes []box) int {
	sum := 0
	for index, box := range boxes {
		if len(box.lensLabelMap) != 0 {
			for _, lens := range box.lensLabelMap {
				sum += (1 + index) * (lens.position + 1) * lens.focalLength
			}
		}
	}
	return sum
}

func labelToBoxId(label string) int {
	currentValue := 0
	for _, char := range label {
		ascii := int(char)
		currentValue += ascii
		currentValue *= 17
		currentValue = currentValue % 256
	}
	return currentValue
}
