package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var grid [][]rune
var gears []*Gear

type Gear struct {
	// coords where the gear is located
	x int
	y int
}

// isSymbol checks if the rune is non-alphanumeric and not a dot
func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

// getWholeNumber returns a whole number from any character in the number
func getWholeNumber(x, y int) (startIdx, endIdx, wholeNumber int) {
	row := grid[y]

	// find the start of the number
	for i := x; i >= 0; i-- {
		if !unicode.IsDigit(row[i]) {
			startIdx = i + 1
			break
		}
	}

	// build the rest of the number
	var number []rune
	for i := startIdx; i < len(row); i++ {
		if !unicode.IsDigit(row[i]) {
			endIdx = i - 1
			break
		}
		number = append(number, row[i])
	}

	// convert the number to an int
	wholeNumber, _ = strconv.Atoi(string(number))
	return startIdx, endIdx, wholeNumber
}

// multiplySurroundingNumbers finds the numbers surrounding the gear and multiplies them
func multiplySurroundingNumbers(gear *Gear) int {
	var numbers []int

	start_x := gear.x - 1
	if start_x < 0 {
		start_x = 0
	}
	end_x := gear.x + 1
	if end_x > len(grid[gear.y])-1 {
		end_x = len(grid[gear.y]) - 1
	}

	// check the row above if it exists
	if gear.y > 0 {
		existingIdx := -1
		rowAbove := grid[gear.y-1]
		// check the cols above and see if it's a number
		for i := start_x; i <= end_x; i++ {
			if unicode.IsDigit(rowAbove[i]) {
				startIdx, _, wholeNumber := getWholeNumber(i, gear.y-1)
				if existingIdx != startIdx {
					existingIdx = startIdx
					numbers = append(numbers, wholeNumber)
				}
			}
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	// check to the left of the gear
	if gear.x > 0 {
		// check the col to the left and see if it's a number
		if unicode.IsDigit(grid[gear.y][gear.x-1]) {
			_, _, wholeNumber := getWholeNumber(gear.x-1, gear.y)
			numbers = append(numbers, wholeNumber)
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	// check to the right of the gear
	if gear.x < len(grid[gear.y])-1 {
		// check the col to the right and see if it's a number
		if unicode.IsDigit(grid[gear.y][gear.x+1]) {
			_, _, wholeNumber := getWholeNumber(gear.x+1, gear.y)
			numbers = append(numbers, wholeNumber)
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	// check the row below if it exists
	if gear.y < len(grid)-1 {
		existingIdx := -1
		rowBelow := grid[gear.y+1]
		// check the cols below and see if it's a number
		for i := start_x; i <= end_x; i++ {
			if unicode.IsDigit(rowBelow[i]) {
				startIdx, _, wholeNumber := getWholeNumber(i, gear.y+1)
				if existingIdx != startIdx {
					existingIdx = startIdx
					numbers = append(numbers, wholeNumber)
				}
			}
		}
	}

	if len(numbers) == 2 {
		// print gear coords and surrounding numbers
		fmt.Printf("gear at %d,%d has numbers %d and %d\n", gear.x, gear.y, numbers[0], numbers[1])
		return numbers[0] * numbers[1]
	}
	return 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file to read")
		os.Exit(1)
	}

	// read file from arg
	buf, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	var totalPartNumbers int

	// read each line into rows and cols
	var row []rune
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		// add the whole line to the grid
		if len(line) == 0 {
			continue
		}
		// append each char individually
		for _, chr := range line {
			row = append(row, rune(chr))
		}
		// add the row to the grid
		grid = append(grid, row)
		row = []rune{}
	}

	// iterate through the grid to look for gears
	for i, row := range grid {
		gear := &Gear{}
		for j, col := range row {
			if col == '*' {
				gear.x = j
				gear.y = i
				gears = append(gears, gear)
			}
		}
	}

	// iterate over the gears to find its surrounding numbers
	for _, gear := range gears {
		totalPartNumbers += multiplySurroundingNumbers(gear)
	}
	fmt.Printf("total is %d\n", totalPartNumbers)
}
