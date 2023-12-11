package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"strings"
)

func main() {

	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

	//fmt.Printf("Running Part 2\n")
	//part2(*util.GetFileAsLines("input"))

}

func part1(lines []string) int {
	fmt.Printf("got %d lines\n", len(lines))

	var grid [][]string

	var startHorizontal int
	var startVertical int

	for i, line := range lines {

		parsedLine := strings.Split(line, "")
		grid = append(grid, parsedLine)

		index := strings.Index(line, "S")
		if index != -1 {
			startVertical = i
			startHorizontal = index
		}
	}

	// ok so we start from each of the 4 adjacent, then follow
	// but we know for now its up so....

	//fmt.Printf("StartX %d, Start Y %d, Grid: %v\n", startHorizontal, startVertical, grid)

	steps := moveAndIncrement(grid, 1, startHorizontal, startVertical+1, "S")

	//fmt.Printf("startHorizontal %d, startVertical %d, Grid: %v\n", startHorizontal, startVertical, grid, steps)

	fmt.Printf("MAX Dist := %d\n", steps/2)
	return (steps - 1) / 2

}

func nextCoords(currentHorizontal int, currentVertical int, directionToTravel string) (int, int) {
	switch directionToTravel {
	case "N":
		return currentHorizontal, currentVertical - 1
	case "S":
		return currentHorizontal, currentVertical + 1
	case "E":
		return currentHorizontal + 1, currentVertical
	case "W":
		return currentHorizontal - 1, currentVertical
	}
	panic("cant get cords")
	return 0, 0
}

func moveAndIncrement(grid [][]string, currentCount int, horizontal int, vertical int, directionTraveled string) int {

	// 0,0 top left
	// first dimension is vertical, 2nd horizontal

	currentCount++
	currentPipe := grid[vertical][horizontal]
	fmt.Printf("steps = %d, horizontal:%d, vertical: %d, pipe %s\n", currentCount, horizontal, vertical, currentPipe)

	var nextHorizontal int
	var nextVertical int
	var nextDirection string
	switch {
	case currentPipe == "|" && directionTraveled == "N":
		{
			nextDirection = "N"
		}
	case currentPipe == "|" && directionTraveled == "S":
		{
			nextDirection = "S"
		}
	case currentPipe == "-" && directionTraveled == "E":
		{
			nextDirection = "E"
		}
	case currentPipe == "-" && directionTraveled == "W":
		{
			nextDirection = "W"
		}
	case currentPipe == "L" && directionTraveled == "S":
		{
			nextDirection = "E"
		}
	case currentPipe == "L" && directionTraveled == "W":
		{
			nextDirection = "N"
		}
	case currentPipe == "J" && directionTraveled == "S":
		{
			nextDirection = "W"
		}
	case currentPipe == "J" && directionTraveled == "E":
		{
			nextDirection = "N"
		}

	case currentPipe == "7" && directionTraveled == "N":
		{
			nextDirection = "W"
		}
	case currentPipe == "7" && directionTraveled == "E":
		{
			nextDirection = "S"
		}

	case currentPipe == "F" && directionTraveled == "N":
		{
			nextDirection = "E"
		}
	case currentPipe == "F" && directionTraveled == "W":
		{
			nextDirection = "S"
		}

	case currentPipe == "S":
		{
			return currentCount
		}

	default:
		panic("aaargh somehting wrong")
	}

	nextHorizontal, nextVertical = nextCoords(horizontal, vertical, nextDirection)
	return moveAndIncrement(grid, currentCount, nextHorizontal, nextVertical, nextDirection)

}
