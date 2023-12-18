package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

var re = regexp.MustCompile(`(.) (.*) \(#(\w{5})(\w)\)`)

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
		matches := re.FindStringSubmatch(line)
		direction := matches[1]
		amount, _ := strconv.Atoi(matches[2])

		length += amount

		nextPoint = move(nextPoint, direction, amount)
		points = append(points, nextPoint)
	}
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
		matches := re.FindStringSubmatch(line)
		//split := strings.Fields(line)
		amount, _ := strconv.ParseInt(matches[3], 16, 64)
		length += int(amount)

		nextPoint = move(nextPoint, directionLookup[matches[4]], int(amount))
		points = append(points, nextPoint)
	}

	area := calc(points, length)
	fmt.Printf("Part 2 Area %d\n", area)
}

func calc(points []Point, length int) int {
	// half the length for straights
	// 3/4 * 4 for the outer corners = 3
	// 3/4 + 1/4 for the other matching corners, so same as half the length, 2 becomes 1.
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

// Shoelace Formula
// https://en.wikipedia.org/wiki/Shoelace_formula
func volume(points []Point) float64 {
	sum1 := points[len(points)-1].x * points[0].y
	sum2 := points[len(points)-1].y * points[0].x

	for i := 0; i < len(points)-1; i++ {
		sum1 += points[i].x * points[i+1].y
		sum2 += points[i].y * points[i+1].x
	}

	return math.Abs(float64(sum1-sum2)) / 2
}
