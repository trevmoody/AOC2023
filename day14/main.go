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

func part2(lines []string) {

	length := len(lines)
	width := len(lines[0])

	rocks := map[Point]bool{}
	mirrors := map[Point]bool{}
	for rowId := 0; rowId < length; rowId++ {
		for colId := 0; colId < width; colId++ {
			point := Point{row: rowId, col: colId}

			char := lines[rowId][colId]
			switch char {
			case 'O':
				mirrors[point] = true
			case '#':
				rocks[point] = true
			default:

			}
		}
	}

	for i := 0; i < 1000000000; i++ {
		rollNorth(mirrors, rocks, length, width)
		rollWest(mirrors, rocks, length, width)
		rollSouth(mirrors, rocks, length, width)
		rollEast(mirrors, rocks, length, width)

		if i%13 == 0 {
			fmt.Printf("i %d, weight %d\n", i, getWeight(mirrors, length, width))
		}

	}

	weight := getWeight(mirrors, length, width)
	fmt.Printf("Total is %d\n", weight)

}

func printout(mirrors map[Point]bool, rocks map[Point]bool, length int, width int) {
	for rowId := 0; rowId < length; rowId++ {
		for colId := 0; colId < width; colId++ {
			point := Point{row: rowId, col: colId}
			okMirror, _ := mirrors[point]
			okRock, _ := rocks[point]

			if okMirror {
				fmt.Print("O")
			} else if okRock {
				fmt.Printf("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	fmt.Print("\n")
}

func getWeight(mirrors map[Point]bool, length int, width int) int {
	weight := 0
	for colId := 0; colId < width; colId++ {

		for rowId := 0; rowId < length; rowId++ {
			ok, _ := mirrors[Point{rowId, colId}]
			if ok {
				weight += length - rowId
			}
		}
	}
	return weight
}

func rollEast(mirrors map[Point]bool, rocks map[Point]bool, length int, width int) {
	for rowId := 0; rowId < length; rowId++ {
		nextFreeCol := -1
		for colId := width - 1; colId >= 0; colId-- {
			currentPoint := Point{row: rowId, col: colId}
			okMirror, _ := mirrors[currentPoint]
			okRock, _ := rocks[currentPoint]

			if okMirror && nextFreeCol != -1 {
				// move mirror to free row
				delete(mirrors, currentPoint)
				mirrors[Point{row: rowId, col: nextFreeCol}] = true
				nextFreeCol = nextFreeCol - 1
			} else if okMirror {
				// do nothing
				continue
			} else if okRock {
				nextFreeCol = -1

			} else { // set free col if not set
				if nextFreeCol == -1 {
					nextFreeCol = colId
				}
			}
		}
	}
}

func rollSouth(mirrors map[Point]bool, rocks map[Point]bool, length int, width int) {
	for colId := 0; colId < width; colId++ {
		nextFreeRow := -1
		for rowId := length - 1; rowId >= 0; rowId-- {

			currentPoint := Point{row: rowId, col: colId}
			okMirror, _ := mirrors[currentPoint]
			okRock, _ := rocks[currentPoint]
			if okMirror && nextFreeRow != -1 {
				// move mirror to free row
				delete(mirrors, currentPoint)
				mirrors[Point{row: nextFreeRow, col: colId}] = true
				nextFreeRow = nextFreeRow - 1
			} else if okMirror {
				// do nothing
				continue
			} else if okRock {
				nextFreeRow = -1

			} else {
				// set free row if not set
				if nextFreeRow == -1 {
					nextFreeRow = rowId
				}
			}
		}
	}
}

func rollWest(mirrors map[Point]bool, rocks map[Point]bool, length int, width int) {
	for rowId := 0; rowId < length; rowId++ {
		nextFreeCol := -1
		for colId := 0; colId < width; colId++ {
			currentPoint := Point{row: rowId, col: colId}
			okMirror, _ := mirrors[currentPoint]
			okRock, _ := rocks[currentPoint]

			if okMirror && nextFreeCol != -1 {
				// move mirror to free row
				delete(mirrors, currentPoint)
				mirrors[Point{row: rowId, col: nextFreeCol}] = true
				nextFreeCol = nextFreeCol + 1
			} else if okMirror {
				// do nothing
				continue
			} else if okRock {
				nextFreeCol = -1

			} else { // set free col if not set
				if nextFreeCol == -1 {
					nextFreeCol = colId
				}
			}
		}
	}
}

func rollNorth(mirrors map[Point]bool, rocks map[Point]bool, length int, width int) {

	for colId := 0; colId < width; colId++ {
		nextFreeRow := -1
		for rowId := 0; rowId < length; rowId++ {

			currentPoint := Point{row: rowId, col: colId}
			okMirror, _ := mirrors[currentPoint]
			okRock, _ := rocks[currentPoint]
			if okMirror && nextFreeRow != -1 {
				// move mirror to free row
				delete(mirrors, currentPoint)
				mirrors[Point{row: nextFreeRow, col: colId}] = true

				nextFreeRow = nextFreeRow + 1

			} else if okMirror {
				// do nothing
				continue
			} else if okRock {
				nextFreeRow = -1

			} else {
				// set free row if not set
				if nextFreeRow == -1 {
					nextFreeRow = rowId
				}
			}
		}
	}
}

func part1(lines []string) {

	totalWeight := 0
	for colId := 0; colId < len(lines[0]); colId++ {
		startWeight := len(lines)
		nextWeight := startWeight
		weight := 0
		for rowId := 0; rowId < len(lines); rowId++ {
			switch lines[rowId][colId] {
			case 'O':
				{
					weight += nextWeight
					nextWeight -= 1
				}
			case '#':
				nextWeight = len(lines) - rowId - 1
			case '.':
				continue
			}
		}

		fmt.Printf("Col %d Weight %d\n", colId, weight)
		totalWeight += weight
	}

	fmt.Printf("Part Total Weight %d\n", totalWeight)
}
