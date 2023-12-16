package main

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type Lens struct {
	FocalLength  int
	Label string
}

var boxes = make(map[int]*list.List, 255)
var lenses = make(map[string]*list.Element)

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	separateStrings := strings.Split(input.StringLines[0], ",")

	for _, label := range separateStrings {
			value := 0
			value = Hash(label, value)
			puzzle1Res += value
	}

	// Puzzle 2
	for i := 0 ; i < 256 ; i++ {
		boxes[i] = list.New()
	}

	for _, label := range separateStrings {
		if strings.Contains(label, "=") {
			addLensToBox(label)
		} else {
			removeLens(strings.Split(label, "-")[0])
		}
	}

	for i := 0; i < len(boxes); i++ {
		counter := 0
		if boxes[i].Len() > 0 {
			for e := boxes[i].Front(); e != nil; e = e.Next() {
				counter++
				lens := e.Value.(Lens)
				puzzle2Res += (1 + i) * counter * lens.FocalLength
			}
		}
	}
}

func removeLens(label string) {
	lens := strings.Split(label, "-")
	hash := Hash(lens[0])
	el, ok := lenses[lens[0]]
	if ok {
		boxes[hash].Remove(el)
		delete(lenses, lens[0])
	}
}

func addLensToBox(label string) {
	lens := strings.Split(label, "=")
	hash := Hash(lens[0])
  focalLength, err := strconv.Atoi(lens[1])
	if err!= nil {
		panic(err)
	}

	currentLens := Lens{FocalLength: focalLength, Label: lens[0]}
	el, ok := lenses[lens[0]]
	if ok{
		// Lens is on the shelf
		newLens := boxes[hash].InsertAfter(currentLens, el)
		lenses[lens[0]] = newLens
		boxes[hash].Remove(el)
	} else {
		lensElement := boxes[hash].PushBack(currentLens)
		lenses[lens[0]] = lensElement
	}
}

func Hash(s string, vals ...int) int {
	value := 0
	if len(vals) > 0 {
		value = vals[0]
	}
	for _, char := range s {
		val := int(char)
		value += val
		value = value *17
		value = value % 256
	}
	return value
}
