package main

import (
	"fmt"
	"github.com/trevmoody/aoc23/util"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2(*util.GetFileAsLines("testinput"))
}

type Range struct {
	start int
	end   int
}

func part2(lines []string) int {

	// read seeds
	seedRanges := getSeedRanges(lines[0])

	rangesToProcess := []Range{}
	unmappedForNextFilter := []Range{}

	// to start the loop
	mappedRangesforThisMappingType := seedRanges

	for lineNo := 1; lineNo < len(lines); lineNo++ {
		line := lines[lineNo]
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 2 {
			// ok this is now a new map type, so any mapped, and unmapped become the new set of ranges.
			rangesToProcess = []Range{}

			//so copy the whole lot (previously mapped, and unmapped as they have a default of no mapping) to unmapped so it is used next iteration.
			unmappedForNextFilter = append(unmappedForNextFilter, mappedRangesforThisMappingType...)
			fmt.Printf("Got Next Map Type Map Name %s, processing ranges of size %d:: %v\n\n", fields[0], len(unmappedForNextFilter), unmappedForNextFilter)

			mappedRangesforThisMappingType = []Range{}
		} else {
			fmt.Printf("Got Next Filter for Type Map Name %s, processing ranges of size %d:: %v\n\n", fields[0], len(unmappedForNextFilter), unmappedForNextFilter)

			//copy the unmapped and reset it so we can collect again for the next filter.
			rangesToProcess = []Range{}
			rangesToProcess = append(unmappedForNextFilter)
			unmappedForNextFilter = []Range{}

			var fields []int
			for _, field := range strings.Fields(line) {
				fieldInt, _ := strconv.Atoi(strings.TrimSpace(field))
				fields = append(fields, fieldInt)
			}

			mapFromRange := Range{fields[1], fields[1] + fields[2]}
			mappingDiff := fields[0] - fields[1]

			for _, rangeToProcess := range rangesToProcess {
				ok, mappedRange, unmappedRanges := mapRange(rangeToProcess, mapFromRange, mappingDiff)
				if ok {
					mappedRangesforThisMappingType = append(mappedRangesforThisMappingType, mappedRange)
				}
				unmappedForNextFilter = append(unmappedForNextFilter, unmappedRanges...)
			}
		}
	}

	//copy the whole lot to unmapped
	unmappedForNextFilter = append(unmappedForNextFilter, mappedRangesforThisMappingType...)

	//get the starts of the last set of mapped ranges
	var starts []int
	for _, r := range unmappedForNextFilter {
		starts = append(starts, r.start)
	}

	//get the min.
	minInRanges := slices.Min(starts)
	fmt.Printf("Min Location is %d\n", minInRanges)
	return minInRanges
}

func mapRange(currentRange Range, mapFromRange Range, mappingDiff int) (bool, Range, []Range) {

	// each range, is either entirely mapped to the next type, or partially mapped, or not at all.
	// if partially mapped, there will be a new range created for the next mapping type AND some new smaller range(s) for the current type that also need
	// to be checked for subsequent mappings of this mapping type
	if mapFromRange.start <= currentRange.end && mapFromRange.end >= currentRange.start {

		rangeToMap := Range{
			start: int(math.Max(float64(currentRange.start), float64(mapFromRange.start))),
			end:   int(math.Min(float64(currentRange.end), float64(mapFromRange.end))),
		}

		mappedRange := Range{
			start: rangeToMap.start + mappingDiff,
			end:   rangeToMap.end + mappingDiff,
		}

		var unmappedRanges []Range

		//unchanged parts
		if currentRange.start < rangeToMap.start {
			unmappedRanges = append(unmappedRanges, Range{currentRange.start, rangeToMap.start - 1})
		}
		if currentRange.end > rangeToMap.end {
			unmappedRanges = append(unmappedRanges, Range{rangeToMap.end, currentRange.end - 1})
		}

		return true, mappedRange, unmappedRanges
	}

	// nothing changed, pass back the raw input
	return false, Range{}, []Range{currentRange}
}

func getSeedRanges(line string) []Range {
	var seedRanges []Range
	seedFields := strings.Fields(strings.Split(line, ":")[1])
	for i := 0; i < len(seedFields); i += 2 {
		start, _ := strconv.Atoi(strings.TrimSpace(seedFields[i]))
		rangeSize, _ := strconv.Atoi(strings.TrimSpace(seedFields[i+1]))
		seedRanges = append(seedRanges, Range{start: start, end: start + rangeSize})

	}

	fmt.Printf("Got Seeds Ranges  %v\n", seedRanges)
	return seedRanges
}
