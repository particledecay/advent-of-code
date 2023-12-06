package main

import (
	"bytes"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maps map[string][]*mapRange
var sourceSeeds []*seedRange

type mapRange struct {
	start, end int
	diff       int
}

type seedRange struct {
	start, end int
}

func atoi(s string) int {
	str, _ := strconv.Atoi(s)
	return str
}

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

	var currentMap string
	for i, line := range bytes.Split(buf, []byte("\n")) {
		line := string(line)
		if len(line) == 0 {
			continue
		}

		// seed line
		if i == 0 {
			maps = make(map[string][]*mapRange)
			sourceSeeds = make([]*seedRange, 0)

			// split line into pairs of seed source and range
			seedSplit := strings.Split(line, ":")
			seedValues := strings.Split(strings.Trim(seedSplit[1], " "), " ")
			for j := 0; j < len(seedValues); j += 2 {
				seedStart := atoi(seedValues[j])
				seedEnd := seedStart + (atoi(seedValues[j+1]) - 1)
				seed := &seedRange{
					start: seedStart,
					end:   seedEnd,
				}
				sourceSeeds = append(sourceSeeds, seed)
			}
			continue
		}

		// map name
		reName := regexp.MustCompile(`[a-z]+-to-([a-z]+) map:`)
		if reName.MatchString(line) {
			currentMap = reName.FindStringSubmatch(line)[1]
			maps[currentMap] = make([]*mapRange, 0)
			continue
		}

		// map ranges
		rangeSplit := strings.Split(line, " ")
		srcStart := atoi(rangeSplit[1])
		srcEnd := srcStart + (atoi(rangeSplit[2]) - 1)
		mapp := &mapRange{
			start: srcStart,
			end:   srcEnd,
			diff:  atoi(rangeSplit[0]) - srcStart,
		}
		maps[currentMap] = append(maps[currentMap], mapp)
	}

	// iterate over all maps
	for _, mapRanges := range maps {
		seeds := list.New()
		for _, seed := range sourceSeeds {
			seeds.PushBack(seed)
		}

		for seed := seeds.Front(); seed != nil; seed = seed.Next() {
			for _, mapp := range mapRanges {
				// check if seed is in range
				thisSeed := seed.Value.(*seedRange)
				if thisSeed.start >= mapp.start && thisSeed.start <= mapp.end {
					// matched, but maybe split up to make the math work
					if thisSeed.end > mapp.end {
						// split up seed
						newSeed := &seedRange{
							start: mapp.end + 1,
							end:   thisSeed.end,
						}
						thisSeed.end = mapp.end
						sourceSeeds = append(sourceSeeds, newSeed)
						seeds.PushBack(newSeed)
					}
					// since we matched, modify the seed for the next map
					thisSeed.start += mapp.diff
					thisSeed.end += mapp.diff
					break
				}
			}
		}
	}

	// find the lowest remaining seed
	lowest := -1
	for _, seed := range sourceSeeds {
		if lowest == -1 || seed.start < lowest {
			lowest = seed.start
		}
	}
	fmt.Println(lowest)
}
