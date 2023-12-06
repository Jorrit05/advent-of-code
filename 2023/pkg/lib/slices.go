package lib

func SliceIncrementRange(start, increment int) []int {
	var rangeSlice []int
	for i := start; i < start+increment; i++ {
		rangeSlice = append(rangeSlice, i)
	}
	return rangeSlice
}
