package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var totals = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func init() {

}

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	processFile2(file)
}

func processFile2(file *os.File) {
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		id, minColCounts := processLine2(line)

		power := 1
		for _, count := range minColCounts {
			power = power * count
		}

		sum = sum + power

		fmt.Printf("input=> %s :: GameID=%d, Power=%d\n", line, id, power)

	}

	fmt.Printf("Total: %d\n", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func processLine2(line string) (int, map[string]int) {

	minColCounts := make(map[string]int)

	split := strings.Split(line, ":")

	var gameIDString = split[0]

	pattern := regexp.MustCompile(`(\d+)$`)
	match := pattern.FindString(gameIDString)
	gameId, _ := strconv.Atoi(match)

	trys := make([]map[string]int, 0)
	var gamesDetails = split[1]
	// Split input into individual games
	tryStrings := strings.Split(gamesDetails, ";")

	trypattern := regexp.MustCompile(`(\d+)\s*(\w+)`)

	for _, tryString := range tryStrings {
		try := make(map[string]int)

		// Extract color and count using regex
		matches := trypattern.FindAllStringSubmatch(tryString, -1)

		// Populate the map for each color
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			try[color] = count
		}

		// Add the game map to the list
		trys = append(trys, try)
	}

	for _, try := range trys {
		for colour, tryCount := range try {

			// check the colour even exitst
			minColourCount, valid := minColCounts[colour]

			if valid == false {
				minColCounts[colour] = tryCount
			} else if minColourCount < tryCount {
				minColCounts[colour] = tryCount
			}

		}
	}

	return gameId, minColCounts
}

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		id, possible := processLine(line)
		if possible {
			sum = sum + id
		}

		fmt.Printf("input=> %s :: GameID=%d, Possible=%t\n", line, id, possible)

	}

	fmt.Printf("Total: %d\n", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func processLine(line string) (int, bool) {

	split := strings.Split(line, ":")

	var gameIDString = split[0]

	pattern := regexp.MustCompile(`(\d+)$`)
	match := pattern.FindString(gameIDString)
	gameId, _ := strconv.Atoi(match)

	trys := make([]map[string]int, 0)
	var gamesDetails = split[1]
	// Split input into individual games
	tryStrings := strings.Split(gamesDetails, ";")

	trypattern := regexp.MustCompile(`(\d+)\s*(\w+)`)

	for _, tryString := range tryStrings {
		try := make(map[string]int)

		// Extract color and count using regex
		matches := trypattern.FindAllStringSubmatch(tryString, -1)

		// Populate the map for each color
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			try[color] = count
		}

		// Add the game map to the list
		trys = append(trys, try)
	}

	for _, try := range trys {
		for colour, tryCount := range try {

			// check the colour even exitst
			totalColourCount, valid := totals[colour]

			if valid == false {
				return gameId, false
			}
			if totalColourCount < tryCount {
				return gameId, false
			}

		}
	}

	return gameId, true
}

type game struct {
	id   int
	trys []map[string]int
}

func parse(line string) game {
	pattern := regexp.MustCompile(`^Game (\d+):(.+)`)
	pattern2 := regexp.MustCompile(`(\d+)\s*([a-zA-Z]+)`)

	matches := pattern.FindAllStringSubmatch(line, -1)

	// Populate the map for each color
	var id int
	for _, match := range matches {
		id, _ = strconv.Atoi(match[1])
		tryStrings := strings.Split(match[2], ";")
		for _, tryString := range tryStrings {
			//need to convert this to a map
			matches := pattern2.FindAllString(tryString, -1)

			// Create a map to store the results
			resultMap := make(map[string]int)

			// Iterate through matches and populate the map
			for _, match := range matches {
				resultMap[match] = 1
			}
		}

	}

	g := game{id: id, trys: nil}
	return g
}
