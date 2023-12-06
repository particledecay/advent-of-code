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
			reTime := regexp.MustCompile(`\d+`)
			timeValues := reTime.FindAllString(line, -1)

			for _, t := range timeValues {
				time, _ := strconv.Atoi(strings.Trim(t, " "))
				times = append(times, time)
			}
		} else {
			reDist := regexp.MustCompile(`\d+`)
			distanceValues := reDist.FindAllString(line, -1)

			for _, d := range distanceValues {
				distance, _ := strconv.Atoi(strings.Trim(d, " "))
				distances = append(distances, distance)
			}
		}
	}

	// multiply them together
	total := 1
	for i := 0; i < len(times); i++ {
		total *= waysToWin(times[i], distances[i])
	}

	fmt.Println(total)
}
