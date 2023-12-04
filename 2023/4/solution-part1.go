package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	idx            int
	winningNumbers []int
	yourNumbers    []int
	score          int
}

// matchingNumbers returns the numbers that match winning numbers
func (c *Card) matchingNumbers() []int {
	var matches []int
	for _, n := range c.winningNumbers {
		for _, m := range c.yourNumbers {
			if n == m {
				matches = append(matches, n)
			}
		}
	}
	return matches
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

	// map of cards
	cards := make(map[int]*Card)

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// get card index
		card := Card{}
		// split by colon
		cardParts := strings.Split(line, ":")

		// get the index
		re := regexp.MustCompile(`\d+`)
		card.idx, _ = strconv.Atoi(re.FindString(cardParts[0]))

		// split by pipe symbol
		numbers := strings.Split(cardParts[1], "|")

		// get winning numbers
		for _, n := range re.FindAllString(numbers[0], -1) {
			num, _ := strconv.Atoi(n)
			card.winningNumbers = append(card.winningNumbers, num)
		}

		// get your numbers
		for _, n := range re.FindAllString(numbers[1], -1) {
			num, _ := strconv.Atoi(n)
			card.yourNumbers = append(card.yourNumbers, num)
		}

		// add card to map
		cards[card.idx] = &card
	}

	// get all the matching numbers for each card
	var total int
	for _, c := range cards {
		matches := c.matchingNumbers()
		//
		if len(matches) == 0 {
			c.score = 0
		} else {
			// score is 2^(N-1) where N is the number of matching numbers
			c.score = int(math.Pow(2, float64(len(matches)-1)))
		}

		fmt.Printf("Card %d: %d\n", c.idx, c.score)
		total += c.score
	}

	fmt.Printf("Total: %d\n", total)
}
