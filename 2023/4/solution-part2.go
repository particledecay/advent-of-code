package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	idx            int
	winningNumbers []int
	yourNumbers    []int
}

// MatchingNumbers returns the numbers that match winning numbers
func (c *Card) MatchingNumbers() []int {
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

	// doubly linked list of cards
	cardMap := make(map[int]*Card)
	cards := list.New()

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

		// add card to map and list for later
		cardMap[card.idx] = &card
		cards.PushBack(card)
	}

	// get all the matching numbers for each card
	var totalCards int
	for c := cards.Front(); c != nil; c = c.Next() {
		totalCards += 1
		card := c.Value.(Card)
		matches := card.MatchingNumbers()

		for i := range matches {
			originalCard := cardMap[card.idx+i+1]
			card := Card{
				idx:            originalCard.idx,
				winningNumbers: originalCard.winningNumbers,
				yourNumbers:    originalCard.yourNumbers,
			}
			cards.PushBack(card)
		}
	}

	// print the total cards
	fmt.Println(totalCards)
}
