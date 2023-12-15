package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// old value 11, new value 29
	// test_calibration_value("two1nine", 29)
	// old value 24, new value 14
	// test_calibration_value("zoneight234", 14)
	// old value 22, new value 13
	// test_calibration_value("abcone2threexyz", 13)
	// return

	// Read text from https://adventofcode.com/2023/day/1/input
	var calibration_codes, err1 = get_calibration_codes()
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}

	// Convert each line to a number
	var calibration_values []int = make([]int, len(calibration_codes))
	for i, line := range calibration_codes {
		var value, err2 = calibration_value(line)
		if err2 != nil {
			fmt.Println("Error:", err2)
			return
		} else {
			calibration_values[i] = value
		}
	}

	var sum int = 0
	for _, value := range calibration_values {
		sum += value
	}

	// Should return 56324
	fmt.Println("Sum:", sum)
}

func get_calibration_codes() ([]string, error) {
	file, err := os.Open("1-trebuchet-data.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	var codes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	return codes, nil
}

// Convert string to number
func calibration_value(line string) (int, error) {
	digits := first_and_last_digit(line)

	if len(digits) == 0 {
		return 0, fmt.Errorf("No number in the string. Digits: %v, Line: %s", digits, line)
	}

	return ((digits[0] * 10) + digits[1]), nil
}

var digit_names = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func first_and_last_digit(line string) []int {
	var digits = make([]int, 0, 2)
	var sequential_letters string = ""

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			sequential_letters = ""
			digits = append_digit(digits, int_from_unicode(runeValue))
		} else {
			sequential_letters += string(runeValue)
			var count = len(sequential_letters)
			for name, value := range digit_names {
				if count >= len(name) && sequential_letters[count-len(name):] == name {
					digits = append_digit(digits, value)
				}
			}
		}
	}

	return digits
}

func append_digit(digits []int, digit int) []int {
	if len(digits) == 0 {
		digits = append(digits, digit)
		digits = append(digits, digit)
	} else {
		digits[1] = digit
	}
	return digits
}

func int_from_unicode(runeValue rune) int {
	return int(runeValue - '0')
}

func test_calibration_value(code string, expected int) {
	var result, _ = calibration_value(code)
	if result != expected {
		fmt.Println(fmt.Sprintf("Failure: expected %d, got %d", expected, result))
		return
	} else {
		fmt.Println("Success")
	}
}
