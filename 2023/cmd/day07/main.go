package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type Hand struct {
	Cards     string
	Bid       int
	TypeName  string
	TypeValue int
}

func main() {
	startTime := time.Now()
	input, err := lib.GetInput("sample.txt", false)
	if err != nil {
		fmt.Println(err)
	}

	hands := make([]Hand, len(input.StringLines))

	totalsPuzzle1 := 0
	totalsPuzzle2 := 0

	getHands(input.StringLines, hands)
	// sortedHands := sortHands(hands)
	for i := range hands {
		determineHand(&hands[i])
	}

	fmt.Printf("Puzzle 1: %d\n", totalsPuzzle1)
	fmt.Printf("Puzzle 2: %d\n", totalsPuzzle2)

	duration := time.Since(startTime)
	fmt.Println("Execution time: ", duration)
}

// func sortHands(hands []Hand) []Hand {
// 	for i, hand := range hands {

// 	}
// }

func getHands(input []string, hands []Hand) {
	for i, line := range input {
		s := strings.Fields(line)
		intVal, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		hands[i] = Hand{
			Cards: s[0],
			Bid:   intVal,
		}
	}
}

func determineHand(hand *Hand) {
	// Count frequencies
	freq := make(map[rune]int)
	for _, letter := range hand.Cards {
		freq[letter]++
	}

	// Extract and sort frequencies
	var freqs []int
	for _, count := range freq {
		freqs = append(freqs, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))

	// Determine hand type
	switch freqs[0] {
	case 5:
		hand.TypeName = "Five of a Kind"
		hand.TypeValue = 7
	case 4:
		hand.TypeName = "Four of a Kind"
		hand.TypeValue = 6
	case 3:
		if freqs[1] == 2 {
			hand.TypeName = "Full House"
			hand.TypeValue = 5
		} else {
			hand.TypeName = "Three of a Kind"
			hand.TypeValue = 4
		}
	case 2:
		if freqs[1] == 2 {
			hand.TypeName = "Two Pair"
			hand.TypeValue = 3
		} else {
			hand.TypeName = "One Pair"
			hand.TypeValue = 2
		}
	default:
		hand.TypeName = "High Card"
		hand.TypeValue = 2
	}
}
