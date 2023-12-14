package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
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
	digits, string_had_digit := first_and_last_digit(line)

	if string_had_digit == false {
		return 0, fmt.Errorf("No number in the string. Digits: %v, Line: %s", digits, line)
	}

	return ((digits[0] * 10) + digits[1]), nil
}

func first_and_last_digit(line string) ([2]int, bool) {
	var digits [2]int
	var passed_digit bool = false

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			value := int_from_unicode(runeValue)
			if passed_digit {
				digits = [2]int{digits[0], value}
			} else {
				digits = [2]int{value, value}
			}
			passed_digit = true
		}
	}

	return digits, passed_digit
}

func int_from_unicode(runeValue rune) int {
	return int(runeValue - '0')
}
