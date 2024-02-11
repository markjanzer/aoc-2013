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
	// fmt.Println("in executeString, instruction: ", instruction)
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
	// fmt.Println(workflow)
	for _, rule := range workflow {
		// fmt.Println(p, rule)
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

*/

func solvePart2(input string) int {
	return 0
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
