package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {

	lines := *util.GetFileAsLines("input")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var points []Point
	nextPoint := Point{0, 0}

	length := 0
	for _, line := range lines {
		split := strings.Fields(line)
		direction := split[0]
		amount, _ := strconv.Atoi(split[1])

		length += amount

		nextPoint = move(nextPoint, direction, amount)
		points = append(points, nextPoint)
	}

	slices.Reverse(points)

	area := calc(points, length)

	fmt.Printf("Part 1 Area %d\n", area)

}

func part2(lines []string) {

	directionLookup := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}
	var points []Point
	nextPoint := Point{0, 0}

	length := 0
	for _, line := range lines {
		//split := strings.Fields(line)
		index := strings.Index(line, "#")
		hexString := line[index+1 : index+6]
		directionString := string(line[len(line)-2])
		amount, _ := strconv.ParseInt(hexString, 16, 64)

		length += int(amount)

		nextPoint = move(nextPoint, directionLookup[directionString], int(amount))
		points = append(points, nextPoint)
	}

	slices.Reverse(points)

	area := calc(points, length)

	fmt.Printf("Part 2 Area %d\n", area)

}

func calc(points []Point, length int) int {
	return int(volume(points)) + (length-4)/2 + 3
}

func move(from Point, direction string, amount int) Point {
	switch direction {
	case "D":
		{
			return Point{from.x + amount, from.y}
		}
	case "U":
		{
			return Point{from.x - amount, from.y}
		}
	case "R":
		{
			return Point{from.x, from.y + amount}
		}
	case "L":
		{
			return Point{from.x, from.y - amount}
		}
	default:
		panic("Aaaargh")
	}
}

func volume(points []Point) float64 {
	sum1 := points[len(points)-1].x * points[0].y
	sum2 := points[len(points)-1].y * points[0].x

	for i := 0; i < len(points)-1; i++ {
		sum1 += points[i].x * points[i+1].y
		sum2 += points[i].y * points[i+1].x
	}

	return math.Abs(float64(sum1-sum2)) / 2

}
