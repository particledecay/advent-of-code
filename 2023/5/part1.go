package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maps map[string][][]int
var mappedSeeds map[int]map[string]int
var sourceSeeds []int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: $0 <input-file>")
		os.Exit(1)
	}

	// read file from arg
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	maps = make(map[string][][]int)
	mappedSeeds = make(map[int]map[string]int)
	var currentMatch string

	for i, line := range bytes.Split(buf, []byte("\n")) {
		line := string(line)
		if len(line) == 0 {
			continue
		}

		if i == 0 { // seed line
			seeds := strings.Split(line, ":")
			seedNumbers := strings.Split(strings.Trim(seeds[1], " "), " ")
			for _, seed := range seedNumbers {
				seed, _ := strconv.Atoi(strings.Trim(seed, " "))
				sourceSeeds = append(sourceSeeds, seed)
			}
			continue
		}

		// build the maps
		re := regexp.MustCompile(`[a-z]+-[a-z]+-([a-z]+) map:`)
		matches := re.FindStringSubmatch(line)

		if len(matches) > 0 {
			currentMatch = matches[1]
			continue
		} else {
			// we're in the mapping lines
			conversion := strings.Split(line, " ")
			dest, _ := strconv.Atoi(strings.Trim(conversion[0], " "))
			source, _ := strconv.Atoi(strings.Trim(conversion[1], " "))
			times, _ := strconv.Atoi(strings.Trim(conversion[2], " "))
			maps[currentMatch] = append(maps[currentMatch], []int{dest, source, times})
		}
	}

	// iterate over seeds
	var previousMapper string
	for _, mapper := range []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"} {
		thisMap := maps[mapper]

		for _, seed := range sourceSeeds {
			if mappedSeeds[seed] == nil {
				mappedSeeds[seed] = make(map[string]int)
			}

			mappedSeeds[seed][mapper] = -1

			// get previous value or seed
			var src int
			if previousMapper == "" {
				src = seed
			} else {
				src = mappedSeeds[seed][previousMapper]
			}

			for _, mapping := range thisMap {
				diff := src - mapping[1]
				if diff > 0 && diff <= mapping[2] {
					mappedSeeds[seed][mapper] = mapping[0] + diff
				}
			}

			// if we didn't find a mapping, use the last one
			if mappedSeeds[seed][mapper] == -1 {
				mappedSeeds[seed][mapper] = src
			}
		}
		previousMapper = mapper
	}

	// find the lowest location
	lowestLocation := -1
	for _, seed := range sourceSeeds {
		if lowestLocation == -1 || mappedSeeds[seed]["location"] < lowestLocation {
			lowestLocation = mappedSeeds[seed]["location"]
		}
	}

	fmt.Println(lowestLocation)
}
