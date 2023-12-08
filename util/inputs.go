package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFileAsLines(fileName string) *[]string {

	currentDir, _ := os.Getwd()

	fmt.Printf("Current DIR: %s\n", currentDir)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic("Error opening file")
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		fmt.Printf("line %d: %s\n", i, line)
	}

	fmt.Printf("line count: %d\n", len(lines))

	return &lines
}

func StringsToInts(line string) []int {
	fields := strings.Fields(line)
	retList := []int{}
	for _, f := range fields {
		i, _ := strconv.Atoi(strings.TrimSpace(f))
		retList = append(retList, i)
	}
	return retList
}
