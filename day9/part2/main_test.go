package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAtIndexValid(t *testing.T) {
	preamble := 5
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	for i := preamble; i < len(numbers); i++ {
		expect := numbers[i] != 127
		valid := IsAtIndexValid(i, preamble, numbers)
		assert.Equal(t, expect, valid)
	}
}

func TestFindConsecutiveGroupByBruteForce(t *testing.T) {
	target := 127
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	group := FindConsecutiveGroupByBruteForce(target, numbers)
	min, max := helpers.FindMinAndMax(group)

	assert.Equal(t, []int{15, 25, 47, 40}, group)
	assert.Equal(t, 15, min)
	assert.Equal(t, 47, max)
	assert.Equal(t, 62, min+max)
}

func TestFindConsecutiveGroupByPrefixSum(t *testing.T) {
	target := 127
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	group := FindConsecutiveGroupByPrefixSum(target, numbers)
	min, max := helpers.FindMinAndMax(group)

	assert.Equal(t, []int{15, 25, 47, 40}, group)
	assert.Equal(t, 15, min)
	assert.Equal(t, 47, max)
	assert.Equal(t, 62, min+max)
}

func BenchmarkFindConsecutiveGroupByBruteForce(b *testing.B) {
	target := 127
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FindConsecutiveGroupByBruteForce(target, numbers)
	}
}

func BenchmarkFindConsecutiveGroupByPrefixSum(b *testing.B) {
	target := 127
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FindConsecutiveGroupByPrefixSum(target, numbers)
	}
}
