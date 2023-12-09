package lib

func SliceIncrementRange(start, increment int) []int {
	var rangeSlice []int
	for i := start; i < start+increment; i++ {
		rangeSlice = append(rangeSlice, i)
	}
	return rangeSlice
}

func ReverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
