package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntStack_Push(t *testing.T) {
	// initially a stack is empty
	s := IntStack{}
	assert.Len(t, s, 0)

	// pushing some elements lets it grow
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Len(t, s, 3)
	assert.EqualValues(t, []int{1, 2, 3}, s)
}

func TestIntStack_Peek(t *testing.T) {
	// peeking an empty stack fails
	s := IntStack{}
	_, err := s.Peek()
	assert.Error(t, err)

	// pushing a value and peeking works
	s.Push(10)
	element1, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 10, element1)

	// peeking multiple times does not change the value
	element2, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, element1, element2)
}

func TestIntStack_Pop(t *testing.T) {
	// popping an empty stack fails
	s := IntStack{}
	_, err := s.Pop()
	assert.Error(t, err)

	// popping after an push works
	s.Push(10)
	s.Push(20)

	element, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, element)

	element, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 10, element)

	_, err = s.Pop()
	assert.Error(t, err)
}

func TestStringStack_Push(t *testing.T) {
	// initially a stack is empty
	s := StringStack{}
	assert.Len(t, s, 0)

	// pushing some elements lets it grow
	s.Push("1")
	s.Push("2")
	s.Push("3")
	assert.Len(t, s, 3)
	assert.EqualValues(t, []string{"1", "2", "3"}, s)
}

func TestStringStack_Peek(t *testing.T) {
	// peeking an empty stack fails
	s := StringStack{}
	_, err := s.Peek()
	assert.Error(t, err)

	// pushing a value and peeking works
	s.Push("foobar")
	element1, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "foobar", element1)

	// peeking multiple times does not change the value
	element2, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, element1, element2)
}

func TestStringStack_Pop(t *testing.T) {
	// popping an empty stack fails
	s := StringStack{}
	_, err := s.Pop()
	assert.Error(t, err)

	// popping after an push works
	s.Push("foo")
	s.Push("bar")

	element, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "bar", element)

	element, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "foo", element)

	_, err = s.Pop()
	assert.Error(t, err)
}