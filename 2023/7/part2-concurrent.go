package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var CardValues = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Original  string
	Type      HandType
	Bid       int
	Cards     []int // sorted
	OrigCards []int // original order
}

func whatHand(cards []int) HandType {
	counts := make(map[int]int)
	jokerCount := 0

	// count our cards
	for _, card := range cards {
		if card == CardValues['J'] {
			jokerCount++
			continue
		}
		counts[card]++
	}

	// find the card with the highest frequency
	var maxFreq int
	var maxFreqCard int
	for card, freq := range counts {
		if freq > maxFreq {
			maxFreq = freq
			maxFreqCard = card
		}
	}

	// increment the highest frequency by the number of jokers
	counts[maxFreqCard] += jokerCount

	switch len(counts) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		if hasValue(counts, 3) {
			return ThreeOfAKind
		}
		return TwoPair
	case 2:
		if hasValue(counts, 4) {
			return FourOfAKind
		}
		return FullHouse
	default:
		return FiveOfAKind
	}
}

func hasValue(m map[int]int, v int) bool {
	for _, val := range m {
		if val == v {
			return true
		}
	}
	return false
}

func parseHand(hand string, bid int) Hand {
	sortedCards := make([]int, len(hand))
	origCards := make([]int, len(hand))
	for i, card := range hand {
		sortedCards[i] = CardValues[card]
		origCards[i] = CardValues[card]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedCards)))
	handType := whatHand(sortedCards)

	return Hand{
		Original:  hand,
		Type:      handType,
		Bid:       bid,
		Cards:     sortedCards,
		OrigCards: origCards,
	}
}

// for sorting hands by their type and value
type byHandType []Hand

// for working with the hands
func (h byHandType) Len() int      { return len(h) }
func (h byHandType) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h byHandType) Less(i, j int) bool {
	if h[i].Type != h[j].Type {
		return h[i].Type < h[j].Type
	}
	for k := 0; k < len(h[i].OrigCards); k++ {
		if h[i].OrigCards[k] != h[j].OrigCards[k] {
			return h[i].OrigCards[k] < h[j].OrigCards[k]
		}
	}
	return false
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

	var wg sync.WaitGroup
	results := make(chan Hand)

	for _, line := range bytes.Split(buf, []byte("\n")) {
		line := string(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		hand := parts[0]

		wg.Add(1)
		go func(handStr string, bid int) {
			defer wg.Done()
			hand := parseHand(parts[0], bid)
			results <- hand
		}(hand, bid)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var hands []Hand
	for hand := range results {
		hands = append(hands, hand)
	}

	// sort the hands
	sort.Sort(byHandType(hands))

	for rank, hand := range hands {
		fmt.Printf("Rank: %d, Bid: %d, Hand: %s\n", rank+1, hand.Bid, hand.Original)
	}

	total := 0
	for rank, hand := range hands {
		total += hand.Bid * (rank + 1)
	}

	fmt.Println("Total Winnings:", total)
}
