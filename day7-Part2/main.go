package main

import (
	"fmt"
	"github.com/trevmoody/aoc23/util"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Running Part 1\n")
	part2(*util.GetFileAsLines("input"))

}

var cardToValue = map[string]int{
	"A": 100,
	"K": 99,
	"Q": 98,
	"T": 96,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
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

const joker = "J"

type hand struct {
	cards    string
	handType int
	bid      int
}
type handSlice []hand

func (x handSlice) Len() int           { return len(x) }
func (x handSlice) Less(i, j int) bool { return less(x[i], x[j]) }
func (x handSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

var handRegex = regexp.MustCompile(`^([A-Za-z0-9]{5})\s?(\d+)$`)

func newHand(cardString string, bid int) hand {
	return hand{
		cards:    cardString,
		handType: handType(cardString),
		bid:      bid,
	}

}
func newHandFromLine(line string) hand {
	match := handRegex.FindAllStringSubmatch(line, -1)
	bid, _ := strconv.Atoi(match[0][2])

	return newHand(match[0][1], bid)
}

func handType(cardString string) int {

	cardCountMap := map[string]int{}
	jokerCount := 0

	for _, card := range strings.Split(cardString, "") {
		if card == joker {
			jokerCount++
			continue
		}
		currentCount, ok := cardCountMap[card]
		if ok {
			cardCountMap[card] = currentCount + 1
		} else {
			cardCountMap[card] = 1
		}
	}

	// we need to add any J's to the most occuring card
	var mostFrequentCard string
	prevCount := 0
	for key, value := range cardCountMap {
		if value > prevCount {
			mostFrequentCard = key
			prevCount = value
		}
	}
	cardCountMap[mostFrequentCard] = cardCountMap[mostFrequentCard] + jokerCount

	cardCountLength := len(cardCountMap)
	cardCounts := make([]int, cardCountLength)
	i := 0
	for _, v := range cardCountMap {
		cardCounts[i] = v
		i++
	}

	//sort them so we can work it out how many cards in most frequent
	sort.Sort(sort.Reverse(sort.IntSlice(cardCounts)))

	var handType int

	switch {
	case cardCountLength == 1:
		handType = FiveOfAKind
	case cardCountLength == 2 && cardCounts[0] == 4:
		handType = FourOfAKind
	case cardCountLength == 2:
		handType = FullHouse
	case cardCountLength == 3 && cardCounts[0] == 3:
		handType = ThreeOfAKind
	case cardCountLength == 3:
		handType = TwoPair
	case cardCountLength == 4:
		handType = OnePair
	case cardCountLength == 5:
		handType = HighCard

	}
	return handType
}

func part2(lines []string) int {

	var hands []hand
	for _, line := range lines {
		hands = append(hands, newHandFromLine(line))
	}

	sort.Sort(handSlice(hands))

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	fmt.Printf("Winnings: %v\n", winnings)
	return winnings
}

func less(h1 hand, h2 hand) bool {

	//fmt.Printf("Comparing %v vs %v\n", h1, h2)
	if h1.handType == h2.handType {
		strings1 := strings.Split(h1.cards, "")
		strings2 := strings.Split(h2.cards, "")

		for k := 0; k < len(strings1); k++ {
			val1 := cardToValue[strings1[k]]
			val2 := cardToValue[strings2[k]]
			if val1 == val2 {
				continue
			}
			retVal := val1 < val2
			//	fmt.Printf("Returning %v < %v : %t\n", val1, val2, retVal)
			return retVal
		}

	} else {
		retVal := h1.handType < h2.handType
		//fmt.Printf("Returning %v < %v : %t\n", h1.handType, h2.handType, retVal)
		return retVal
	}
	//fmt.Printf("Returning default TRUE\n")
	return true
}
