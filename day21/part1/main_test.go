package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var input = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`

func TestExtractFoods(t *testing.T) {
	expectedFoods := []Food{
		{[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, []string{"dairy", "fish"}},
		{[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, []string{"dairy"}},
		{[]string{"sqjhc", "fvjkl"}, []string{"soy"}},
		{[]string{"sqjhc", "mxmxvkd", "sbzzf"}, []string{"fish"}},
	}

	lines := helpers.ReadLineByLine(strings.NewReader(input))
	foods := ExtractFoods(lines)
	assert.ElementsMatch(t, expectedFoods, foods)
}

func TestBuildEliminationTable(t *testing.T) {
	expectedTable := map[string][]string{
		"dairy": {"mxmxvkd"},
		"fish":  {"mxmxvkd", "sqjhc"},
		"soy":   {"fvjkl", "sqjhc"},
	}

	foods := []Food{
		{[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, []string{"dairy", "fish"}},
		{[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, []string{"dairy"}},
		{[]string{"sqjhc", "fvjkl"}, []string{"soy"}},
		{[]string{"sqjhc", "mxmxvkd", "sbzzf"}, []string{"fish"}},
	}

	table := BuildEliminationTable(foods)
	assert.Len(t, table, 3)

	for k, v := range expectedTable {
		assert.Contains(t, table, k)
		assert.ElementsMatch(t, v, table[k].GetAll())
	}
}

func TestBuildAllergenMapping(t *testing.T) {
	expectedMapping := map[string]string{
		"dairy": "mxmxvkd",
		"fish":  "sqjhc",
		"soy":   "fvjkl",
	}

	foods := []Food{
		{[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, []string{"dairy", "fish"}},
		{[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, []string{"dairy"}},
		{[]string{"sqjhc", "fvjkl"}, []string{"soy"}},
		{[]string{"sqjhc", "mxmxvkd", "sbzzf"}, []string{"fish"}},
	}

	mapping := BuildAllergenMapping(foods)
	assert.Equal(t, expectedMapping, mapping)
}

func TestCountOccurrencesOfHarmlessIngredients(t *testing.T) {
	lines := helpers.ReadLineByLine(strings.NewReader(input))
	foods := ExtractFoods(lines)
	mapping := BuildAllergenMapping(foods)

	count := CountOccurrencesOfHarmlessIngredients(foods, mapping)
	assert.Equal(t, 5, count)
}
