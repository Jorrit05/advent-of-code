package main

import (
	"fmt"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
	"github.com/gammazero/deque"
)

var directions = map[string][]int{
	"N": {-1, 0},
	"E": {0, 1},
	"S": {1, 0},
	"W": {0, -1},
}
var cameFrom = map[string]int{
	"N": 2,
	"E": 3,
	"S": 0,
	"W": 1,
}
var x [][]string
var MaxLength = 0
var MaxWidth = 0

var visitedSet = lib.Set[lib.MazeLocation]{}
func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)
	for i, line := range input.Lines {
			x = append(x, make([]string, len(line)))
			copy(x[i], line)
	}

	MaxWidth = len(input.Lines[0]) -1
	MaxLength = len(input.Lines)-1

	puzzle1Res = energize(input.Lines)
	count := 0
	for _, line := range x {
		for _, char := range line {
			if char == "#" {
				count++
			}
		}
	}

	fmt.Println(count)
}

func energize(input [][]string) int {
	energizedSet := lib.Set[lib.Coordinate]{}

	start := getStartPosition(input)

	var q deque.Deque[*lib.MazeLocation]
	q.PushBack(start)
	for {
		queueLength := q.Len()

		// for i := 0; i < q.Len() ; i++{
		// 	x := q.PopFront()

		// 	fmt.Println(x)
		// 	q.PushBack(x)
		// }
		// fmt.Println("----")

		if queueLength == 0 {
			break
		}
		for i := 0; i < queueLength; i++ {
			current := q.PopFront()
			energizedSet.Add(current.Coordinate)

			next, err := getNext(current)
			if err != nil {
				continue
			}
			secondDirection := determineNewDirections(current, next, input)

			if !visitedSet.Has(*next) {
				q.PushBack(next)
			}

			if secondDirection != nil && secondDirection.Direction != ""{
				if !visitedSet.Has(*secondDirection) {
					q.PushBack(secondDirection)
				}
			}
			visitedSet.Add(*current)
			// x[current.Coordinate.Row][current.Coordinate.Col] = "#"
			// lib.PrettyPrint(x)
			// fmt.Printf("\n\n\n\n\n\n")

		}
	}
	for coord := range energizedSet {
		x[coord.Row][coord.Col] = "#"
	}
	return len(energizedSet)
}

var backslashMap = map[string]string {
	"N"  : "W",
	"E"  : "S",
	"S"  : "E",
	"W" : "N",
}

var forwardslashMap = map[string]string {
	"N"  : "E",
	"E"  : "N",
	"S"  : "W",
	"W" : "S",
}

func determineNewDirections(current *lib.MazeLocation, next *lib.MazeLocation, maze [][]string) *lib.MazeLocation{
	char := maze[next.Coordinate.Row][next.Coordinate.Col]
	copy := &lib.MazeLocation{}

	switch char {
	case ".":
			next.Direction = current.Direction
	case "\\":
			next.Direction = backslashMap[current.Direction]

	case "/":
			next.Direction = forwardslashMap[current.Direction]

	case "-":
			if current.Direction == "W" || current.Direction == "E" {
					next.Direction = current.Direction
			} else {
				next.Direction = "W"
				copy = lib.CopyMazeLocation(next)
				copy.Direction = "E"
			}

	case "|":
			if current.Direction == "N" || current.Direction == "S" {
					next.Direction = current.Direction
			} else {
					next.Direction = "N"
					copy = lib.CopyMazeLocation(next)
					copy.Direction = "S"
			}
	}

	return copy
}

var start = true
func getNext(current *lib.MazeLocation) (*lib.MazeLocation, error) {
	direction := []int{}
	if start {
		direction = []int{1, 0}
		current.Direction = "S"
		start = false
	} else {
	  direction = directions[current.Direction]
	}
	newCoord := &lib.Coordinate{}
	newCoord.Row = current.Coordinate.Row + direction[0]
	newCoord.Col = current.Coordinate.Col + direction[1]

	if newCoord.Row < 0 || newCoord.Row > MaxLength ||
			newCoord.Col < 0  || newCoord.Col >  MaxWidth {
				return &lib.MazeLocation{},fmt.Errorf("index out of range")
	}

	return &lib.MazeLocation{
		Coordinate: *newCoord,
	}, nil
}

func getStartPosition(input [][]string) *lib.MazeLocation {
	return &lib.MazeLocation{
		Coordinate: lib.Coordinate{
			Row: 0,
			Col: 0,
		},
		Direction: "E",
	}
}
