package main

import (
	"fmt"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	startTime, input := lib.Init()
	defer lib.Close(startTime)
	puzzle1Res := 0
	puzzle2Res := 0

	routes := make(map[string]lib.Pair[lib.StringCompare], len(input.StringLines))

	getRoutes(input, routes)

	fmt.Printf("Puzzle 1: %d\n", puzzle1Res)
	fmt.Printf("Puzzle 2: %d\n", puzzle2Res)
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
		fmt.Println(fields)
		fmt.Println(fst)
	}
}
