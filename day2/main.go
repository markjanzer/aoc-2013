package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const RedMax = 12
const GreenMax = 13
const BlueMax = 14

const TestString = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

const DataFile string = "./data.txt"

// Set type as defined earlier
type Set map[string]int

// Game struct with an ID and a slice of Sets
type Game struct {
	ID   int
	Sets []Set
}

var validColors = []string{"red", "green", "blue"}

func parseGame(gameString string) (Game, error) {
	game := Game{}

	re := regexp.MustCompile(`Game (\d+): (.*)`)
	matches := re.FindStringSubmatch(gameString)
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return game, err
	}
	game.ID = id

	setsString := strings.Split(matches[2], ";")
	for _, setString := range setsString {
		set := Set{}
		colorStrings := strings.Split(setString, ",")
		for _, colorString := range colorStrings {
			colorString := strings.TrimSpace(colorString)
			colorInfo := strings.Split(colorString, " ")
			count, err := strconv.Atoi(colorInfo[0])
			if err != nil {
				return game, err
			}
			color := colorInfo[1]

			var validColorsMap = make(map[string]bool)
			for _, validColor := range validColors {
				validColorsMap[validColor] = true
			}

			if _, valid := validColorsMap[color]; !valid {
				return game, fmt.Errorf("invalid color: %s", color)
			}

			set[color] = count

		}

		for _, validColor := range validColors {
			if _, ok := set[validColor]; !ok {
				set[validColor] = 0
			}
		}

		game.Sets = append(game.Sets, set)
	}

	return game, nil
}

func gameIsPossible(game Game) bool {
	for _, set := range game.Sets {
		if set["red"] > RedMax || set["green"] > GreenMax || set["blue"] > BlueMax {
			return false
		}
	}

	return true
}

func minimumPossibleColors(game Game) Set {
	minimumSet := Set{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range game.Sets {
		for _, color := range validColors {
			if set[color] > minimumSet[color] {
				minimumSet[color] = set[color]
			}
		}
	}

	return minimumSet
}

func gamePower(game Game) int {
	minimumSet := minimumPossibleColors(game)

	return minimumSet["red"] * minimumSet["green"] * minimumSet["blue"]
}

func makeGames(input string) (games []Game) {
	gameStrings := strings.Split(input, "\n")

	for _, gameString := range gameStrings {
		game, err := parseGame(gameString)
		if err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}

		games = append(games, game)
	}

	return
}

func solvePart1(input string) int {
	games := makeGames(input)

	sum := 0
	for _, game := range games {
		if gameIsPossible(game) {
			sum += game.ID
		}
	}

	return sum
}

func solvePart2(input string) int {
	games := makeGames(input)

	sum := 0
	for _, game := range games {
		sum += gamePower(game)
	}

	return sum
}

func main() {
	lib.AssertEqual(8, solvePart1(TestString))
	lib.AssertEqual(2286, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
