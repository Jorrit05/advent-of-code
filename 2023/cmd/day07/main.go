package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type Hand struct {
	Cards     string
	Bid       int
	TypeName  string
	TypeValue int
}

type ByTypeValue []Hand

func (a ByTypeValue) Len() int      { return len(a) }
func (a ByTypeValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTypeValue) Less(i, j int) bool {
	if a[i].TypeValue == a[j].TypeValue {
		return customSort(a[i], a[j])
	}
	return a[i].TypeValue < a[j].TypeValue
}

func customSort(a, b Hand) bool {
	for i := 0; i < len(a.Cards) && i < len(b.Cards); i++ {
		if a.Cards[i] != b.Cards[i] {
			return a.Cards[i] < b.Cards[i]
		}
	}
	return false
}

func main() {
	startTime, input := lib.Init()
	defer lib.Close(startTime)

	hands := make([]Hand, len(input.StringLines))
	handsTwo := make([]Hand, len(input.StringLines))

	totalsPuzzle1 := 0
	totalsPuzzle2 := 0

	getHands(input.StringLines, hands, handsTwo)

	for i := range hands {
		determineHand(&hands[i])
		determineHandPuzzleTwo(&handsTwo[i])
	}

	for i, _ := range handsTwo {
		handsTwo[i].Cards = replaceMultiple(handsTwo[i].Cards, replacementsTwo)
	}

	sort.Sort(ByTypeValue(hands))
	sort.Sort(ByTypeValue(handsTwo))

	// Puzzle 1
	for rank, hand := range hands {
		rank++
		totalsPuzzle1 += (hand.Bid * rank)
	}
	for rank, hand := range handsTwo {
		rank++
		totalsPuzzle2 += (hand.Bid * rank)
	}
	fmt.Printf("Puzzle 1: %d\n", totalsPuzzle1)
	fmt.Printf("Puzzle 2: %d\n", totalsPuzzle2)
}

func getHands(input []string, hands []Hand, handsTwo []Hand) {
	for i, line := range input {
		s := strings.Fields(line)
		intVal, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		str := replaceMultiple(s[0], replacements)
		hands[i] = Hand{
			Cards: str,
			Bid:   intVal,
		}

		handsTwo[i] = Hand{
			Cards: s[0],
			Bid:   intVal,
		}
	}
}

var replacements = map[string]string{
	"T": "B",
	"J": "C",
	"Q": "D",
	"K": "E",
	"A": "F",
}
var replacementsTwo = map[string]string{
	"T": "B",
	"J": "1",
	"Q": "D",
	"K": "E",
	"A": "F",
}

func replaceMultiple(s string, replacements map[string]string) string {
	for oldChar, newChar := range replacements {
		s = strings.ReplaceAll(s, oldChar, newChar)
	}
	return s
}

func determineHand(hand *Hand) {
	// Count frequencies
	freq := make(map[rune]int)
	for _, letter := range hand.Cards {
		freq[letter]++
	}

	var freqs []int
	for _, count := range freq {
		freqs = append(freqs, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))

	// Determine hand type
	switch freqs[0] {
	case 5:
		hand.TypeName = "Five of a Kind"
		hand.TypeValue = 6
	case 4:
		hand.TypeName = "Four of a Kind"
		hand.TypeValue = 5
	case 3:
		if freqs[1] == 2 {
			hand.TypeName = "Full House"
			hand.TypeValue = 4
		} else {
			hand.TypeName = "Three of a Kind"
			hand.TypeValue = 3
		}
	case 2:
		if freqs[1] == 2 {
			hand.TypeName = "Two Pair"
			hand.TypeValue = 2
		} else {
			hand.TypeName = "One Pair"
			hand.TypeValue = 1
		}
	default:
		hand.TypeName = "High Card"
		hand.TypeValue = 0
	}
}

type Pair struct {
	Letter string
	Freq   int
}
type ByFreq []Pair

func (a ByFreq) Len() int           { return len(a) }
func (a ByFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFreq) Less(i, j int) bool { return a[i].Freq < a[j].Freq }

func determineHandPuzzleTwo(hand *Hand) {
	// Count frequencies
	m := make(map[string]int)
	for _, letter := range strings.Split(hand.Cards, "") {
		m[letter]++
	}
	containsJokers := strings.Contains(hand.Cards, "J")
	pairs := make([]Pair, 0, 5)
	for k, v := range m {
		pairs = append(pairs, Pair{
			Letter: k,
			Freq:   v,
		})
	}

	if containsJokers {
		sort.Sort(sort.Reverse(ByFreq(pairs)))

		if pairs[0].Letter == "J" {
			//J is highest
			if pairs[0].Freq < 5 {
				pairs[0].Freq += pairs[1].Freq
				pairs[1].Freq = 0
			}
		} else {
			for i, v := range pairs {
				if v.Letter == "J" {
					pairs[0].Freq += pairs[i].Freq
					pairs[i].Freq = 0
				}
			}
		}
	}
	sort.Sort(sort.Reverse(ByFreq(pairs)))

	// Determine hand type
	switch pairs[0].Freq {
	case 5:
		hand.TypeName = "Five of a Kind"
		hand.TypeValue = 6
	case 4:
		hand.TypeName = "Four of a Kind"
		hand.TypeValue = 5
	case 3:
		if pairs[1].Freq == 2 {
			hand.TypeName = "Full House"
			hand.TypeValue = 4
		} else {
			hand.TypeName = "Three of a Kind"
			hand.TypeValue = 3
		}
	case 2:
		if pairs[1].Freq == 2 {
			hand.TypeName = "Two Pair"
			hand.TypeValue = 2
		} else {
			hand.TypeName = "One Pair"
			hand.TypeValue = 1
		}
	default:
		hand.TypeName = "High Card"
		hand.TypeValue = 0
	}
}
