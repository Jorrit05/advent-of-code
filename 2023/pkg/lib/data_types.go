package lib

type Coordinate struct {
	Row int
	Col int
}

// type IntPair struct {
// 	Lhs, Rhs int
// }

// type SortLhs []IntPair

// func (a SortLhs) Len() int           { return len(a) }
// func (a SortLhs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortLhs) Less(i, j int) bool { return a[i].Lhs < a[j].Lhs }

// type SortRhs []IntPair

// func (a SortRhs) Len() int           { return len(a) }
// func (a SortRhs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortRhs) Less(i, j int) bool { return a[i].Rhs < a[j].Rhs }

// // Warning not safe for string number comparison: "11" < "2" = true
// type StringPair struct {
// 	Lhs, Rhs string
// }

// type SrtLhs []StringPair

// func (a SrtLhs) Len() int           { return len(a) }
// func (a SrtLhs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SrtLhs) Less(i, j int) bool { return a[i].Lhs < a[j].Lhs }

// type SrtRhs []StringPair

// func (a SrtRhs) Len() int           { return len(a) }
// func (a SrtRhs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SrtRhs) Less(i, j int) bool { return a[i].Rhs < a[j].Rhs }

type Comparer[T any] interface {
	Compare(other T) int
}

type Pair[T Comparer[T]] struct {
	Fst, Snd T
}

type Slice[T Comparer[T]] []Pair[T]

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Slice[T]) Less(i, j int) bool {
	return s[i].Fst.Compare(s[j].Fst) < 0
}

type SliceBySecond[T Comparer[T]] []Pair[T]

func (s SliceBySecond[T]) Len() int {
	return len(s)
}

func (s SliceBySecond[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SliceBySecond[T]) Less(i, j int) bool {
	return s[i].Snd.Compare(s[j].Snd) < 0
}

// Example of an integer type implementing Comparer
type IntCompare int

func (i IntCompare) Compare(other IntCompare) int {
	if i < other {
		return -1
	} else if i > other {
		return 1
	}
	return 0
}

type StringCompare string

func (i StringCompare) Compare(other StringCompare) int {
	if i < other {
		return -1
	} else if i > other {
		return 1
	}
	return 0
}

func GenerateIntPairs(rangeEnd int) []Pair[IntCompare] {
	var pairs []Pair[IntCompare]

	for i := 1; i < rangeEnd; i++ {
		for j := i + 1; j <= rangeEnd; j++ {
			pairs = append(pairs, Pair[IntCompare]{IntCompare(i), IntCompare(j)})
		}
	}

	return pairs
}
