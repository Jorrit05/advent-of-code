package main

import (
	"slices"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type Card struct {
	Numbers        []string
	WinningNumbers []string
}

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	cards := getCards(*input)

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
			puzzle1Res += 1
		} else {
			puzzle1Res += 1 << (count - 1)
		}

		for i := 0; i < puzzle2Map[currentCard]; i++ {
			for i := 0; i < count; i++ {
				newCard := currentCard + 1 + i
				puzzle2Map[newCard] += 1
			}
		}
	}

	for _, v := range puzzle2Map {
		puzzle2Res += v
	}
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
