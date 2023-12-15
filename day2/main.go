package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const RedMax = 12
const GreenMax = 13
const BlueMax = 14

const DataFile string = "./data.txt"

// Set type as defined earlier
type Set map[string]int

// Game struct with an ID and a slice of Sets
type Game struct {
	ID   int
	Sets []Set
}

func getGameStrings() ([]string, error) {
	file, err := os.Open(DataFile)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	// var games []Game
	var gameStrings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameStrings = append(gameStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	return gameStrings, nil
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
				return game, fmt.Errorf("Invalid color: %s", color)
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

func main() {
	gameStrings, err := getGameStrings()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	games := make([]Game, 0, len(gameStrings))
	for _, gameString := range gameStrings {
		game, err := parseGame(gameString)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		if gameIsPossible(game) {
			sum += game.ID
		}
	}

	fmt.Println("Sum:", sum)

	// result, err := parseGame("Game 1: 12 red, 2 green, 5 blue; 9 red, 6 green, 4 blue; 10 red, 2 green, 5 blue; 8 blue, 9 red")

	fmt.Println("done")
}
