package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strings"
)

const TestString string = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const TestString2 string = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	First step is to create a queue of instructions
	Then we initialize the modules

	1000 times do this:
		Then add a low pulse to the broadcaster to the queue and run it
		Each pulse is added to a count of low and high pulses before it is ran
		When the queue is empty, send the low pulse again

	Multiply the low and high pulses


	broadcaster -> a, b, c
	%a -> b
	%b -> c
	%c -> inv
	&inv -> a

*/

type instruction struct {
	toMod   string
	fromMod string
	pulse   string
}

type module struct {
	moduleType string
	name       string
	targets    []string
	state      string
	memory     map[string]string
}

func (toMod module) sendPulse(pulse string, queue *[]instruction) {
	for _, target := range toMod.targets {
		*queue = append((*queue), instruction{target, toMod.name, pulse})
	}
}

func pulseModule(toModName, fromModName string, pulse string, queue *[]instruction, modules *map[string]module) {
	toMod := (*modules)[toModName]
	if toMod.moduleType == "broadcaster" {
		toMod.sendPulse(pulse, queue)
	} else if toMod.moduleType == "flip-flop" {
		if pulse == "low" {
			if toMod.state == "off" {
				toMod.state = "on"
				toMod.sendPulse("high", queue)
			} else if toMod.state == "on" {
				toMod.state = "off"
				toMod.sendPulse("low", queue)
			}
		}
	} else if toMod.moduleType == "conjunction" {
		toMod.memory[fromModName] = pulse
		allHigh := true
		for _, memoryPulse := range toMod.memory {
			if memoryPulse == "low" {
				allHigh = false
				break
			}
		}
		if allHigh {
			toMod.sendPulse("low", queue)
		} else {
			toMod.sendPulse("high", queue)
		}
	} else {
		// fmt.Println("In else")
		// fmt.Println("toModName: ", toModName)
		// fmt.Println(toMod.moduleType)
	}

	// Save changes to modules
	(*modules)[toModName] = toMod
}

func createModules(input string) map[string]module {
	modules := map[string]module{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		moduleDescriptor := parts[0]

		moduleType := ""
		name := ""
		if moduleDescriptor[0] == '%' {
			moduleType = "flip-flop"
			name = moduleDescriptor[1:]
		} else if moduleDescriptor[0] == '&' {
			moduleType = "conjunction"
			name = moduleDescriptor[1:]
		} else if moduleDescriptor == "broadcaster" {
			moduleType = moduleDescriptor
			name = moduleDescriptor
		} else {
			fmt.Println(line)
			fmt.Println(moduleDescriptor)
			panic("invalid module type")
		}

		targets := strings.Split(parts[1], ", ")
		newModule := module{moduleType, name, targets, "off", map[string]string{}}

		modules[name] = newModule
	}

	// Set memory for conjunctions and set module for output
	for _, mod := range modules {
		for _, target := range mod.targets {
			if _, ok := modules[target]; !ok {
				modules[target] = module{"output", target, []string{}, "off", map[string]string{}}
			}

			if modules[target].moduleType == "conjunction" {
				modules[target].memory[mod.name] = "low"
			}
		}
	}

	return modules
}

func runQueueUntilEmpty(queue *[]instruction, modules *map[string]module) (lowPulseCount, highPulseCount int) {
	for len(*queue) > 0 {
		instruction := (*queue)[0]
		*queue = (*queue)[1:]

		if instruction.pulse == "low" {
			lowPulseCount++
		} else if instruction.pulse == "high" {
			highPulseCount++
		}

		pulseModule(instruction.toMod, instruction.fromMod, instruction.pulse, queue, modules)
	}
	return lowPulseCount, highPulseCount
}

func solvePart1(input string) int {
	queue := []instruction{}
	modules := createModules(input)

	lowPulseCount := 0
	highPulseCount := 0

	i := 0
	for i < 1000 {
		queue = append(queue, instruction{"broadcaster", "button", "low"})

		lc, hc := runQueueUntilEmpty(&queue, &modules)
		lowPulseCount += lc
		highPulseCount += hc
		i++
	}

	return lowPulseCount * highPulseCount
}

/*
	Part 2 Notes

*/

func allKeysHaveValue(track map[string]int) bool {
	allHaveValue := true
	for _, trackValue := range track {
		if trackValue == 0 {
			allHaveValue = false
			break
		}
	}
	return allHaveValue
}

func allValues(track map[string]int) []int {
	values := []int{}
	for _, trackValue := range track {
		values = append(values, trackValue)
	}
	return values
}

func solvePart2(input string) int {
	queue := []instruction{}
	modules := createModules(input)

	trackLx := map[string]int{}
	for mod := range modules["lx"].memory {
		trackLx[mod] = 0
	}

	i := 0
	for i < 10000 {
		queue = append(queue, instruction{"broadcaster", "button", "low"})

		for len(queue) > 0 {
			instruction := queue[0]
			queue = queue[1:]

			// Set the cycle that the key sent high pulse to lx
			if instruction.toMod == "lx" && instruction.pulse == "high" && trackLx[instruction.fromMod] == 0 {
				trackLx[instruction.fromMod] = i + 1
			}

			pulseModule(instruction.toMod, instruction.fromMod, instruction.pulse, &queue, &modules)
		}

		if allKeysHaveValue(trackLx) {
			return lib.LcmOfSlice(allValues(trackLx))
		}

		i++
	}

	return -1
}

func main() {
	lib.AssertEqual(32000000, solvePart1(TestString))
	lib.AssertEqual(11687500, solvePart1(TestString2))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
