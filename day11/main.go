package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"math"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

}

type Point struct {
	row, col int
}

func part1(lines []string) int {

	planetsList := []Point{}

	fmt.Printf("got %d lines\n", len(lines))

	emptyRowCount := 0

	for rowId, line := range lines {
		emptyRow := true
		for colId := 0; colId < len(line); colId++ {
			if line[colId] == byte('#') {
				emptyRow = false
				planetsList = append(planetsList, Point{rowId + emptyRowCount, colId})
			}
		}
		if emptyRow {
			emptyRowCount += 1000000 - 1
		}
	}

	fmt.Printf("PlantList %v\n", planetsList)
	emptyColCount := 0
	for colId := 0; colId < len(lines[0]); colId++ {
		emptyCol := true
		for rowId := 0; rowId < len(lines); rowId++ {
			if lines[rowId][colId] == byte('#') {
				emptyCol = false
			}
		}
		if emptyCol {
			fmt.Printf("Empty Col at Id %d\n", colId)
			// ok so any col > colId we increase
			for i, point := range planetsList {
				if point.col-emptyColCount > colId { // original col position
					planetsList[i] = Point{point.row, point.col + (1000000 - 1)}
				}
			}
			//how much have we already moved...
			emptyColCount += 1000000 - 1
		}
	}

	fmt.Printf("PlantList %v\n", planetsList)

	sum := 0
	for i1, point1 := range planetsList {
		for i2, point2 := range planetsList {
			if i2 <= i1 {
				continue
			} else {
				rowDistance := math.Abs(float64(point2.row - point1.row))
				colDistance := math.Abs(float64(point2.col - point1.col))

				sum += int(rowDistance)
				sum += int(colDistance)
			}

		}
	}

	fmt.Printf("TOTAL DISTANCE %d\n\n", sum)

	return 0

}
