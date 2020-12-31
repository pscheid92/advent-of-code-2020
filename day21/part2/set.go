package main

import (
	"sort"
	"strings"
)

var empty struct{}

type StringSet struct {
	data map[string]struct{}
}

func NewStringSet() StringSet {
	return StringSet{make(map[string]struct{})}
}

func (s StringSet) Add(x string) {
	s.data[x] = empty
}

func (s StringSet) AddAll(items ...string) {
	for _, i := range items {
		s.data[i] = empty
	}
}

func (s StringSet) Remove(x string) {
	delete(s.data, x)
}

func (s StringSet) Union(set StringSet) {
	for _, x := range set.GetAll() {
		s.Add(x)
	}
}

func (s StringSet) Intersect(set StringSet) {
	for _, x := range s.GetAll() {
		if !set.Contains(x) {
			delete(s.data, x)
		}
	}
}

func (s StringSet) Contains(x string) bool {
	_, found := s.data[x]
	return found
}

func (s StringSet) GetAll() []string {
	elements := make([]string, 0, len(s.data))
	for x := range s.data {
		elements = append(elements, x)
	}
	return elements
}

func (s StringSet) Len() int {
	return len(s.data)
}

func (s StringSet) String() string {
	// gather keys and sort them
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// prepare builder
	builder := strings.Builder{}
	builder.WriteString("{")

	// append middle parts (with comma separation)
	for i := 0; i < len(keys)-1; i++ {
		builder.WriteString(keys[i])
		builder.WriteRune(',')
	}

	// write last element if there is one
	if len(keys)-1 >= 0 {
		builder.WriteString(keys[len(keys)-1])
	}

	// end string and return it
	builder.WriteString("}")
	return builder.String()
}
