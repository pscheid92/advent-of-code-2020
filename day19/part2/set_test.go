package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSet_Add(t *testing.T) {
	s := NewStringSet()
	s.Add("1")
	s.Add("2")
	assert.Len(t, s.data, 2)
}

func TestStringSet_Union(t *testing.T) {
	s := NewStringSet()
	s.Add("a")
	s.Add("b")
	assert.Equal(t, 2, s.Len())

	// build an empty set and add it
	x := NewStringSet()
	s.Union(x)
	assert.Equal(t, 2, s.Len())

	// build a set with other elements and add it
	y := NewStringSet()
	y.Add("b")	// overlapping element
	y.Add("c")
	y.Add("d")
	s.Union(y)
	assert.Equal(t, 4, s.Len())
}

func TestStringSet_Contains(t *testing.T) {
	s := NewStringSet()
	assert.False(t, s.Contains("x"))

	s.Add("x")
	assert.True(t, s.Contains("x"))
}

func TestStringSet_GetAll(t *testing.T) {
	s := NewStringSet()

	// an empty set contains nothing
	all := s.GetAll()
	assert.Len(t, all, 0)
	assert.Equal(t, []string{}, all)

	// add some elements
	s.Add("foo")
	s.Add("bar")

	// now we can see them
	all = s.GetAll()
	assert.Len(t, all, 2)
	assert.ElementsMatch(t, []string{"foo", "bar"}, all)
}

func TestStringSet_Len(t *testing.T) {
	s := NewStringSet()
	assert.Equal(t, 0, s.Len())

	s.Add("a")
	assert.Equal(t, 1, s.Len())

	s.Add("b")
	assert.Equal(t, 2, s.Len())
}
