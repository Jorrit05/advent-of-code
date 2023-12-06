package main

type Destinations struct {
	DestRange   int
	SourceRange int
	Length      int
}

type ConversionMap struct {
	SourceName      string
	DestinationName string
	Destinations    []Destinations
}

// NewDestinations creates a new Destinations struct with the given values.
func NewDestinations(destRange, sourceRange, length int) *Destinations {
	return &Destinations{
		DestRange:   destRange,
		SourceRange: sourceRange,
		Length:      length,
	}
}
