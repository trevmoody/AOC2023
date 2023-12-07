package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

	processFile(file)

}

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		fmt.Printf("line %d: %s\n", i, line)
	}

	fmt.Printf("line count: %d\n", len(lines))

	processLines(lines)

}

type void struct{}

func processLines(lines []string) {
	cardCountById := map[int]int{}

	for i := 1; i <= len(lines); i++ {
		cardCountById[i] = 1
	}

	cardIdPattern := regexp.MustCompile(`^Card\s+(\d+):(.+)`)
	for _, line := range lines {
		match := cardIdPattern.FindAllStringSubmatch(line, -1)
		id, _ := strconv.Atoi(match[0][1])
		restSplit := strings.Split(match[0][2], "|")

		var winningNumbers []int
		for _, numString := range strings.Fields(restSplit[0]) {
			num, _ := strconv.Atoi(numString)
			winningNumbers = append(winningNumbers, num)
		}

		var allNumbers = map[int]void{}
		for _, numString := range strings.Fields(restSplit[1]) {
			num, _ := strconv.Atoi(numString)
			allNumbers[num] = void{}
		}

		//fmt.Printf("Id: %d, wining Numbers %v, all Numbers %v\n", id, winningNumbers, allNumbers)

		winCount := 0
		for _, number := range winningNumbers {
			if _, ok := allNumbers[number]; ok {
				winCount = winCount + 1
			}
		}

		numberOfCopies, _ := cardCountById[id]

		fmt.Printf("CardId %d wincount: %d, totalCopies %d\n", id, winCount, numberOfCopies)

		for i := id + 1; i <= id+winCount; i++ {
			if currentCopyCount, ok := cardCountById[i]; ok {
				cardCountById[i] = currentCopyCount + numberOfCopies
			} else {
				panic("should not happen as total copies initialised to day1")
			}
		}
	}

	totalCards := 0
	for _, count := range cardCountById {
		totalCards = totalCards + count
	}

	fmt.Printf("TOTAL CARDS : %d\n", totalCards)

}
