package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var numberWords map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// getFirstNumIndex returns the index of the first number in a string
func getFirstNumIndex(line []byte) (int, int) {
	for i, char := range line {
		if char >= '0' && char <= '9' {
			return i, int(char - '0')
		}
	}
	return -1, -1
}

// getLastNumIndex returns the index of the last number in a string
func getLastNumIndex(line []byte) (int, int) {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return i, int(line[i] - '0')
		}
	}
	return -1, -1
}

// getFirstWordNumIndex returns the earliest index of the given words in a string
func getFirstWordNumIndex(line []byte) (int, string) {
	var earliestIndex int
	var earliestWord string
	for word := range numberWords {
		index := bytes.Index(line, []byte(word))
		if index != -1 && (earliestWord == "" || index < earliestIndex) {
			earliestWord = word
			earliestIndex = index
		}
	}
	if earliestWord != "" {
		return earliestIndex, earliestWord
	}
	return -1, ""
}

// getLastWordNumIndex returns the latest index of the given words in a string
func getLastWordNumIndex(line []byte) (int, string) {
	var latestIndex int
	var latestWord string
	for word := range numberWords {
		index := bytes.LastIndex(line, []byte(word))
		if index != -1 && index > latestIndex {
			latestWord = word
			latestIndex = index
		}
	}
	if latestWord != "" {
		return latestIndex, latestWord
	}
	return -1, ""
}

func main() {
	// read the input file
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// read input file line by line and extract only the numbers from each line
	var total int
	for _, line := range bytes.Split(buf, []byte("\n")) {
		var numbers []string

		// check whether first match is a number or a word
		numIndex, num := getFirstNumIndex(line)
		wordIndex, word := getFirstWordNumIndex(line)

		// move on if nothing is there
		if numIndex == -1 && wordIndex == -1 {
			continue
		}

		if (numIndex != -1 && numIndex < wordIndex) || wordIndex == -1 {
			// first match is a number
			numbers = append(numbers, strconv.Itoa(num))
		} else {
			// first match is a word
			numbers = append(numbers, strconv.Itoa(numberWords[word]))
		}

		// check whether last match is a number or a word (or nothing)
		numIndex, num = getLastNumIndex(line)
		wordIndex, word = getLastWordNumIndex(line)
		if (numIndex != -1 && numIndex > wordIndex) || wordIndex == -1 {
			// last match is a number
			numbers = append(numbers, strconv.Itoa(num))
		} else {
			// last match is a word
			numbers = append(numbers, strconv.Itoa(numberWords[word]))
		}

		// combine all found numbers into one single number
		var combined int
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			combined = int(combined)*10 + num
		}

		// add the combined number to the total
		total += combined
	}

	// print the total
	fmt.Println(total)
}
