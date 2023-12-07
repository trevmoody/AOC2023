package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var NumberToWord map[string]string
var ReversedWordToNumber map[string]string
var WordToNumber map[string]string

func init() {
	NumberToWord = map[string]string{
		"day1": "one",
		"day2": "two",
		"day3": "three",
		"day4": "four",
		"5":    "five",
		"6":    "six",
		"7":    "seven",
		"8":    "eight",
		"9":    "nine",
	}

	ReversedWordToNumber = make(map[string]string)
	for number, word := range NumberToWord {
		ReversedWordToNumber[reverse(word)] = number
	}

	WordToNumber = make(map[string]string)
	for number, word := range NumberToWord {
		WordToNumber[word] = number
	}
}

// part1regexp := regexp.MustCompile(`\d`)
var part2regexp = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
var part2regexpreversed = regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)

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
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstMatch, lastMatch, number := processLine(line)

		sum = sum + number

		fmt.Printf("input: %s, first digit %s, last digit %s , number %d, sum=%d\n", line, firstMatch, lastMatch, number, sum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Total: %d\n", sum)
}

func processLine(line string) (string, string, int) {
	firstMatch := part2regexp.FindString(line)
	lastMatch := part2regexpreversed.FindString(reverse(line))
	convertedFirstMatch := convert(firstMatch, WordToNumber)
	convertedLastMatch := convert(lastMatch, ReversedWordToNumber)

	number, err := strconv.Atoi(fmt.Sprintf("%s%s", convertedFirstMatch, convertedLastMatch))

	if err != nil {
		fmt.Println("ERROR", err)
	}

	return convertedFirstMatch, convertedLastMatch, number
}

// ripped from internet somewhere
func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func convert(s string, m map[string]string) string {
	var number, ok = m[s]
	if ok {
		return number
	}
	return s
}
