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

// // Implement this to refactor a little
// func (mod module)sendPulse(pulse string, queue []instruction) {
// 	for _, target := range toMod.targets {
// 		queue = append(queue, instruction{target, pulse})
// 	}
// }

// type flipFlop struct {
// 	name string
// 	state string
// }

func pulseModule(toMod module, fromModName string, pulse string, queue *[]instruction) {
	if toMod.moduleType == "broadcaster" {
		for _, target := range toMod.targets {
			*queue = append(*queue, instruction{target, toMod.name, pulse})
		}
	} else if toMod.moduleType == "flip-flop" {
		if pulse == "low" {
			if toMod.state == "off" {
				toMod.state = "on"
				for _, target := range toMod.targets {
					*queue = append(*queue, instruction{target, toMod.name, "high"})
				}
			} else if toMod.state == "on" {
				toMod.state = "off"
				for _, target := range toMod.targets {
					*queue = append(*queue, instruction{target, toMod.name, "low"})
				}
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
			for _, target := range toMod.targets {
				*queue = append(*queue, instruction{target, toMod.name, "low"})
			}
		} else {
			for _, target := range toMod.targets {
				*queue = append(*queue, instruction{target, toMod.name, "high"})
			}
		}
	}
}

func solvePart1(input string) int {
	queue := []instruction{}

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

	// Set memory for conjunctions
	for modName, mod := range modules {
		for _, target := range mod.targets {
			if modules[target].moduleType == "conjunction" {
				modules[target].memory[modName] = "low"
			}
		}
	}

	lowPulseCount := 0
	highPulseCount := 0

	queue = append(queue, instruction{"broadcaster", "button", "low"})
	for len(queue) > 0 {
		instruction := queue[0]
		queue = queue[1:]

		fmt.Println(instruction.fromMod, instruction.pulse, instruction.toMod)

		if instruction.pulse == "low" {
			lowPulseCount++
		} else if instruction.pulse == "high" {
			highPulseCount++
		}

		pulseModule(modules[instruction.toMod], instruction.fromMod, instruction.pulse, &queue)
	}

	fmt.Println(lowPulseCount, highPulseCount)

	return lowPulseCount * highPulseCount
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	// Only running once, it should be 8 * 4
	lib.AssertEqual(32, solvePart1(TestString))
	// lib.AssertEqual(32000000, solvePart1(TestString))
	// lib.AssertEqual(11687500, solvePart1(TestString2))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
