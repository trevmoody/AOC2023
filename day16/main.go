package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"math"
)

func main() {

	lines := *util.GetFileAsLines("input")
	part1(lines)
	part2(lines)
}

type Point struct {
	row, col int
}

type PointDirection struct {
	row, col  int
	direction string
}

func part2(lines []string) {
	currentMax := 0

	// row 0
	for colId := 0; colId < len(lines[0]); colId++ {
		energisedPoints := map[Point]bool{}
		populateEnergisedPointMap(0, colId, lines, "S", energisedPoints, map[PointDirection]bool{})
		currentMax = int(math.Max(float64(currentMax), float64(len(energisedPoints))))
	}

	// row last
	for colId := 0; colId < len(lines[0]); colId++ {
		energisedPoints := map[Point]bool{}
		populateEnergisedPointMap(len(lines)-1, colId, lines, "N", energisedPoints, map[PointDirection]bool{})
		currentMax = int(math.Max(float64(currentMax), float64(len(energisedPoints))))
	}

	// col 0
	for rowId := 0; rowId < len(lines); rowId++ {
		energisedPoints := map[Point]bool{}
		populateEnergisedPointMap(rowId, 0, lines, "E", energisedPoints, map[PointDirection]bool{})
		currentMax = int(math.Max(float64(currentMax), float64(len(energisedPoints))))
	}

	// col last
	for rowId := 0; rowId < len(lines); rowId++ {
		energisedPoints := map[Point]bool{}
		populateEnergisedPointMap(rowId, len(lines[0]), lines, "W", energisedPoints, map[PointDirection]bool{})
		currentMax = int(math.Max(float64(currentMax), float64(len(energisedPoints))))
	}

	fmt.Printf("Part 2 Max Energised Tiles Count: %d\n", currentMax)
}

func part1(lines []string) {
	energisedPoints := map[Point]bool{}
	previouslyVisited := map[PointDirection]bool{}

	populateEnergisedPointMap(0, 0, lines, "E", energisedPoints, previouslyVisited)

	fmt.Printf("Part 1 Energised Tiles Count: %d\n", len(energisedPoints))
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

	var nextDirections []string
	var nextRowId int
	var nextColId int

	switch currentChar {
	case '.':
		{
			switch direction {
			case "N":
				nextDirections = append(nextDirections, "N")
			case "E":
				nextDirections = append(nextDirections, "E")
			case "S":
				nextDirections = append(nextDirections, "S")
			case "W":
				nextDirections = append(nextDirections, "W")
			}
		}

	case '/':
		{
			switch direction {
			case "N":
				nextDirections = append(nextDirections, "E")
			case "E":
				nextDirections = append(nextDirections, "N")
			case "S":
				nextDirections = append(nextDirections, "W")
			case "W":
				nextDirections = append(nextDirections, "S")
			}
		}

	case '\\':
		{

			switch direction {
			case "N":
				nextDirections = append(nextDirections, "W")
			case "E":
				nextDirections = append(nextDirections, "S")
			case "S":
				nextDirections = append(nextDirections, "E")
			case "W":
				nextDirections = append(nextDirections, "N")
			}
		}

	case '-':
		{
			switch direction {
			case "N": // split
				nextDirections = append(nextDirections, "W")
				nextDirections = append(nextDirections, "E")
			case "E":
				nextDirections = append(nextDirections, "E")
			case "S":
				nextDirections = append(nextDirections, "W")
				nextDirections = append(nextDirections, "E")
			case "W":
				nextDirections = append(nextDirections, "W")
			}
		}
	case '|':
		{
			switch direction {
			case "N":
				nextDirections = append(nextDirections, "N")
			case "E":
				nextDirections = append(nextDirections, "N")
				nextDirections = append(nextDirections, "S")
			case "S":
				nextDirections = append(nextDirections, "S")
			case "W":
				nextDirections = append(nextDirections, "N")
				nextDirections = append(nextDirections, "S")
			}
		}
	}

	for _, nextDirection := range nextDirections {
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
}
