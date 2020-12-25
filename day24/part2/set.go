package main

var empty = struct{}{}

type CoordinateSet map[Coordinate]struct{}

func NewCoordinateSet(prealloc int) CoordinateSet {
	return make(CoordinateSet, prealloc)
}

func (s CoordinateSet) Add(coordinate Coordinate, neighbors ...Coordinate) {
	s[coordinate] = empty
	for _, n := range neighbors {
		s[n] = empty
	}
}

func (s CoordinateSet) ToSlice() []Coordinate {
	result := make([]Coordinate, 0, len(s))
	for c := range s {
		result = append(result, c)
	}
	return result
}
