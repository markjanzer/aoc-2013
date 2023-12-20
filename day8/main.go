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


*/

type Journey struct {
	Start        string
	Current      string
	ResolvedOnce bool
}

func newJourney(key string) Journey {
	return Journey{
		Start:   key,
		Current: key,
	}
}

func (journey *Journey) atDestination() bool {
	atDestination := string(journey.Current[2]) == "Z"
	if atDestination && !journey.ResolvedOnce {
		journey.ResolvedOnce = true
	}
	return atDestination
}

func (journey *Journey) travel(direction string, nodeMap NodeMap) {
	newLocation := nodeMap[journey.Current][direction]
	journey.Current = newLocation
}

type Journeys []Journey

func (journeys Journeys) allAtDestination() bool {
	for i := range journeys {
		if !journeys[i].atDestination() {
			return false
		}
	}
	return true
}

func (journeys Journeys) travel(direction string, nodeMap NodeMap) {
	for i := range journeys {
		journeys[i].travel(direction, nodeMap)
	}
}

func solvePart2(input string) int {
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

	count := 0
	for !journeys.allAtDestination() && count < len(inputString) {
		inputIndex := count % len(inputString)
		direction := inputString[inputIndex]
		fmt.Println(direction)
		journeys.travel(direction, nodeMap)
		fmt.Println(journeys)
		count++
	}
	fmt.Println(journeys)
	return count
}

func main() {
	lib.AssertEqual(2, solvePart1(TestString))
	lib.AssertEqual(6, solvePart1(SmallTestString))

	lib.AssertEqual(6, solvePart2(TestString2))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	dataString := lib.GetDataString(DataFile)
	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
