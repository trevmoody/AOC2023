package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

	lineCount := len(lines)
	//for i, line := range lines {
	//	fmt.Printf("line %d: %s\n", i, line)
	//}
	//
	//fmt.Printf("line count: %d", lineCount)

	sumRatio := 0

	for i := 0; i < lineCount; i++ {
		if i == 0 {
			sumRatio = sumRatio + processLine("", lines[i], lines[i+1]) //first line

		} else if i == lineCount-1 {
			sumRatio = sumRatio + processLine(lines[i-1], lines[i], "") //last line

		} else {
			sumRatio = sumRatio + processLine(lines[i-1], lines[i], lines[i+1])

		}
	}

	fmt.Printf("sum of ratios  is %d\n", sumRatio)

}

func processLine(prev string, current string, next string) int {
	fmt.Printf("PREV ROW::%s\n", prev)
	fmt.Printf("CURR ROW::%s\n", current)
	fmt.Printf("NEXT ROW::%s\n", next) //ok we have a symbol

	sumGearRatio := 0
	currentRunes := []rune(current)
	prevRunes := []rune(prev)
	nextRunes := []rune(next)

	var partNumber int
	var found bool

	//iterate down current line, if we find a symbol, need to cheeck prev, next, and above.
	for i := 0; i < len(currentRunes); i++ {
		char := currentRunes[i]
		if char == '*' {
			fmt.Printf("founds char %s at position %d\n", string(char), i)
			partNumbers := make([]int, 0)

			//check prev row
			if len(prev) != 0 {
				//ok is there a number in i-day1, i , i+day1 then if so read whole number.

				// check i, if not a digit, we can check forwards and backwards
				if !unicode.IsDigit(prevRunes[i]) {
					//check current row
					partNumber, found = checkBackwards(prevRunes, i)
					fmt.Printf("PREV ROW checkbackwards: %t %d\n ", found, partNumber)
					if found {
						partNumbers = append(partNumbers, partNumber)
					}

					partNumber, found = checkForwards(prevRunes, i)
					fmt.Printf("PREV ROW checkforwards: %t %d\n ", found, partNumber)
					if found {
						partNumbers = append(partNumbers, partNumber)
					}
				} else {

					//it is a number so roll back to start of number, then check forwards.
					if unicode.IsDigit(prevRunes[i-1]) {
						// we now are atleast a day2 digit number
						if unicode.IsDigit(prevRunes[i-2]) {
							partNumber, found = checkForwards(prevRunes, i-3)
							fmt.Printf("PREV ROW checkforwards: %t %d\n ", found, partNumber)
							if found {
								partNumbers = append(partNumbers, partNumber)
							}

						} else {
							partNumber, found = checkForwards(prevRunes, i-2)
							fmt.Printf("PREV ROW checkforwards: %t %d\n ", found, partNumber)
							if found {
								partNumbers = append(partNumbers, partNumber)
							}
						}

					} else {
						partNumber, found = checkForwards(prevRunes, i-1)
						fmt.Printf("PREV ROW checkforwards: %t %d\n ", found, partNumber)
						if found {
							partNumbers = append(partNumbers, partNumber)
						}
					}
				}
			}

			//check current row
			partNumber, found = checkBackwards(currentRunes, i)
			fmt.Printf("CURR ROW checkbackwards: %t %d\n ", found, partNumber)
			if found {
				partNumbers = append(partNumbers, partNumber)
			}

			partNumber, found = checkForwards(currentRunes, i)
			fmt.Printf("CURR ROW checkforwards: %t %d\n ", found, partNumber)
			if found {
				partNumbers = append(partNumbers, partNumber)
			}

			//check next

			if len(next) != 0 {
				//ok is there a number in i-day1, i , i+day1 then if so read whole number.

				// check i, if not a digit, we can check forwards and backwards
				if !unicode.IsDigit(nextRunes[i]) {
					//check current row
					partNumber, found = checkBackwards(nextRunes, i)
					fmt.Printf("NEXT ROW checkbackwards: %t %d\n ", found, partNumber)
					if found {
						partNumbers = append(partNumbers, partNumber)
					}

					partNumber, found = checkForwards(nextRunes, i)
					fmt.Printf("NEXT ROW checkforwards: %t %d\n ", found, partNumber)
					if found {
						partNumbers = append(partNumbers, partNumber)
					}
				} else {

					//it is a number so roll back to start of number, then check forwards.
					if unicode.IsDigit(nextRunes[i-1]) {
						// we now are atleast a day2 digit number
						if unicode.IsDigit(nextRunes[i-2]) {
							partNumber, found = checkForwards(nextRunes, i-3)
							fmt.Printf("NEXT ROW checkforwards: %t %d\n ", found, partNumber)
							if found {
								partNumbers = append(partNumbers, partNumber)
							}

						} else {
							partNumber, found = checkForwards(nextRunes, i-2)
							fmt.Printf("NEXT ROW checkforwards: %t %d\n ", found, partNumber)
							if found {
								partNumbers = append(partNumbers, partNumber)
							}
						}

					} else {
						partNumber, found = checkForwards(nextRunes, i-1)
						fmt.Printf("NEXT ROW checkforwards: %t %d\n ", found, partNumber)
						if found {
							partNumbers = append(partNumbers, partNumber)
						}
					}
				}
			}

			if len(partNumbers) == 2 {

				gearRatio := partNumbers[0] * partNumbers[1]
				sumGearRatio = sumGearRatio + gearRatio
			}
		}

	}

	fmt.Printf("found gear ratios sum is  %d\n", sumGearRatio)
	return sumGearRatio
}

func checkForwards(line []rune, position int) (int, bool) {
	if unicode.IsDigit(line[position+1]) {
		if unicode.IsDigit(line[position+2]) {
			if unicode.IsDigit(line[position+3]) {
				// day3  digit number
				n := fmt.Sprintf("%s%s%s", string(line[position+1]), string(line[position+2]), string(line[position+3]))
				number, _ := strconv.Atoi(n)
				return number, true

			} else {
				//day2 digit number
				n := fmt.Sprintf("%s%s", string(line[position+1]), string(line[position+2]))
				number, _ := strconv.Atoi(n)
				return number, true
			}

		} else {
			// day1 digit number
			n := fmt.Sprintf("%s", string(line[position+1]))
			number, _ := strconv.Atoi(n)
			return number, true

		}
	} else {
		// no number
		return -1, false
	}
}

func checkBackwards(line []rune, position int) (int, bool) {
	if unicode.IsDigit(line[position-1]) {
		if unicode.IsDigit(line[position-2]) {
			if unicode.IsDigit(line[position-3]) {
				// day3  digit number
				n := fmt.Sprintf("%s%s%s", string(line[position-3]), string(line[position-2]), string(line[position-1]))
				number, _ := strconv.Atoi(n)
				return number, true

			} else {
				//day2 digit number
				n := fmt.Sprintf("%s%s", string(line[position-2]), string(line[position-1]))
				number, _ := strconv.Atoi(n)
				return number, true
			}

		} else {
			// day1 digit number
			n := fmt.Sprintf("%s", string(line[position-1]))
			number, _ := strconv.Atoi(n)
			return number, true

		}
	} else {
		// no number
		return -1, false
	}
}
