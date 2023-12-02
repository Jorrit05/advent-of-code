package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	input, err := lib.GetInput("sample.txt", false)
	if err != nil {
		fmt.Println(err)
	}

	totals := 0

	var iterator [][]string
	puzzleOne := false
	if puzzleOne {
		iterator = input.Lines
	} else {
		iterator = replaceNumberStrings(*input)

	}
	// Puzzle 1:
	for _, word := range iterator {
		var numbersOnLine []int

		numbersOnLine = getNumbersOnline(word, numbersOnLine)
		totals += extractValue(numbersOnLine)
	}
	fmt.Println(totals)
}

// Puzzle 2:
func replaceNumberStrings(input lib.PuzzleInput) [][]string {
	digitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var newlines [][]string

	for _, word := range input.StringLines {

		output := replaceFirstLastOccurrences(word, digitMap)
		fmt.Println(output)
		// Iterate over each word in input.StringLines
		// Append the final result to newlines
		lineSlice := strings.Split(output, "")
		newlines = append(newlines, lineSlice)
	}
	// Compile regular expressions for each word in the map
	// regexMap := make(map[*regexp.Regexp]string)
	// for word, digit := range digitMap {
	// 	// Regular expression pattern matching the word directly
	// 	pattern := word
	// 	r := regexp.MustCompile(pattern)
	// 	regexMap[r] = digit
	// }

	// var newlines [][]string

	// // Iterate over each word in input.StringLines
	// for _, word := range input.StringLines {
	// 	line := word

	// 	// Apply all regex replacements to the line
	// 	for regex, digit := range regexMap {
	// 		line = regex.ReplaceAllString(line, digit)
	// 	}

	// 	// Append the final result to newlines
	// 	lineSlice := strings.Split(line, "")
	// 	newlines = append(newlines, lineSlice)
	// }
	// fmt.Println(newlines)
	return newlines
}

func getNumbersOnline(word []string, numbersOnLine []int) []int {
	for _, letter := range word {
		if val, err := strconv.Atoi(letter); err == nil {
			numbersOnLine = append(numbersOnLine, val)
		}
	}
	return numbersOnLine
}

func extractValue(numbersOnLine []int) int {
	if len(numbersOnLine) == 1 {
		return calcValue(numbersOnLine[0], numbersOnLine[0])
	} else {
		return calcValue(numbersOnLine[0], numbersOnLine[len(numbersOnLine)-1])
	}
}

func calcValue(a int, b int) int {
	return (a * 10) + b
}

func replaceFirstLastOccurrences(s string, digitMap map[string]string) string {
	for word, digit := range digitMap {
		s = replaceFirstOccurrence(s, word, digit)
		s = replaceLastOccurrence(s, word, digit)
	}
	return s
}

func replaceFirstOccurrence(s, word, replacement string) string {
	idx := strings.Index(s, word)
	if idx != -1 {
		return s[:idx] + replacement + s[idx+len(word):]
	}
	return s
}

func replaceLastOccurrence(s, word, replacement string) string {
	idx := strings.LastIndex(s, word)
	if idx != -1 {
		return s[:idx] + replacement + s[idx+len(word):]
	}
	return s
}
