package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var grid [][]rune

type Number struct {
	// row where number is located
	row int
	// start and end index of the number
	startIdx int
	endIdx   int
	// the number itself
	number []rune
}

// Value returns the numerical value of the number as an int
func (n *Number) Value() int {
	var value int
	for _, digit := range n.number {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		value = int(value)*10 + num
	}

	return value
}

// isSymbol checks if the rune is non-alphanumeric and not a dot
func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

// checkAroundNumber checks the surrounding grid indices for any symbol that's not a dot and returns true if found
func checkAroundNumber(number *Number) bool {
	var rows [][]string

	left_index := number.startIdx - 1
	if left_index < 0 {
		left_index = 0
	}
	right_index := number.endIdx + 1
	if right_index > len(grid[number.row])-1 {
		right_index = len(grid[number.row]) - 1
	}

	// check the row above if it exists
	if number.row > 0 {
		var topRow []string
		for i := left_index; i <= right_index; i++ {
			if isSymbol(grid[number.row-1][i]) {
				return true
			}
			topRow = append(topRow, string(grid[number.row-1][i]))
		}
		rows = append(rows, topRow)
	}

	// check around the number in its row
	// check left
	var thisRow []string
	if left_index > 0 {
		if isSymbol(grid[number.row][left_index]) {
			return true
		}
		thisRow = append(thisRow, string(grid[number.row][left_index]))
	}
	for _, char := range number.number {
		thisRow = append(thisRow, string(char))
	}
	// check right
	if right_index != left_index {
		if isSymbol(grid[number.row][right_index]) {
			return true
		}
		thisRow = append(thisRow, string(grid[number.row][right_index]))
	}
	rows = append(rows, thisRow)

	// check the row below if exists
	if number.row < len(grid)-1 {
		var bottomRow []string
		for i := left_index; i <= right_index; i++ {
			if isSymbol(grid[number.row+1][i]) {
				return true
			}
			bottomRow = append(bottomRow, string(grid[number.row+1][i]))
		}
		rows = append(rows, bottomRow)
	}

	return false
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

	// iterate through the grid to look for groups of characters forming a number
	for i, row := range grid {
		var inNumber bool
		number := &Number{}
		for j, col := range row {
			if !unicode.IsDigit(col) {
				// end the previous number's index
				if inNumber {
					number.endIdx = j - 1
				}
				inNumber = false
				// if we have a number, check it
				if len(number.number) > 0 {
					valid := checkAroundNumber(number)
					if valid {
						// fmt.Printf("VALID: %s\n", string(number.number))
						totalPartNumbers += number.Value()
					} else {
						fmt.Printf("%s is not valid\n", string(number.number))
					}
				}
				number = &Number{}
			} else {
				// if col is a digit, start a new number
				if _, err := strconv.Atoi(string(col)); err == nil {
					if !inNumber {
						number.startIdx = j
						number.row = i
						inNumber = true
					}
					number.number = append(number.number, col)
				}
			}
		}
	}

	fmt.Printf("total is %d\n", totalPartNumbers)
}
