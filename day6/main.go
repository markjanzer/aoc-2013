package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strings"
)

const SmallTestString string = ``

const TestString string = `Time:      7  15   30
Distance:  9  40  200`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	For each game we have to find out how many different ways there are to beat the best score
	Then we multiply them together

	To do this let's start by defining the games with their time and their current records

*/

type Game struct {
	Time             int
	Record           int
	PossibleOutcomes []Outcome
}

type Outcome struct {
	HeldFor      int
	TravelledFor int
	Distance     int
}

func makeGames(input string) (games []Game) {
	lines := strings.Split(input, "\n")
	timesLine, recordsLine := lines[0], lines[1]
	times := lib.IntsFromString(strings.Split(timesLine, ":")[1])
	records := lib.IntsFromString(strings.Split(recordsLine, ":")[1])

	for i := 0; i < len(times); i += 1 {
		newGame := Game{Time: times[i], Record: records[i]}
		newGame.PossibleOutcomes = determineOutcomes(newGame)
		games = append(games, newGame)
	}

	return
}

func determineOutcomes(game Game) (outcomes []Outcome) {
	for i := 1; i <= game.Time; i++ {
		heldFor := i
		travelledFor := game.Time - i
		distance := heldFor * travelledFor
		newOutcome := Outcome{HeldFor: heldFor, TravelledFor: travelledFor, Distance: distance}
		outcomes = append(outcomes, newOutcome)
	}
	return
}

func numberOfWinningOutcomes(game Game) (number int) {
	for _, outcome := range game.PossibleOutcomes {
		if outcome.Distance > game.Record {
			number += 1
		}
	}
	return
}

func solvePart1(input string) int {
	games := makeGames(input)
	result := lib.Reduce(games, func(sum int, game Game) int {
		return sum * numberOfWinningOutcomes(game)
	}, 1)

	// fmt.Println(games)
	// fmt.Println(result)

	return result
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	// lib.AssertEqual(49, solvePart1(SmallTestString))
	// lib.AssertEqual(49, solvePart2(SmallTestString))

	lib.AssertEqual(288, solvePart1(TestString))
	// lib.AssertEqual(46, solvePart2(TestString))

	dataString := lib.GetDataString(DataFile)

	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
