package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const SmallTestString string = ``

const TestString string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const DataFile string = "data.txt"

const HIGH_CARD string = "high card"
const ONE_PAIR string = "one pair"
const TWO_PAIR string = "two pair"
const THREE_OF_A_KIND string = "three of a kind"
const FULL_HOUSE string = "full house"
const FOUR_OF_A_KIND string = "four of a kind"
const FIVE_OF_A_KIND string = "five of a kind"

// Can I use constants here?
var TypeStrengths = map[string]int{
	"high card":       0,
	"one pair":        1,
	"two pair":        2,
	"three of a kind": 3,
	"full house":      4,
	"four of a kind":  5,
	"five of a kind":  6,
}

var CardStrengths = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type Hand struct {
	Cards string
	Bid   int
	Type  string
}

/*
	Part 1 Notes

	Create Hands (wth bids)
	Order by strength
	Calculate total winnings by multiplying bidding by rank
*/

func makeHands(input string) (hands []Hand) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hands = append(hands, makeHand(line))
	}
	return
}

func makeHand(line string) (hand Hand) {
	parts := strings.Split(line, " ")
	hand.Cards = parts[0]
	hand.Bid, _ = strconv.Atoi(parts[1])
	hand.Type = determineType(hand)
	return
}

func determineType(hand Hand) string {
	cardMap := lib.FrequencyMap(strings.Split(hand.Cards, ""))
	orderedCardCounts := []int{}
	for _, count := range cardMap {
		orderedCardCounts = append(orderedCardCounts, count)
	}
	// Sort in reverse order
	sort.Slice(orderedCardCounts, func(i, j int) bool {
		return orderedCardCounts[i] > orderedCardCounts[j]
	})

	if orderedCardCounts[0] == 1 {
		return HIGH_CARD
	} else if orderedCardCounts[0] == 2 {
		if orderedCardCounts[1] == 2 {
			return TWO_PAIR
		} else {
			return ONE_PAIR
		}
	} else if orderedCardCounts[0] == 3 {
		if orderedCardCounts[1] == 2 {
			return FULL_HOUSE
		} else {
			return THREE_OF_A_KIND
		}
	} else if orderedCardCounts[0] == 4 {
		return FOUR_OF_A_KIND
	} else if orderedCardCounts[0] == 5 {
		return FIVE_OF_A_KIND
	} else {
		panic("Unknown hand type")
	}
}

func orderHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		firstHand := hands[i]
		secondHand := hands[j]
		if TypeStrengths[firstHand.Type] < TypeStrengths[secondHand.Type] {
			return true
		} else if TypeStrengths[firstHand.Type] > TypeStrengths[secondHand.Type] {
			return false
		} else {
			return compareCards(firstHand.Cards, secondHand.Cards)
		}
	})
	return hands
}

func compareCards(firstCards, secondCards string) bool {
	for i := 0; i < len(firstCards); i++ {
		if CardStrengths[string(firstCards[i])] < CardStrengths[string(secondCards[i])] {
			return true
		} else if CardStrengths[string(firstCards[i])] > CardStrengths[string(secondCards[i])] {
			return false
		} else {
			break
		}
	}
	return false
}

func solvePart1(input string) int {
	hands := makeHands(input)
	orderedHands := orderHands(hands)
	totalWinnings := 0
	for i := 0; i < len(orderedHands); i++ {
		winnings := orderedHands[i].Bid * (i + 1)
		fmt.Println(orderedHands[i], winnings)
		totalWinnings += winnings
	}
	return totalWinnings
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(6440, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
