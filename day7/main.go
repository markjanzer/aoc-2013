package main

import (
	"advent-of-code-2023/lib"
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
	Cards []string
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
		hand := Hand{}
		parts := strings.Split(line, " ")
		hand.Cards = strings.Split(parts[0], "")
		hand.Bid, _ = strconv.Atoi(parts[1])

		orderedCardCounts := orderedCardCounts(hand.Cards)
		hand.Type = determineType(orderedCardCounts)
		hands = append(hands, hand)
	}
	return
}

func determineType(orderedCardCounts []int) string {
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

func orderedCardCounts(cards []string) []int {
	if len(cards) == 0 {
		return []int{0}
	}

	cardMap := lib.FrequencyMap(cards)
	orderedCardCounts := []int{}
	for _, count := range cardMap {
		orderedCardCounts = append(orderedCardCounts, count)
	}
	// Sort in reverse order
	sort.Slice(orderedCardCounts, func(i, j int) bool {
		return orderedCardCounts[i] > orderedCardCounts[j]
	})

	return orderedCardCounts
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

func compareCards(firstCards, secondCards []string) bool {
	for i := 0; i < len(firstCards); i++ {
		if CardStrengths[firstCards[i]] < CardStrengths[secondCards[i]] {
			return true
		} else if CardStrengths[string(firstCards[i])] > CardStrengths[string(secondCards[i])] {
			return false
		}
	}
	return true
}

func solvePart1(input string) int {
	hands := makeHands(input)
	orderedHands := orderHands(hands)
	totalWinnings := 0
	for i := 0; i < len(orderedHands); i++ {
		winnings := orderedHands[i].Bid * (i + 1)
		totalWinnings += winnings
	}
	return totalWinnings
}

/*
	Part 2 Notes

	Jacks are now Jokers.

	I'll need new CardStrengths that set J to -1
	I'll need to revamp my frequency map code
	Pretty much, J will be ignored until the end, when the
	count of the Jacks will be added to the highest frequency value
*/

var CardStrengthsPart2 = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": -1,
	"Q": 10,
	"K": 11,
	"A": 12,
}

// I'd like for this to share a bit more logic
func makeHandsPart2(input string) (hands []Hand) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hand := Hand{}
		parts := strings.Split(line, " ")
		hand.Cards = strings.Split(parts[0], "")
		hand.Bid, _ = strconv.Atoi(parts[1])

		orderedCardCounts := orderedCardCountsPart2(hand.Cards)
		hand.Type = determineType(orderedCardCounts)
		hands = append(hands, hand)
	}
	return
}

func orderedCardCountsPart2(cards []string) []int {
	jokerCount := 0
	cardsWithoutJokers := []string{}
	for _, card := range cards {
		if card == "J" {
			jokerCount++
		} else {
			cardsWithoutJokers = append(cardsWithoutJokers, card)
		}
	}

	orderedCardCounts := orderedCardCounts(cardsWithoutJokers)
	firstValue := orderedCardCounts[0]
	firstValue += jokerCount
	orderedCardCounts[0] = firstValue
	return orderedCardCounts
}

// Again, A LOT of duplication
func orderHandsPart2(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		firstHand := hands[i]
		secondHand := hands[j]
		if TypeStrengths[firstHand.Type] < TypeStrengths[secondHand.Type] {
			return true
		} else if TypeStrengths[firstHand.Type] > TypeStrengths[secondHand.Type] {
			return false
		} else {
			return compareCardsPart2(firstHand.Cards, secondHand.Cards)
		}
	})
	return hands
}

func compareCardsPart2(firstCards, secondCards []string) bool {
	for i := 0; i < len(firstCards); i++ {
		if CardStrengthsPart2[firstCards[i]] < CardStrengthsPart2[secondCards[i]] {
			return true
		} else if CardStrengthsPart2[string(firstCards[i])] > CardStrengthsPart2[string(secondCards[i])] {
			return false
		}
	}
	return true
}

func solvePart2(input string) int {
	hands := makeHandsPart2(input)
	orderedHands := orderHandsPart2(hands)
	totalWinnings := 0
	for i := 0; i < len(orderedHands); i++ {
		winnings := orderedHands[i].Bid * (i + 1)
		totalWinnings += winnings
	}
	return totalWinnings
}

func main() {
	lib.AssertEqual(6440, solvePart1(TestString))
	lib.AssertEqual(5905, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
