package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var times []int
var distances []int

func waysToWin(t, d int) int {
	var ways int
	for i := 1; i < t; i++ {
		if (t-i)*i > d {
			ways++
		}
	}
	return ways
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

	for _, line := range bytes.Split(buf, []byte("\n")) {
		line := string(line)
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Time") {
			reTime := regexp.MustCompile(`\d`)
			timeValues := reTime.FindAllString(line, -1)

			var combined int
			for _, t := range timeValues {
				num, _ := strconv.Atoi(t)
				combined = int(combined)*10 + num
			}
			times = append(times, combined)
		} else {
			reDist := regexp.MustCompile(`\d`)
			distanceValues := reDist.FindAllString(line, -1)

			var combined int
			for _, d := range distanceValues {
				num, _ := strconv.Atoi(d)
				combined = int(combined)*10 + num
			}
			distances = append(distances, combined)
		}
	}

	// multiply them together
	total := 1
	for i := 0; i < len(times); i++ {
		total *= waysToWin(times[i], distances[i])
	}

	fmt.Println(total)
}
