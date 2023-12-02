package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cubeLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// isGamePossible returns the game index if possible, or -1 if not
func isGamePossible(line []byte) int {
	// split on colon
	parts := strings.Split(string(line), ":")

	// first part is the game index
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(parts[0])
	gameIndex, err := strconv.Atoi(match)
	if err != nil {
		panic(err)
	}

	// second part is the handfuls
	hands := strings.Split(parts[1], ";")
	for _, hand := range hands {
		// split on comma
		cubes := strings.Split(hand, ",")
		for _, cube := range cubes {
			// split on space
			colorCount := strings.Split(strings.Trim(cube, " "), " ")
			color := colorCount[1]
			count, err := strconv.Atoi(colorCount[0])
			if err != nil {
				panic(err)
			}

			// check that the count is within the limits
			if count > cubeLimits[color] {
				return -1
			}
		}
	}

	return gameIndex
}

func main() {
	// read from the input file
	buf, err := os.ReadFile("sample-part1.txt")
	if err != nil {
		panic(err)
	}

	// hold the total index value
	var totalIndex int

	// for each line in the input file, evaluate a game
	for _, line := range bytes.Split(buf, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		// parse the line to determine the number of cubes in each handful in this format:
		gameIndex := isGamePossible(line)
		if gameIndex == -1 {
			continue
		}

		totalIndex += gameIndex
	}

	fmt.Println(totalIndex)
}
