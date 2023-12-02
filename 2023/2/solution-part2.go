package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handful struct {
	red   int
	green int
	blue  int
}

type Game struct {
	index      int
	hands      []Handful
	isPossible bool
}

// minimumCubePower determines the minimum possible cubes for each color and multiplies them together
func minimumCubePower(line []byte) int {
	// split on colon
	parts := strings.Split(string(line), ":")

	// second part is the handfuls
	var maxCubes map[string]int = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
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

			// update the max cubes if this one has higher count
			if count > maxCubes[color] {
				maxCubes[color] = count
			}
		}
	}

	// multiply the max cubes together
	totalPower := 1
	for _, count := range maxCubes {
		totalPower *= count
	}

	return totalPower
}

func main() {
	// read from the input file
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var totalPower int

	// for each line in the input file, evaluate a game
	for _, line := range bytes.Split(buf, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		// calculate the total power for each game
		gamePower := minimumCubePower(line)
		totalPower += gamePower
	}

	fmt.Println(totalPower)
}
