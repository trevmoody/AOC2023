package main

import (
	"AOC2023/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("input"))

	fmt.Printf("Running Part 2\n")
	part2()
}

func part1(lines []string) int {
	times := util.StringsToInts(strings.Split(lines[0], ":")[1])
	distances := util.StringsToInts(strings.Split(lines[1], ":")[1])

	fmt.Printf("Got times %v\n", times)
	fmt.Printf("Got Distances %v\n", distances)

	totalProduct := 1
	for raceId := 0; raceId < len(times); raceId++ {
		wayToWinCount := 0
		for buttonHoldTime := 0; buttonHoldTime < times[raceId]; buttonHoldTime++ {
			if buttonHoldTime*(times[raceId]-buttonHoldTime) > distances[raceId] {
				wayToWinCount += 1
			}
		}

		fmt.Printf("Race :%d, got %d ways to win\n", raceId, wayToWinCount)
		totalProduct *= wayToWinCount
	}

	fmt.Printf("TotalProduct %d\n", totalProduct)
	return totalProduct
}

func part2() int {
	raceTime := 35937366
	raceRecordDistance := 212206012011044

	wayToWinCount := 0
	for buttonHoldTime := 0; buttonHoldTime < raceTime; buttonHoldTime++ {
		if buttonHoldTime*(raceTime-buttonHoldTime) > raceRecordDistance {
			wayToWinCount += 1
		}
	}

	fmt.Printf("wayToWinCount = %d\n", wayToWinCount)
	return wayToWinCount
}
