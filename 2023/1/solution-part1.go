package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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
		line := string(line)
		re := regexp.MustCompile(`\d`)
		// print the first match
		matches := re.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		// append only first match and last match (even if it's the same)
		numbers = append(numbers, matches[0])
		numbers = append(numbers, matches[len(matches)-1])

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
