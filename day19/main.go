package main

import (
	"advent-of-code-2023/lib"
	"regexp"
	"strconv"
	"strings"
)

const SmallTestString string = ``

const TestString string = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

const DataFile string = "data.txt"

/*
	Part 1 Notes

*/

type part struct {
	x int
	m int
	a int
	s int
}

type workflowMap map[string][]string

func parseParts(input string) (parts []part) {
	partLines := strings.Split(input, "\n")

	for _, line := range partLines {
		newPart := part{}

		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)

		newPart.x, _ = strconv.Atoi(matches[0])
		newPart.m, _ = strconv.Atoi(matches[1])
		newPart.a, _ = strconv.Atoi(matches[2])
		newPart.s, _ = strconv.Atoi(matches[3])

		parts = append(parts, newPart)
	}

	return
}

func parseWorkflows(input string) (workflows workflowMap) {
	workflows = make(workflowMap)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, "{")
		rules := strings.Trim(parts[1], "}")

		workflows[parts[0]] = strings.Split(rules, ",")
	}

	return
}

func sumPart(p part) int {
	return p.x + p.m + p.a + p.s
}

func executeString(instruction string, p part, workflows workflowMap) bool {
	if instruction == "R" {
		return false
	} else if instruction == "A" {
		return true
	} else {
		return followWorkflow(p, workflows[instruction], workflows)
	}
}

func comparePart(p part, key, comparison string, value int) bool {
	partValue := 0
	switch key {
	case "x":
		partValue = p.x
	case "m":
		partValue = p.m
	case "a":
		partValue = p.a
	case "s":
		partValue = p.s
	default:
		panic("we done goofed")
	}

	if comparison == "<" {
		return partValue < value
	} else {
		return partValue > value
	}
}

func followWorkflow(p part, workflow []string, workflows workflowMap) bool {
	for _, rule := range workflow {
		if !strings.Contains(rule, ":") {
			return executeString(rule, p, workflows)
		} else {
			splitString := strings.Split(rule, ":")
			conditional := splitString[0]
			execute := splitString[1]
			key, comparison, valueString := "", "", ""
			if strings.Contains(conditional, "<") {
				key = strings.Split(conditional, "<")[0]
				comparison = "<"
				valueString = strings.Split(conditional, "<")[1]
			} else {
				key = strings.Split(conditional, ">")[0]
				comparison = ">"
				valueString = strings.Split(conditional, ">")[1]
			}
			value, _ := strconv.Atoi(valueString)

			if comparePart(p, key, comparison, value) {
				return executeString(execute, p, workflows)
			}
		}
	}
	panic("workflow ended without action")
}

func solvePart1(input string) int {
	groups := strings.Split(input, "\n\n")

	parts := parseParts(groups[1])
	workflows := parseWorkflows(groups[0])

	sum := 0

	for _, part := range parts {
		if followWorkflow(part, workflows["in"], workflows) {
			sum += sumPart(part)
		}
	}

	return sum
}

/*
	Part 2 Notes

	Change the structure of parts to have xMin xMax, mMin, mMax etc.
	Change partSum to partPossibilities, which multiplies the lengths of each of the values
	Change the logic so that we start out with potentialParts with ranges from 1 to 4000 for every value
		When we reach a coniditional, we do recursive logic and return the possibilities of each added together

	We're going to need to be able to start at different rules within the workflows for this to work
*/

type partPossibilities struct {
	xMin int
	xMax int
	mMin int
	mMax int
	aMin int
	aMax int
	sMin int
	sMax int
}

func originalPossibilities() partPossibilities {
	return partPossibilities{1, 4000, 1, 4000, 1, 4000, 1, 4000}
}

func (p partPossibilities) possibilityCount() int {
	return (p.xMax - p.xMin + 1) * (p.mMax - p.mMin + 1) * (p.aMax - p.aMin + 1) * (p.sMax - p.sMin + 1)
}

func splitPossibility(p partPossibilities, key, comparison string, value int) (partPossibilities, partPossibilities) {
	truePossibility := p
	falsePossibility := p

	switch key {
	case "x":
		if comparison == "<" {
			truePossibility.xMax = lib.Min(truePossibility.xMax, value-1)
			falsePossibility.xMin = lib.Max(truePossibility.xMin, value)
		} else {
			truePossibility.xMin = lib.Max(truePossibility.xMin, value+1)
			falsePossibility.xMax = lib.Min(truePossibility.xMax, value)
		}
	case "m":
		if comparison == "<" {
			truePossibility.mMax = lib.Min(truePossibility.mMax, value-1)
			falsePossibility.mMin = lib.Max(truePossibility.mMin, value)
		} else {
			truePossibility.mMin = lib.Max(truePossibility.mMin, value+1)
			falsePossibility.mMax = lib.Min(truePossibility.mMax, value)
		}
	case "a":
		if comparison == "<" {
			truePossibility.aMax = lib.Min(truePossibility.aMax, value-1)
			falsePossibility.aMin = lib.Max(truePossibility.aMin, value)
		} else {
			truePossibility.aMin = lib.Max(truePossibility.aMin, value+1)
			falsePossibility.aMax = lib.Min(truePossibility.aMax, value)
		}
	case "s":
		if comparison == "<" {
			truePossibility.sMax = lib.Min(truePossibility.sMax, value-1)
			falsePossibility.sMin = lib.Max(truePossibility.sMin, value)
		} else {
			truePossibility.sMin = lib.Max(truePossibility.sMin, value+1)
			falsePossibility.sMax = lib.Min(truePossibility.sMax, value)
		}
	}

	return truePossibility, falsePossibility
}

func executeString2(instruction string, p partPossibilities, workflows workflowMap) int {
	if instruction == "R" {
		return 0
	} else if instruction == "A" {
		return p.possibilityCount()
	} else {
		return possibilities(p, workflows[instruction], 0, workflows)
	}
}

func possibilities(p partPossibilities, workflow []string, ruleIndex int, workflows workflowMap) int {
	rule := workflow[ruleIndex]

	if !strings.Contains(rule, ":") {
		return executeString2(rule, p, workflows)
	} else {
		splitString := strings.Split(rule, ":")
		conditional := splitString[0]
		execute := splitString[1]
		key, comparison, valueString := "", "", ""
		if strings.Contains(conditional, "<") {
			key = strings.Split(conditional, "<")[0]
			comparison = "<"
			valueString = strings.Split(conditional, "<")[1]
		} else {
			key = strings.Split(conditional, ">")[0]
			comparison = ">"
			valueString = strings.Split(conditional, ">")[1]
		}
		value, _ := strconv.Atoi(valueString)

		truePossibility, falsePossibility := splitPossibility(p, key, comparison, value)
		return executeString2(execute, truePossibility, workflows) + possibilities(falsePossibility, workflow, ruleIndex+1, workflows)
	}
}

func solvePart2(input string) int {
	groups := strings.Split(input, "\n\n")
	workflows := parseWorkflows(groups[0])
	allPossibleParts := originalPossibilities()

	return possibilities(allPossibleParts, workflows["in"], 0, workflows)
}

func main() {
	lib.AssertEqual(19114, solvePart1(TestString))
	lib.AssertEqual(167409079868000, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
