package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type SeedRange struct {
	Start     int
	Increment int
}

var locations = make(map[int]int)
var locationsTwo = make(map[int]int)
var mutex = &sync.Mutex{}
var mutex2 = &sync.Mutex{}

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	seeds := getSeeds(input.StringLines[0])
	seedPairs := getSeedPairs(input.StringLines[0])

	soilMap := make(map[string]ConversionMap)
	createSoilMap(input.StringLines, soilMap)

	for _, seed := range seeds {
		findLocations(seed, soilMap, locations, mutex)
	}
	puzzle1Res, _ = KeyWithLowestValue(locations)

	findLocationsTwo(seedPairs, soilMap, locationsTwo, mutex)
	puzzle2Res, _ = KeyWithLowestValue(locationsTwo)
}

var alreadyChecked = make(map[int]rune)

func findLocationsTwo(seedPairs []SeedRange, soilMap map[string]ConversionMap, locations map[int]int, mutex *sync.Mutex) {
	var wg sync.WaitGroup

	for _, seedPair := range seedPairs {
		wg.Add(1) // Increment the WaitGroup counter
		go func(seedPair SeedRange) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			for _, destination := range soilMap["seed"].Destinations {
				if seedPair.Start >= destination.SourceRange && seedPair.Start < destination.SourceRange+destination.Length {
					end := min(seedPair.Start+seedPair.Increment, destination.SourceRange+destination.Length)
					for i := seedPair.Start; i <= end; i++ {
						mutex2.Lock()
						if _, exists := alreadyChecked[i]; !exists {
							alreadyChecked[i] = 'i'
							mutex2.Unlock()
							findLocations(i, soilMap, locations, mutex)
						} else {
							mutex2.Unlock()
						}
					}
					fmt.Println(seedPair)
				}
			}
		}(seedPair)
	}
	fmt.Println("waiting")
	wg.Wait()
}

func getSeedPairs(s string) []SeedRange {
	var res []SeedRange

	stringArray := strings.Fields(s)
	for i := 1; i < len(stringArray); i += 2 {
		start, _ := strconv.Atoi(stringArray[i])
		increment, _ := strconv.Atoi(stringArray[i+1])

		var pair SeedRange
		pair.Start = start
		pair.Increment = increment
		res = append(res, pair)

	}
	return res
}

func findLocations(seed int, soilMap map[string]ConversionMap, locations map[int]int, mutex *sync.Mutex) {

	tmpValues := make(map[int]int)
	tmpValues[0] = seed
	dest := "seed"
	for range soilMap {
		latestVal := tmpValues[0]

		currentMap := soilMap[dest]
		newValue := -1

		for _, v := range currentMap.Destinations {
			// seed is in source range value, do some math
			if latestVal >= v.SourceRange && latestVal < v.SourceRange+v.Length {
				newValue = (v.DestRange - v.SourceRange) + latestVal
				break
			}
		}

		if newValue == -1 {
			newValue = latestVal
		}
		tmpValues[0] = newValue
		dest = currentMap.DestinationName
		if currentMap.DestinationName == "location" {
			mutex.Lock() // Lock the mutex before accessing the map
			locations[newValue] = newValue
			mutex.Unlock()
		}
	}
}

func createSoilMap(lines []string, soilMap map[string]ConversionMap) {
	currentEntry := []string{}
	conversionMap := &ConversionMap{}

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			if conversionMap.SourceName != "" {
				soilMap[currentEntry[0]] = *conversionMap
				conversionMap = &ConversionMap{}
			}
			continue
		}
		splitLine := strings.Fields(line)

		if strings.Contains(line, "map:") {
			currentEntry = strings.Split(splitLine[0], "-")
			conversionMap.SourceName = currentEntry[0]
			conversionMap.DestinationName = currentEntry[2]
			continue
		} else {
			destVals := []int{}
			for _, value := range strings.Fields(line) {
				val, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}
				destVals = append(destVals, val)
			}

			dest := NewDestinations(destVals[0], destVals[1], destVals[2])
			conversionMap.Destinations = append(conversionMap.Destinations, *dest)
		}
	}
	soilMap[currentEntry[0]] = *conversionMap
	conversionMap = &ConversionMap{}
}

func getSeeds(firstLine string) []int {
	var seeds []int
	for _, seedNr := range strings.Fields(firstLine) {
		value, err := strconv.Atoi(seedNr)
		if err != nil {
			continue
		}
		seeds = append(seeds, value)
	}
	return seeds
}
func KeyWithLowestValue(m map[int]int) (keyWithLowestValue int, found bool) {
	if len(m) == 0 {
		// If the map is empty, return that no key was found
		return 0, false
	}

	var minValue int
	for key, value := range m {
		if !found || value < minValue {
			// For the first item or if the current value is lower than the lowest so far
			minValue = value
			keyWithLowestValue = key
			found = true
		}
	}
	return keyWithLowestValue, found
}
