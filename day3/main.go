package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const TestString string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

const DataFile string = "data.txt"

// We can assume that each line of the string is the same length

/*
  How to go about this?

	What I could do is identify all of the numbers, along with their coordinates
	I could also identify all of the non . characters, and their coordinates

	Then I could write a function that takes a given set of coordinates, and returns a range or coordinates
	that are within one step of the given coordinates

	Then I could write a function that takes a set of coordinates and determines if it is with the plane

	Then for each number, I could see if there is any number within the given set of coordinates

	Finally, I would add the numbers together that are adjacent to a symbol
*/

type Set map[string]int

// Game struct with an ID and a slice of Sets
type Game struct {
	ID   int
	Sets []Set
}

type Plane struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

type Point struct {
	x int
	y int
}

type Number struct {
	value int
	plane Plane
}

const TestIdentify string = `42..
$...
..35
...*`

func getDataString() (data string) {
	file, err := os.Open(DataFile)
	assertNoError((err))
	defer file.Close()

	// Read file content into a byte slice
	byteContent, err := io.ReadAll(file)
	assertNoError(err)

	return string(byteContent)
}

func stringToMatrix(input string) (matrix [][]byte) {

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}

	return
}

func byteIsDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func intFromByte(b byte) int {
	return int(b) - 48
}

func byteIsPeriod(b byte) bool {
	return b == 46
}

func identifyNumbersAndSymbols(matrix [][]byte) (numbers []Number, symbols []Point) {
	for y, line := range matrix {
		usingNumber := false
		tempNumber := Number{}
		for x, b := range line {
			if byteIsDigit(b) {
				usingNumber = true
				tempNumber = addToNumber(tempNumber, x, y, intFromByte(b))
			} else {
				if usingNumber {
					numbers = append(numbers, tempNumber)
					tempNumber = Number{}
					usingNumber = false
				}
			}

			if !byteIsDigit(b) && !byteIsPeriod(b) {
				symbols = append(symbols, Point{x, y})
			}
		}
		if usingNumber {
			numbers = append(numbers, tempNumber)
		}
		tempNumber = Number{}
		usingNumber = false
	}
	return
}

func addToNumber(num Number, x, y, b int) Number {
	if num.value == 0 {
		num.value = b
		num.plane = Plane{
			xMin: x,
			xMax: x,
			yMin: y,
			yMax: y,
		}
	} else {
		num.value = num.value*10 + b
		num.plane = Plane{
			xMin: num.plane.xMin,
			xMax: x,
			yMin: num.plane.yMin,
			yMax: y,
		}
	}
	return num
}

func findNumbersNextToSymbol(numbers []Number, symbols []Point) (numbersNextToSymbol []Number) {
	for _, number := range numbers {
		includedPlane := expandedPlane(number.plane)
		for _, symbol := range symbols {
			if pointInPlane(symbol, includedPlane) {
				numbersNextToSymbol = append(numbersNextToSymbol, number)
			}
		}
	}
	return
}

func findGearsNextToNumber()

func expandedPlane(plane Plane) Plane {
	return Plane{
		xMin: plane.xMin - 1,
		xMax: plane.xMax + 1,
		yMin: plane.yMin - 1,
		yMax: plane.yMax + 1,
	}
}

func pointInPlane(point Point, plane Plane) bool {
	return (point.x >= plane.xMin &&
		point.x <= plane.xMax &&
		point.y >= plane.yMin &&
		point.y <= plane.yMax)
}

// Might use this
// func planesOverlap(plane1, plane2 Plane) bool {
// 	return (plane1.xMin <= plane2.xMax &&
// 		plane1.xMax >= plane2.xMin &&
// 		plane1.yMin <= plane2.yMax &&
// 		plane1.yMax >= plane2.yMin)
// }

func sumNumbers(numbers []Number) (sum int) {
	for _, number := range numbers {
		sum += number.value
	}
	return
}

func solve1(input string) int result {
	matrix := stringToMatrix(input)
	numbers, symbols := identifyNumbersAndSymbols(matrix)
	numbersNextToSymbols := findNumbersNextToSymbol(numbers, symbols)
	result := sumNumbers(numbersNextToSymbols)
	return
}

func solve2(input string) int result {
	// matrix := stringToMatrix(input)
	// numbers, symbols := identifyNumbersAndSymbols(matrix)
	// numbersNextToSymbols := findNumbersNextToSymbol(numbers, symbols)
	// result := sumNumbers(numbersNextToSymbols)
	return 0
}


func main() {


	// fmt.Println("Testing identifyNumbers function")
	// in := identifyNumbers(TestString)

	assertEqual(4361, solve1(TestString))
	assertEqual(467835, solve2(TestString))

	dataString := getDataString()
	result1 := solve1(dataString)
	result2 := solve2(dataString)

	fmt.Println(result1)
	fmt.Println(result2)
}

// Helpers
func assertNoError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func assertEqual(expected, actual int) {
	if expected != actual {
		fmt.Println(fmt.Sprintf("Test failed \n\texpected: %d, got: %d", expected, actual))
	} else {
		fmt.Println("Test passed")
	}
}
