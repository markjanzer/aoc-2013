package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// old value 11, new value 29
	test_calibration_value("two1nine", 11)
	// old value 24, new value 14
	test_calibration_value("zoneight234", 24)
	// old value 22, new value 13
	test_calibration_value("abcone2threexyz", 22)

	return

	// Read text from https://adventofcode.com/2023/day/1/input
	// var calibration_codes, err1 = get_calibration_codes()
	// if err1 != nil {
	// 	fmt.Println("Error:", err1)
	// 	return
	// }

	// // Convert each line to a number
	// var calibration_values []int = make([]int, len(calibration_codes))
	// for i, line := range calibration_codes {
	// 	var value, err2 = calibration_value(line)
	// 	if err2 != nil {
	// 		fmt.Println("Error:", err2)
	// 		return
	// 	} else {
	// 		calibration_values[i] = value
	// 	}
	// }

	// var sum int = 0
	// for _, value := range calibration_values {
	// 	sum += value
	// }

	// fmt.Println("Sum:", sum)
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

func first_and_last_digit(line string) ([]int, bool) {
	var digits = make([]int, 0, 2)
	var passed_digit bool = false
	var sequential_letters string = ""

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			sequential_letters = ""
			value := int_from_unicode(runeValue)
			if len(digits) == 0 {
				digits = append(digits, value)
				digits = append(digits, value)
			} else {
				digits[1] = value
			}
			passed_digit = true
		} else {
			sequential_letters += string(runeValue)
		}
	}

	return digits, passed_digit
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
