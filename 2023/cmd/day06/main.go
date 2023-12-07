package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type Race struct {
	Time     int
	Distance int
}

func NewRace(time, distance int) *Race {
	return &Race{
		Time:     time,
		Distance: distance,
	}
}

func main() {
	startTime := time.Now()
	input, err := lib.GetInput("input.txt", false)
	if err != nil {
		fmt.Println(err)
	}
	races, racePuzzleTwo := getRaces(input.StringLines)
	totalsPuzzle1 := 1
	totalsPuzzle2 := 1

	for _, race := range races {
		totalsPuzzle1 *= getNumberOfPossibleWins(race)
	}

	totalsPuzzle2 *= getNumberOfPossibleWins(racePuzzleTwo)
	fmt.Printf("Puzzle 1: %d\n", totalsPuzzle1)

	// races := getRaces(input.StringLines)

	fmt.Printf("Puzzle 2: %d\n", totalsPuzzle2)

	duration := time.Since(startTime)
	fmt.Println("Execution time: ", duration)
}

func getNumberOfPossibleWins(race Race) int {
	count := 0
	for i := 1; i < race.Time; i++ {
		distance := i * (race.Time - i)
		if distance > race.Distance {
			count++
		}
	}
	return count
}

func getRaces(input []string) ([]Race, Race) {
	var races []Race

	time := strings.Fields(input[0])
	distance := strings.Fields(input[1])

	for i := 1; i < len(time); i++ {
		tm, _ := strconv.Atoi(time[i])
		dist, _ := strconv.Atoi(distance[i])
		races = append(races, *NewRace(tm, dist))
	}

	raceLength := strings.Split(strings.Join(time, ""), ":")[1]
	raceDistance := strings.Split(strings.Join(distance, ""), ":")[1]
	tm, _ := strconv.Atoi(raceLength)
	dist, _ := strconv.Atoi(raceDistance)

	return races, *NewRace(tm, dist)
}
