package main

import (
	"errors"
)

var StackEmptyError = errors.New("stack is empty")

type IntStack []int

func (s *IntStack) Push(x int) {
	*s = append(*s, x)
}

func (s IntStack) Peek() (int, error) {
	if len(s) == 0 {
		return 0, StackEmptyError
	}
	return s[len(s)-1], nil
}

func (s *IntStack) Pop() (int, error) {
	if n := len(*s); n == 0 {
		return 0, StackEmptyError
	} else {
		result := (*s)[n-1]
		*s = (*s)[:n-1]
		return result, nil
	}
}

type StringStack []string

func (s *StringStack) Push(x string) {
	*s = append(*s, x)
}

func (s StringStack) Peek() (string, error) {
	if len(s) == 0 {
		return "", StackEmptyError
	}
	return s[len(s)-1], nil
}

func (s *StringStack) Pop() (string, error) {
	if n := len(*s); n == 0 {
		return "", StackEmptyError
	} else {
		result := (*s)[n-1]
		*s = (*s)[:n-1]
		return result, nil
	}
}
