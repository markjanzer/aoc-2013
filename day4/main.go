package main

import (
	"advent-of-code-2023/lib"
	"math"
	"regexp"
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

type Card struct {
	Number          int
	WinningNumbers  []int
	GivenNumbers    []int
	NumberOfMatches int
	Score1          int
	Score2          int
}

func parseCard(input string) (card Card) {
	splitCard := strings.Split(input, ":")
	titleArea, cardValues := splitCard[0], splitCard[1]

	re := regexp.MustCompile(`Card\s+(\d+)`)
	matches := re.FindStringSubmatch(titleArea)
	card.Number, _ = strconv.Atoi(matches[1])

	splitValues := strings.Split(cardValues, "|")
	winningNumbersString, givenNumbersString := splitValues[0], splitValues[1]

	card.WinningNumbers = lib.IntsFromString(winningNumbersString)
	card.GivenNumbers = lib.IntsFromString(givenNumbersString)

	card.NumberOfMatches = len(sharedValues(card.WinningNumbers, card.GivenNumbers))
	card.Score1 = score1(card.NumberOfMatches)

	return
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

func score1(numberOfMatches int) int {
	if numberOfMatches == 0 {
		return 0
	}

	return intPower(2, numberOfMatches-1)
}

func intPower(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

/*
  Sum scores of each card

	For each card, get the winningNumbers and the drawnNumbers
	Get the count of drawn numbers that are winning numbers
	If count > 0
		Return 2 ^ (count -1)
	else
		return 0

*/

func solve1(input string) int {
	lines := strings.Split(input, "\n")
	cards := lib.Reduce(lines, func(agg []Card, line string) []Card {
		return append(agg, parseCard(line))
	}, []Card{})

	return lib.Reduce(cards, func(agg int, card Card) int {
		return agg + card.Score1
	}, 0)
}

// This function assumes that all cards after it have a score2
func getScore2(card Card, allCards map[int]Card) int {
	result := 1
	for i := 0; i < card.NumberOfMatches; i++ {
		result += allCards[card.Number+i+1].Score2
	}
	return result
}

/*
	Alright, time to solve part 2

	First step will be to get the numbers of the cards
	Then we're going to store the cards in a map of allCards
	Also we'll figure out the new scores for the cards.

	Then either we iterate over them all, and do some caching or something like that
	OR
	We iterate over them backwards and determine the scoreCardsGenerated for each card
	Then we sum them all up
*/

func solve2(input string) int {
	lines := strings.Split(input, "\n")
	cardsMap := map[int]Card{}
	for _, lines := range lines {
		card := parseCard(lines)
		cardsMap[card.Number] = card
	}

	result := 0

	for i := len(lines); i > 0; i-- {
		card := cardsMap[i]
		card.Score2 = getScore2(card, cardsMap)
		cardsMap[i] = card
		result += card.Score2
	}

	return result
}

func main() {
	lib.AssertEqual(4, solve1(SmallTestString))
	lib.AssertEqual(13, solve1(TestString))
	lib.AssertEqual(30, solve2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solve1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solve2(dataString)
	// fmt.Println(result2)
}
