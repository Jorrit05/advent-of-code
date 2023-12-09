package main

import (
	"strconv"
	"strings"

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
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	puzzle1Res = 1
	puzzle2Res = 1
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	races, racePuzzleTwo := getRaces(input.StringLines)

	for _, race := range races {
		puzzle1Res *= getNumberOfPossibleWins(race)
	}

	puzzle2Res *= getNumberOfPossibleWins(racePuzzleTwo)
}

func getNumberOfPossibleWins(race Race) int {
	for i := 1; i < race.Time; i++ {
		distance := i * (race.Time - i)
		if distance > race.Distance {
			return (race.Time - i) - i + 1
		}
	}
	return 1
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
