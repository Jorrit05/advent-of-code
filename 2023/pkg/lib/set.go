package lib

// Define a set type as a map of strings to empty structs
type Set map[string]struct{}

// Add a method to add an element to the set
func (s Set) Add(element string) {
	s[element] = struct{}{}
}

// Add a method to remove an element from the set
func (s Set) Remove(element string) {
	delete(s, element)
}

// Add a method to check if an element is in the set
func (s Set) Has(element string) bool {
	_, ok := s[element]
	return ok
}

type IntSet map[int]struct{}

// Add a method to add an element to the set
func (s IntSet) Add(element int) {
	s[element] = struct{}{}
}

// Add a method to remove an element from the set
func (s IntSet) Remove(element int) {
	delete(s, element)
}

// Add a method to check if an element is in the set
func (s IntSet) Has(element int) bool {
	_, ok := s[element]
	return ok
}
