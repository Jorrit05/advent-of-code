package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type Card struct {
	Numbers        []string
	WinningNumbers []string
}

func main() {
	input, err := lib.GetInput("input.txt", false)
	if err != nil {
		fmt.Println(err)
	}

	cards := getCards(*input)
	result := 0
	puzzle2Map := make(map[int]int)

	for i, card := range cards {
		currentCard := i + 1
		puzzle2Map[currentCard] += 1
		count := 0

		for _, v := range card.Numbers {
			if slices.Contains(card.WinningNumbers, v) {
				count++
			}
		}
		if count == 0 {
			continue
		} else if count == 1 {
			result += 1
		} else {
			result += 1 << (count - 1)
		}

		for i := 0; i < puzzle2Map[currentCard]; i++ {
			for i := 0; i < count; i++ {
				newCard := currentCard + 1 + i
				puzzle2Map[newCard] += 1
			}
		}
	}

	totalsPuzzle2 := 0
	for _, v := range puzzle2Map {
		totalsPuzzle2 += v
	}
	fmt.Printf("Puzzle 1: %d\n", result)
	fmt.Printf("Puzzle 2: %d\n", totalsPuzzle2)
}

func getCards(input lib.PuzzleInput) []Card {
	var cards []Card

	for _, line := range input.StringLines {
		var card Card

		splitLine := strings.Split(line, ":")
		scores := strings.Split(splitLine[1], "|")

		card.Numbers = strings.Fields(scores[0])
		card.WinningNumbers = strings.Fields(scores[1])

		cards = append(cards, card)
	}
	return cards
}
