package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part1(*util.GetFileAsLines("testinput"))

}

var cardValues map[string]int = map[string]int{"A": 100,
	"K": 99,
	"Q": 98,
	"J": 97,
	"T": 96,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

const ( // iota is reset to 0
	HighCard     = iota // c0 == 0
	OnePair      = iota // c1 == 1
	TwoPair      = iota // c2 == 2
	ThreeOfAKind = iota // c0 == 0
	FullHouse    = iota // c1 == 1
	FourOfAKind  = iota
	FiveOfAKind  = iota
)

type hand struct {
	cards    string
	handType int
	bid      int
}

var handRegex = regexp.MustCompile(`^([A-Za-z0-9]{5})\s?(\d+)$`)

func newHand(line string) hand {

	match := handRegex.FindAllStringSubmatch(line, -1)

	cardString := match[0][1]
	bid, _ := strconv.Atoi(match[0][2])

	cardCountMap := map[string]int{}
	for _, card := range []rune(cardString) {
		currentCount, ok := cardCountMap[string(card)]
		if ok {
			cardCountMap[string(card)] = currentCount + 1
		} else {
			cardCountMap[string(card)] = 1
		}
	}

	mapLength := len(cardCountMap)
	values := make([]int, mapLength)
	i := 0
	for _, v := range cardCountMap {
		values[i] = v
		i++
	}

	sort.Ints(values)

	var handType int

	if mapLength == 1 {
		handType = FiveOfAKind
	} else if mapLength == 2 {
		if values[1] == 4 {
			handType = FourOfAKind
		} else {
			handType = FullHouse
		}
	} else if mapLength == 3 {
		if values[2] == 3 {
			handType = ThreeOfAKind
		} else {
			handType = TwoPair
		}
	} else if mapLength == 4 {
		handType = OnePair
	} else {
		handType = HighCard
	}

	h := hand{
		cards:    cardString,
		handType: handType,
		bid:      bid,
	}

	return h
}

func part1(lines []string) int {

	fmt.Printf("UNSorted Lines: %v\n", lines)
	var hands []hand
	for _, line := range lines {
		hands = append(hands, newHand(line))
	}

	fmt.Printf("UN Sorted: %v\n", hands)
	sort.Slice(hands, func(i, j int) bool {
		//less function
		return less(hands[i], hands[j])

	})

	fmt.Printf("Sorted: %v\n", hands)
	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	fmt.Printf("Winnings: %v\n", winnings)
	return winnings
}

func less(h1 hand, h2 hand) bool {

	fmt.Printf("Comparing %v vs %v\n", h1, h2)
	if h1.handType == h2.handType {
		strings1 := strings.Split(h1.cards, "")
		strings2 := strings.Split(h2.cards, "")

		for k := 0; k < len(strings1); k++ {
			s2 := strings2[k]
			val1 := cardValues[strings1[k]]
			val2 := cardValues[s2]
			if val1 == val2 {
				continue
			}
			retVal := val1 < val2
			fmt.Printf("Returning %v < %v : %t\n", val1, val2, retVal)
			return retVal
		}

	} else {
		retVal := h1.handType < h2.handType
		fmt.Printf("Returning %v < %v : %t\n", h1.handType, h2.handType, retVal)
		return retVal
	}
	fmt.Printf("Returning default TRUE\n")
	return true
}
