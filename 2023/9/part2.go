package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculateDifferences calculates the differences between consecutive numbers
func calculateDifferences(numbers []int) []int {
	differences := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		differences[i] = numbers[i+1] - numbers[i]
	}
	return differences
}

// isConstant checks if all numbers in the slice are the same
func isConstant(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[0] {
			return false
		}
	}
	return true
}

// findNextNumber finds the next number in the sequence
func findNextNumber(numbers []int) int {
	if len(numbers) <= 1 || isConstant(numbers) {
		return numbers[len(numbers)-1]
	}

	differences := calculateDifferences(numbers)
	nextDifference := findNextNumber(differences)
	return numbers[len(numbers)-1] + nextDifference
}

// reverseSlice reverses the order of the slice
func reverseSlice(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: $0 <input-file>")
		os.Exit(1)
	}

	// read file from arg
	buf, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	scanner := bufio.NewScanner(buf)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var numbers []int
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error reading number:", err)
				return
			}
			numbers = append(numbers, number)
		}

		reversed := reverseSlice(numbers)
		sum += findNextNumber(reversed)
	}

	fmt.Println(sum)
}
