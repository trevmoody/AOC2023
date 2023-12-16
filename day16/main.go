package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
)

func main() {

	part1(*util.GetFileAsLines("input"))
	part2(*util.GetFileAsLines("input"))
}

type Point struct {
	row, col int
}

type PointDirection struct {
	row, col  int
	direction string
}

func part2(lines []string) {
	// start top left

	currentMax := 0

	// row 0
	for colId := 0; colId < len(lines[0]); colId++ {
		energisedPoints := map[Point]bool{}
		previouslyVisited := map[PointDirection]bool{}

		populateEnergisedPointMap(0, colId, lines, "S", energisedPoints, previouslyVisited)

		countForRun := len(energisedPoints)
		if countForRun > currentMax {
			currentMax = countForRun
		}
	}

	// row last
	for colId := 0; colId < len(lines[0]); colId++ {
		energisedPoints := map[Point]bool{}
		previouslyVisited := map[PointDirection]bool{}

		populateEnergisedPointMap(len(lines)-1, colId, lines, "N", energisedPoints, previouslyVisited)

		countForRun := len(energisedPoints)
		if countForRun > currentMax {
			currentMax = countForRun
		}
	}

	// col 0
	for rowId := 0; rowId < len(lines); rowId++ {
		energisedPoints := map[Point]bool{}
		previouslyVisited := map[PointDirection]bool{}

		populateEnergisedPointMap(rowId, 0, lines, "E", energisedPoints, previouslyVisited)

		countForRun := len(energisedPoints)
		if countForRun > currentMax {
			currentMax = countForRun
		}
	}

	// col last
	for rowId := 0; rowId < len(lines); rowId++ {
		energisedPoints := map[Point]bool{}
		previouslyVisited := map[PointDirection]bool{}

		populateEnergisedPointMap(rowId, len(lines[0]), lines, "W", energisedPoints, previouslyVisited)

		countForRun := len(energisedPoints)
		if countForRun > currentMax {
			currentMax = countForRun
		}
	}

	fmt.Printf("Part 2 Max EnergisedTiles Count: %d\n", currentMax)
}

func part1(lines []string) {
	// start top left

	energisedPoints := map[Point]bool{}
	previouslyVisited := map[PointDirection]bool{}

	populateEnergisedPointMap(0, 0, lines, "E", energisedPoints, previouslyVisited)

	fmt.Printf("EnergisedTiles Count: %d\n", len(energisedPoints))
}

func populateEnergisedPointMap(rowId int, colId int, lines []string, direction string, energisedPoints map[Point]bool, previouslyVisited map[PointDirection]bool) {
	if rowId < 0 || rowId >= len(lines) || colId < 0 || colId >= len(lines[0]) {
		return
	}

	_, ok := previouslyVisited[PointDirection{rowId, colId, direction}]
	if ok {
		return
	}

	currentChar := lines[rowId][colId]
	energisedPoints[Point{rowId, colId}] = true
	previouslyVisited[PointDirection{rowId, colId, direction}] = true

	populateEnergisedPointMapForChar(currentChar, rowId, colId, lines, direction, energisedPoints, previouslyVisited)

}

func populateEnergisedPointMapForChar(currentChar uint8, rowId int, colId int, lines []string, direction string, energisedPoints map[Point]bool, previouslyVisited map[PointDirection]bool) {

	var nextDirection string
	var nextRowId int
	var nextColId int

	switch currentChar {
	case '.':
		{
			switch direction {
			case "N":
				nextDirection = "N"

			case "E":
				nextDirection = "E"

			case "S":
				nextDirection = "S"

			case "W":
				nextDirection = "W"
			}

		}

	case '/':
		{
			switch direction {
			case "N":
				nextDirection = "E"
			case "E":
				nextDirection = "N"
			case "S":
				nextDirection = "W"
			case "W":
				nextDirection = "S"
			}
		}

	case '\\':
		{

			switch direction {
			case "N":
				nextDirection = "W"
			case "E":
				nextDirection = "S"
			case "S":
				nextDirection = "E"
			case "W":
				nextDirection = "N"
			}
		}

	case '-':
		{
			switch direction {
			case "N": // split
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "E", energisedPoints, previouslyVisited)
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "W", energisedPoints, previouslyVisited)
				return
			case "E":
				nextDirection = "E"
			case "S":
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "E", energisedPoints, previouslyVisited)
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "W", energisedPoints, previouslyVisited)
				return
			case "W":
				nextDirection = "W"
			}
		}
	case '|':
		{
			switch direction {
			case "N":
				nextDirection = "N"
			case "E":
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "N", energisedPoints, previouslyVisited)
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "S", energisedPoints, previouslyVisited)
				return
			case "S":
				nextDirection = "S"
			case "W":
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "N", energisedPoints, previouslyVisited)
				populateEnergisedPointMapForChar('.', rowId, colId, lines, "S", energisedPoints, previouslyVisited)
				return
			}
		}
	}

	switch nextDirection {
	case "N":
		nextRowId = rowId - 1
		nextColId = colId
	case "E":
		nextRowId = rowId
		nextColId = colId + 1
	case "S":
		nextRowId = rowId + 1
		nextColId = colId
	case "W":
		nextRowId = rowId
		nextColId = colId - 1
	}

	populateEnergisedPointMap(nextRowId, nextColId, lines, nextDirection, energisedPoints, previouslyVisited)

}
