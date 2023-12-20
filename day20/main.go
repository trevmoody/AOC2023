package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"slices"
	"strings"
)

type ModuleType int

const ( // iota is reset to 0
	broadcast ModuleType = iota
	flipFlop
	conjunction
)

type PulseType int

const (
	low PulseType = iota
	high
)

type PulseData struct {
	destinationName string
	pulseType       PulseType
	sourceName      string
}

type Module struct {
	name                     string
	moduleType               ModuleType
	destinationModuleNames   []string
	flipFlopState            bool
	conjunctionModuleHistory map[string]PulseType
}

func main() {
	lines := *util.GetFileAsLines("input")
	part1(lines)
	part2(lines, "rx")
}

func part2(lines []string, destinationToTrack string) {
	moduleMap := parseInput(lines)

	// assume only one conjunction sends to wanted destination
	var parentConjunction Module
	for _, module := range moduleMap {
		if slices.Contains(module.destinationModuleNames, destinationToTrack) {
			parentConjunction = module
		}
	}

	// assume only conjunctions send to parent conjunction
	destinationsToTrack := map[string]int{}
	for _, module := range moduleMap {
		if slices.Contains(module.destinationModuleNames, parentConjunction.name) {
			destinationsToTrack[module.name] = 0
		}
	}

	for count := 1; ; count++ {
		var pulseDestinationsToProcess []PulseData
		pulseDestinationsToProcess = append(pulseDestinationsToProcess, PulseData{"broadcaster", low, "button"})

		for i := 0; i < len(pulseDestinationsToProcess); i++ {
			pd := pulseDestinationsToProcess[i]

			if pd.destinationName == parentConjunction.name && pd.pulseType == high {
				destinationsToTrack[pd.sourceName] = count
			}

			destinationModule := moduleMap[pd.destinationName]
			nextPulses, updatedModule := destinationModule.handlePulse(pd)
			pulseDestinationsToProcess = append(pulseDestinationsToProcess, nextPulses...)
			moduleMap[pd.destinationName] = updatedModule
		}

		// check if we have a count now for all the ones we care about
		completeCount := 0
		var iterationCounts []int
		for _, destCount := range destinationsToTrack {
			// expectation here is that its cyclical
			if destCount != 0 {
				completeCount += 1
				iterationCounts = append(iterationCounts, destCount)
			}
		}

		if completeCount == len(destinationsToTrack) {
			lcm := util.GetLeastCommonMultiple(iterationCounts)
			fmt.Printf("PART 2 %d\n", lcm)
			return
		}
	}
}

func (m Module) handlePulse(pulseData PulseData) ([]PulseData, Module) {
	//fmt.Printf("handling pulse %s -%d-> %s \n", pulseData.sourceName, pulseData.pulseType, pulseData.destinationName)

	var generatedPulses []PulseData
	switch m.moduleType {
	case broadcast:
		generatedPulses = generatePulses(pulseData.pulseType, m)
	case flipFlop:
		if pulseData.pulseType == low {
			// ok flip
			if m.flipFlopState == false {
				m.flipFlopState = true
				generatedPulses = generatePulses(high, m)
			} else {
				m.flipFlopState = false
				generatedPulses = generatePulses(low, m)
			}
		}
	case conjunction:
		m.conjunctionModuleHistory[pulseData.sourceName] = pulseData.pulseType

		foundLow := false
		for _, pulseType := range m.conjunctionModuleHistory {
			if pulseType == low {
				foundLow = true
				break
			}
		}

		if foundLow {
			generatedPulses = generatePulses(high, m)
		} else {
			generatedPulses = generatePulses(low, m)
		}
	default:
		panic("UnSupported Module Type")
	}

	return generatedPulses, m
}

func generatePulses(pulseType PulseType, m Module) []PulseData {
	// forward same pulse to all destinations
	var generatedPulses []PulseData
	for _, destinationModuleName := range m.destinationModuleNames {
		generatedPulses = append(generatedPulses, PulseData{
			pulseType:       pulseType,
			destinationName: destinationModuleName,
			sourceName:      m.name},
		)
	}
	return generatedPulses
}

func part1(lines []string) int {

	lowCount, highCount := 0, 0
	moduleMap := parseInput(lines)

	for i := 0; i < 1000; i++ {
		var pulseDestinationsToProcess []PulseData
		pulseDestinationsToProcess = append(pulseDestinationsToProcess, PulseData{"broadcaster", low, "button"})

		for i := 0; i < len(pulseDestinationsToProcess); i++ {
			pd := pulseDestinationsToProcess[i]

			if pd.pulseType == low {
				lowCount += 1
			} else {
				highCount += 1
			}
			destinationModule := moduleMap[pd.destinationName]
			nextPulses, updatedModule := destinationModule.handlePulse(pd)
			pulseDestinationsToProcess = append(pulseDestinationsToProcess, nextPulses...)

			moduleMap[pd.destinationName] = updatedModule

		}

	}
	fmt.Printf("Part 1, lowCount : %d, highCount: %d, Product: %d \n", lowCount, highCount, lowCount*highCount)
	return lowCount * highCount
}

func parseInput(lines []string) map[string]Module {
	moduleMap := map[string]Module{}

	var parsedModule Module
	for _, line := range lines {
		splitLine := strings.Split(line, " -> ")
		destinationModulesNames := strings.Split(splitLine[1], ",")

		for i := range destinationModulesNames {
			destinationModulesNames[i] = strings.TrimSpace(destinationModulesNames[i])
		}

		if splitLine[0] == "broadcaster" {
			parsedModule = Module{name: "broadcaster", moduleType: broadcast, destinationModuleNames: destinationModulesNames}
		} else if strings.HasPrefix(splitLine[0], "%") {
			parsedModule = Module{name: splitLine[0][1:], moduleType: flipFlop, destinationModuleNames: destinationModulesNames, flipFlopState: false}
		} else if strings.HasPrefix(splitLine[0], "&") {
			parsedModule = Module{name: splitLine[0][1:], moduleType: conjunction, destinationModuleNames: destinationModulesNames, conjunctionModuleHistory: map[string]PulseType{}}
		}

		moduleMap[parsedModule.name] = parsedModule

		// ok now we need to set the conjunction connected modules.
		for name, module := range moduleMap {
			for _, destinationName := range module.destinationModuleNames {
				destinationModule := moduleMap[destinationName]
				if destinationModule.moduleType == conjunction {
					destinationModule.conjunctionModuleHistory[name] = low //default of low.
				}
			}
		}
	}

	return moduleMap
}
