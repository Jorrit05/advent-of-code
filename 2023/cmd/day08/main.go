package main

import (
	"regexp"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	routes := make(map[string]lib.Pair[lib.StringCompare], len(input.StringLines))

	getRoutes(input, routes)
	directions := input.Lines[0]
	// startString := "AAA"
	// found := false

	// for i := 1; i > 0; {
	// 	found, puzzle1Res, startString = findZZZ(directions, startString, routes, puzzle1Res)
	// 	if found {
	// 		break
	// 	}
	// }

	startPositions := getStartPosition(routes)
	Graphs(input, startPositions, directions)

	// found = false
	// fmt.Println(startPositions)
	// for i := 1; i > 0; {
	// 	for _, direction := range directions {
	// 		// fmt.Println(startPositions)
	// 		found, puzzle2Res, startPositions = findAllxxZ(direction, startPositions, routes, puzzle2Res)
	// 		if found {
	// 			break
	// 		}
	// 	}
	// 	if found {
	// 		break
	// 	}
	// }
	// fmt.Println(startPositions)
}

func getStartPosition(routes map[string]lib.Pair[lib.StringCompare]) []string {
	var startPositions []string
	pattern := "^[A-Z0-9]{2}A"
	re := regexp.MustCompile(pattern)
	for k := range routes {
		if re.MatchString(k) {
			startPositions = append(startPositions, k)
		}
	}
	return startPositions
}

// function ConcurrentBFS(graph, startNodes)
//     create a queue Q
//     create a map M to store the state of each path (node and step count)
//     for each node in startNodes
//         add node to Q
//         add node to M with step count 0

//     while Q is not empty
//         size = length of Q
//         for i from 0 to size
//             current = Q.dequeue()
//             currentStep = M[current]

//             if current ends with 'Z' and all nodes in M end with 'Z'
//                 return currentStep

//             for each node n that is adjacent to current (based on the instruction)
//                 if n is not in M or M[n] > currentStep + 1
//                     M[n] = currentStep + 1
//                     enqueue n onto Q

//     return failure

func findAllxxZ(direction string, startPositions []string, routes map[string]lib.Pair[lib.StringCompare], puzzle2Res int) (bool, int, []string) {
	nextPositions := make([]string, 0, len(startPositions))

	for _, startPos := range startPositions {
		if direction == "L" {
			nextPositions = append(nextPositions, string(routes[startPos].Fst))
		} else {
			nextPositions = append(nextPositions, string(routes[startPos].Snd))
		}
	}
	puzzle2Res++
	if checkIfFinished(nextPositions) {
		return true, puzzle2Res, nextPositions
	}

	return false, puzzle2Res, nextPositions
}

func checkIfFinished(nextPositions []string) bool {
	for _, nextPos := range nextPositions {
		if nextPos[2] != 'Z' {
			return false
		}
	}
	return true
}

func findZZZ(directions []string, startString string, routes map[string]lib.Pair[lib.StringCompare], puzzle1Res int) (bool, int, string) {
	for _, direction := range directions {
		if direction == "L" {
			startString = string(routes[startString].Fst)
		} else {
			startString = string(routes[startString].Snd)
		}
		puzzle1Res++
		if startString == "ZZZ" {
			return true, puzzle1Res, startString
		}
	}
	return false, puzzle1Res, startString
}

func getRoutes(input *lib.PuzzleInput, routes map[string]lib.Pair[lib.StringCompare]) {
	for i := 2; i < len(input.StringLines); i++ {
		fields := strings.Fields(input.StringLines[i])
		fst := lib.StringCompare(strings.Trim(fields[2], "(,"))
		snd := lib.StringCompare(strings.Trim(fields[3], "(,)"))
		routes[fields[0]] = lib.Pair[lib.StringCompare]{
			Fst: fst,
			Snd: snd,
		}
	}
}
