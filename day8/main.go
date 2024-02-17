package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"regexp"
	"strings"
)

const SmallTestString string = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const TestString string = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const TestString2 string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

const DataFile string = "data.txt"

/*
	Part 1 Notes

*/

type Node map[string]string
type NodeMap map[string]Node

func makeNode(left string, right string) Node {
	node := make(Node)
	node["L"] = left
	node["R"] = right
	return node
}

func makeNodeMap(nodesString string) NodeMap {
	var nodeMap NodeMap = make(NodeMap)
	lines := strings.Split(nodesString, "\n")
	for _, line := range lines {
		// Get the name, left and right from AAA = (BBB, CCC)
		re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
		matches := re.FindStringSubmatch(line)
		name, left, right := matches[1], matches[2], matches[3]
		newNode := makeNode(left, right)
		nodeMap[name] = newNode
	}
	return nodeMap
}

func navigateMap(inputString []string, nodeMap NodeMap) int {
	currentNode := "AAA"
	count := 0
	for currentNode != "ZZZ" {
		inputIndex := count % len(inputString)
		direction := inputString[inputIndex]
		currentNode = nodeMap[currentNode][direction]
		count++
	}
	return count
}

func solvePart1(input string) int {
	splitString := strings.Split(input, "\n\n")
	inputString := strings.Split(splitString[0], "")
	nodesString := splitString[1]
	nodeMap := makeNodeMap(nodesString)

	steps := navigateMap(inputString, nodeMap)
	return steps
}

/*
	Part 2 Notes

	Find all of the starting paths (every node that ends with A)
	Change to be able to iterate while navigating
	Check all of them when iterating


	This approach does not work, the potential number is too big
	If each journey has a firstComplete and loopInterval then we can look into another solution
*/

type Journey struct {
	Start       string
	Current     string
	CompletedAt []int
}

func newJourney(key string) Journey {
	return Journey{
		Start:   key,
		Current: key,
	}
}

func (journey *Journey) checkCompletion(count int) {
	atDestination := string(journey.Current[2]) == "Z"
	if atDestination {
		journey.CompletedAt = append(journey.CompletedAt, count)
	}
}

func (journey *Journey) travel(direction string, nodeMap NodeMap) {
	newLocation := nodeMap[journey.Current][direction]
	journey.Current = newLocation
}

func (journey Journey) hasConsistentLoopInterval() bool {
	if len(journey.CompletedAt) < 3 {
		panic("Not enough data")
	}

	firstLoop := journey.CompletedAt[1] - journey.CompletedAt[0]
	secondLoop := journey.CompletedAt[1] - journey.CompletedAt[0]

	return firstLoop == secondLoop
}

func (journey Journey) initialComplete() int {
	return journey.CompletedAt[0]
}

func (journey Journey) loop() int {
	return journey.CompletedAt[2] - journey.CompletedAt[1]
}

type Journeys []Journey

func getDirection(directions []string, count int) string {
	inputIndex := count % len(directions)
	return directions[inputIndex]
}

const MaxIterations int = 10000000

func solvePart2(input string) uint64 {
	splitString := strings.Split(input, "\n\n")
	inputString := strings.Split(splitString[0], "")
	nodesString := splitString[1]
	nodeMap := makeNodeMap(nodesString)

	journeys := Journeys{}
	for nodeKey := range nodeMap {
		if string(nodeKey[2]) == "A" {
			journeys = append(journeys, newJourney(nodeKey))
		}
	}

	// Starting points
	// PQA  CQA  TGA  AAA  BLA  DFA

	count := 0
	allJourneysHaveTwoLoops := false
	for !allJourneysHaveTwoLoops && count < MaxIterations {
		direction := getDirection(inputString, count)

		for i := range journeys {
			journeys[i].travel(direction, nodeMap)
			journeys[i].checkCompletion(count)
		}

		allJourneysHaveTwoLoops = lib.All(journeys, func(journey Journey) bool {
			return len(journey.CompletedAt) > 2
		})

		count++
	}

	allJourneysHaveConsistentLoopInterval := lib.All(journeys, func(journey Journey) bool {
		return journey.hasConsistentLoopInterval()
	})
	if !allJourneysHaveConsistentLoopInterval {
		panic("All journeys do not have consistent loop intervals")
	}

	var loops = []uint64{}
	for _, journey := range journeys {
		loops = append(loops, uint64(journey.loop()))
	}

	result := lib.LcmOfSlice(loops)

	// Only the first value is getting set to be true
	// fmt.Println(journeys)
	return result
}

func main() {
	// lib.AssertEqual(2, solvePart1(TestString))
	// lib.AssertEqual(6, solvePart1(SmallTestString))

	lib.AssertEqual(6, int(solvePart2(TestString2)))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// fmt.Println(gcd(200, 80))
	// fmt.Println(lib.LcmOfSlice([]uint64{3, 6, 10}))

	dataString := lib.GetDataString(DataFile)
	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
