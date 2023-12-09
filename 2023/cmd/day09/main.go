package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	startTime, input := lib.Init()
	defer lib.Close(startTime)
	puzzle1Res := 0
	puzzle2Res := 0

	environmentalReport := getNumbers(input.StringLines)
	puzzle1Res = extrapolateForecast(environmentalReport)
	// fmt.Println(environmentalReport)
	fmt.Printf("Puzzle 1: %d\n", puzzle1Res)
	fmt.Printf("Puzzle 2: %d\n", puzzle2Res)
}

type Forecast struct {
	Collection lib.IntSet
	NewLine    []int
}

func extrapolateForecast(environmentalReport [][]int) int {
	result := 0
	for _, lineHistory := range environmentalReport {
		forecasts := []*Forecast{{
			NewLine: lineHistory,
		}}
		current := lineHistory

		for i := 1; i > 0; {
			forecast := &Forecast{}

			getDifferences(current, forecast)
			if len(forecast.Collection) == 1 {
				result += calculateNewForecast(forecast, forecasts)
				break
			}
			forecasts = append(forecasts, forecast)

			current = forecast.NewLine
		}
	}
	return result
}

func calculateNewForecast(forecast *Forecast, forecasts []*Forecast) int {
	plus := forecast.NewLine[0]
	for i := len(forecasts) - 1; i >= 0; i-- {
		plus += forecasts[i].NewLine[len(forecasts[i].NewLine)-1]
	}
	return plus
}

func getDifferences(lineHistory []int, forecast *Forecast) {
	forecast.Collection = make(lib.IntSet, len(lineHistory))
	forecast.NewLine = make([]int, len(lineHistory)-1)

	for i := 0; i < len(lineHistory)-1; i++ {
		diff := lineHistory[i+1] - lineHistory[i]
		forecast.Collection.Add(diff)
		forecast.NewLine[i] = diff
	}
}

func getNumbers(s []string) [][]int {
	numbers := make([][]int, len(s))

	for i, line := range s {
		correctLine := strings.Fields(line)
		numbersLine := make([]int, len(correctLine))

		for i, numberStr := range correctLine {
			nmbr, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			numbersLine[i] = nmbr
		}
		numbers[i] = numbersLine
	}

	return numbers
}
