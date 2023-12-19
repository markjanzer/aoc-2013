package main

import (
	"advent-of-code-2023/lib"
	"fmt"
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

var HIGH_CARD string = "high card"
var ONE_PAIR string = "one pair"
var TWO_PAIR string = "two pair"
var THREE_OF_A_KIND string = "three of a kind"
var FULL_HOUSE string = "full house"
var FOUR_OF_A_KIND string = "four of a kind"
var FIVE_OF_A_KIND string = "five of a kind"

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
	hand.Type = "dunno"
	// hand.Type = determineType(hand)
	return
}

func solvePart1(input string) int {
	hands := makeHands(input)
	fmt.Println(hands)
	return 0
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
