package main

import (
	"advent-of-code-2023/lib"
	"regexp"
	"strings"
)

const SmallTestString string = `seeds: 53

seed-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4`

const TestString string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

const DataFile string = "data.txt"

type Map struct {
	MapLines []MapLine
	Name     string
}
type MapLine struct {
	Destination int
	Source      int
	Range       int
}

/*
	Part 1 Notes

	Iterate over the string
	For each Map found, append it to the almanac
		For each line in the Map, add a MapLine with Destination, Source, and Range

	For each seed number, iterate over the maps, getting the new number and passing it to
	the next Map. If that number is the lowest, set it as the result.

	To iterate over maps, we have a function that takes an number and a map and returns
	the new int
*/

func getSeeds(input string) (seeds []int) {
	splitString := strings.Split(input, "\n\n")
	seedsString := strings.Split(splitString[0], ":")[1]
	seeds = lib.IntsFromString(seedsString)
	return
}

func makeMaps(input string) (maps []Map) {
	mapsStrings := strings.Split(input, "\n\n")[1:]

	return lib.Map(mapsStrings, func(mapString string) Map {
		return makeMap(mapString)
	})
}

func makeMap(mapString string) (thisMap Map) {
	splitString := strings.Split(mapString, "\n")
	thisMap.Name = extractName(splitString[0])
	mapLines := splitString[1:]

	for _, line := range mapLines {
		thisMap.MapLines = append(thisMap.MapLines, makeMapLine(line))
	}
	return
}

func extractName(line string) string {
	re := regexp.MustCompile(`-to-(\w+)\s`)
	matches := re.FindStringSubmatch(line)
	return matches[1]
}

func makeMapLine(line string) MapLine {
	numbers := lib.IntsFromString(line)
	destination, source, rang := numbers[0], numbers[1], numbers[2]
	return MapLine{destination, source, rang}
}

func transformSeedThroughAllMaps(seed int, maps []Map) int {
	for _, thisMap := range maps {
		seed = transformSeedThroughMap(seed, thisMap)
	}
	return seed
}

func transformSeedThroughMap(seed int, thisMap Map) int {
	for _, mapLine := range thisMap.MapLines {
		seed, transformed := transformSeedThroughMapLine(seed, mapLine)
		if transformed {
			return seed
		}
	}
	return seed
}

// Determines if seed is within range of MapLine based on source and range
// Need to do -1 or else range would be extened
func seedIsInRange(seed int, mapLine MapLine) bool {
	return lib.IntIsInRange(seed, mapLine.Source, (mapLine.Source + mapLine.Range - 1))
}

func transformSeedThroughMapLine(seed int, mapLine MapLine) (int, bool) {
	if seedIsInRange(seed, mapLine) {
		transformation := mapLine.Destination - mapLine.Source
		return (seed + transformation), true
	} else {
		return seed, false
	}
}

func smallestResultingValue(seeds []int, maps []Map) int {
	minResult := 1000000000
	for _, seed := range seeds {
		result := transformSeedThroughAllMaps(seed, maps)
		if result < minResult {
			minResult = result
		}
	}

	return minResult
}

func solvePart1(input string) int {
	seeds := getSeeds(input)
	maps := makeMaps(input)

	return smallestResultingValue(seeds, maps)
}

/*
	Part 2 Notes

	The seed numbers are ranges now. We're just going to change the way that
	seeds are calculated and then run it again.
*/

func getSeedsPart2(input string) (seeds []int) {
	numbers := getSeeds(input)

	for i := 0; i < len(numbers); i += 2 {
		start := numbers[i]
		length := numbers[i+1]
		for seedValue := start; seedValue < start+length; seedValue++ {
			seeds = append(seeds, seedValue)
		}
	}
	return
}

func solvePart2(input string) int {
	seeds := getSeedsPart2(input)
	maps := makeMaps(input)

	return smallestResultingValue(seeds, maps)
}

func main() {
	lib.AssertEqual(49, solvePart1(SmallTestString))
	lib.AssertEqual(35, solvePart1(TestString))
	lib.AssertEqual(46, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
