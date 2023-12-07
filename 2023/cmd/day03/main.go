package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

var allowedChars = map[string]bool{
	".": true,
	"0": true,
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

func main() {
	input, err := lib.GetInput("input.txt", false)
	if err != nil {
		fmt.Println(err)
	}

	numberIndexSlice := getIndexesOfNumbers(*input)
	totalsPuzzle1 := getValidEngineParts(numberIndexSlice, input)
	totalsPuzzle2 := getCorrectGears(numberIndexSlice, input)

	fmt.Printf("Puzzle 1: %d\n", totalsPuzzle1)
	fmt.Printf("Puzzle 2: %d\n", totalsPuzzle2)
}

func getValidEngineParts(numberIndexSlice [][][]int, input *lib.PuzzleInput) int {
	result := 0
	for lineNumber, lineOfIndexes := range numberIndexSlice {
		for _, singleIndexSlice := range lineOfIndexes {
			result += checkNeighbours(lineNumber, singleIndexSlice, input.Lines)
		}
	}
	return result
}

func getCorrectGears(numberIndexSlice [][][]int, input *lib.PuzzleInput) int {
	resultMap := make(map[string][]int)

	for lineNumber, lineOfIndexes := range numberIndexSlice {
		for _, singleIndexSlice := range lineOfIndexes {
			checkBelowForGear(lineNumber, singleIndexSlice, input.Lines, resultMap)
		}
	}

	result := 0
	for _, v := range resultMap {
		if len(v) == 2 {
			result += v[0] * v[1]
		}
	}
	return result
}

func checkNeighbours(lineNumber int, singleIndexSlice []int, input [][]string) int {
	if len(singleIndexSlice) == 0 {
		return 0
	}

	result := processLine(lineNumber, lineNumber, singleIndexSlice, input)

	if result > 0 {
		return result
	}

	if lineNumber > 0 {
		result = processLine(lineNumber-1, lineNumber, singleIndexSlice, input)
		if result > 0 {
			return result
		}
	}

	if lineNumber == len(input)-1 {
		return 0
	} else {
		return processLine(lineNumber+1, lineNumber, singleIndexSlice, input)
	}
}

func checkBelowForGear(lineNumber int, singleIndexSlice []int, input [][]string, resultMap map[string][]int) {
	if len(singleIndexSlice) == 0 {
		return
	}

	processLinePuzzleTwo(lineNumber, lineNumber, singleIndexSlice, input, resultMap)

	if lineNumber > 0 {
		processLinePuzzleTwo(lineNumber-1, lineNumber, singleIndexSlice, input, resultMap)
	}

	if lineNumber == len(input)-1 {
		return
	} else {
		processLinePuzzleTwo(lineNumber+1, lineNumber, singleIndexSlice, input, resultMap)
	}
}

func processLinePuzzleTwo(lineNumber int, originalLineNumber int, singleIndexSlice []int, input [][]string, resultMap map[string][]int) {
	first := singleIndexSlice[0]
	last := singleIndexSlice[1]

	if last > len(input[lineNumber])-1 {
		last = len(input[lineNumber]) - 1
	}
	if first > 0 {
		first = first - 1
	}

	for i := first; i <= last; i++ {
		if input[lineNumber][i] == "*" {
			fullNumber := getFullNumber(singleIndexSlice, input[originalLineNumber])
			mapKey := fmt.Sprintf("%d_$%d", lineNumber, i)
			resultMap[mapKey] = append(resultMap[mapKey], fullNumber)
		}
	}
}

func processLine(lineNumber int, originalLineNumber int, singleIndexSlice []int, input [][]string) int {
	first := singleIndexSlice[0]
	last := singleIndexSlice[1]

	if last > len(input[lineNumber])-1 {
		last = len(input[lineNumber]) - 1
	}
	if first > 0 {
		first = first - 1
	}
	for i := first; i <= last; i++ {
		if isSymbol(input[lineNumber][i]) {
			return getFullNumber(singleIndexSlice, input[originalLineNumber])
		}

		if input[lineNumber][i] == "*" {
			return getFullNumber(singleIndexSlice, input[originalLineNumber])
		}
	}
	return 0
}

func getFullNumber(singleIndexSlice []int, input []string) int {
	// Extract the digits between the specified indexes
	digits := input[singleIndexSlice[0]:singleIndexSlice[1]]

	// Convert the string to an integer
	number, err := strconv.Atoi(strings.Join(digits, ""))
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}

	return number
}

func getIndexesOfNumbers(input lib.PuzzleInput) [][][]int {
	re := regexp.MustCompile(`\d+`)
	var allMatches [][][]int
	for _, line := range input.StringLines {
		matches := re.FindAllIndex([]byte(line), -1)
		allMatches = append(allMatches, matches)
	}

	return allMatches
}

// Use pre-defined map for constant processing speed
func isSymbol(char string) bool {
	_, exists := allowedChars[char]
	return !exists
}
