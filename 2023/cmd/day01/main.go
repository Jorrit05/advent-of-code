package main

import (
	"regexp"
	"strconv"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

var digitMap = map[string]string{
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

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	for _, word := range input.Lines {
		puzzle1Numbers := getNumbersOnline(word)
		puzzle1Res += extractValue(puzzle1Numbers)
	}

	for _, word := range input.StringLines {
		puzzle2Numbers := getPuzzleTwoNumbers(word)
		puzzle2Res += extractValue(puzzle2Numbers)
	}
}

func getPuzzleTwoNumbers(word string) []int {
	pattern := "\\d|one|two|three|four|five|six|seven|eight|nine"
	reversePattern := "\\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno"
	r := regexp.MustCompile(pattern)
	reverserR := regexp.MustCompile(reversePattern)

	first := string(r.FindAll([]byte(word), -1)[0])
	last := reverseString(string(reverserR.FindAll([]byte(reverseString(word)), -1)[0]))

	first = getDigitValue(first)
	last = getDigitValue(last)

	fst, _ := strconv.Atoi(first)
	lst, _ := strconv.Atoi(last)
	return []int{fst, lst}
}

func getDigitValue(key string) string {
	if len(key) > 1 {
		return digitMap[key]
	}
	return key
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getNumbersOnline(word []string) []int {
	var numbersOnLine []int
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
