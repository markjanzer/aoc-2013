package main

import (
	"fmt"
)

const RedMax = 12
const GreenMax = 13
const BlueMax = 14

// Set type as defined earlier
type Set map[string]int

// Game struct with an ID and a slice of Sets
type Game struct {
	ID        int
	ColorMaps []Set
}

func main() {
	// old value 11, new value 29
	// test_calibration_value("two1nine", 29)
	// old value 24, new value 14
	// test_calibration_value("zoneight234", 14)
	// old value 22, new value 13
	// test_calibration_value("abcone2threexyz", 13)
	// return

	fmt.Println("Sum:", sum)
}
