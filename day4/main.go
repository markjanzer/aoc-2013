package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const TestString string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

const SmallTestString string = `Card 1: 10 11 12 13 | 11 12 13 14`

const DataFile string = "data.txt"

/*
  Sum scores of each card

	For each card, get the winningNumbers and the drawnNumbers
	Get the count of drawn numbers that are winning numbers
	If count > 0
		Return 2 ^ (count -1)
	else
		return 0

*/

type Card struct {
	WinningNumbers []int
	GivenNumbers   []int
}

func parseCards(input string) []Card {
	lines := strings.Split(input, "\n")
	return lib.Reduce(lines, func(agg []Card, line string) []Card {
		return append(agg, parseCard(line))
	}, []Card{})
}

func parseCard(input string) (card Card) {
	cardValues := strings.Split(input, ":")[1]
	splitValues := strings.Split(cardValues, "|")
	winningNumbersString, givenNumbersString := splitValues[0], splitValues[1]

	card.WinningNumbers = getInts(winningNumbersString)
	card.GivenNumbers = getInts(givenNumbersString)

	return
}

func getInts(input string) []int {
	return lib.Map(strings.Fields(input), func(value string) int {
		result, _ := strconv.Atoi(value)
		return result
	})
}

func sharedValues(a, b []int) (shared []int) {
	for _, aVal := range a {
		for _, bVal := range b {
			if bVal == aVal {
				shared = append(shared, bVal)
				break
			}
		}
	}
	return
}

func score(card Card) int {
	numberOfMatches := len(sharedValues(card.WinningNumbers, card.GivenNumbers))
	if numberOfMatches == 0 {
		return 0
	}
	result := intPower(2, numberOfMatches-1)
	fmt.Println(card)
	fmt.Println(sharedValues(card.WinningNumbers, card.GivenNumbers))
	fmt.Println(result)

	return intPower(2, numberOfMatches-1)
}

func intPower(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

func solve1(input string) int {
	cards := parseCards(input)

	return lib.Reduce(cards, func(agg int, card Card) int {
		return agg + score(card)
	}, 0)
}

// func solve2(input string) int {
// 	return 0
// }

func main() {
	lib.AssertEqual(4, solve1(SmallTestString))
	lib.AssertEqual(13, solve1(TestString))

	// lib.AssertEqual(467835, solve2(TestString))

	// lib.AssertEqual(4361, solve1(TestString))
	// lib.AssertEqual(467835, solve2(TestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solve1(dataString)
	// result2 := solve2(dataString)

	fmt.Println(result1)
	// fmt.Println(result2)
}
