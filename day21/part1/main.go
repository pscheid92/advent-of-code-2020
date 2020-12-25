package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strings"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day21/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	foods := ExtractFoods(lines)
	mapping := BuildAllergenMapping(foods)
	count := CountOccurrencesOfHarmlessIngredients(foods, mapping)
	fmt.Printf("solution code: %d\n", count)
}

type Food struct {
	ingredients []string
	allergens   []string
}

func NewFood(ingredients []string, allergens []string) Food {
	return Food{
		ingredients: ingredients,
		allergens:   allergens,
	}
}

func ExtractFoods(lines []string) []Food {
	foods := make([]Food, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " (contains ")
		parts[1] = parts[1][:len(parts[1])-1]

		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(parts[1], ", ")
		foods[i] = NewFood(ingredients, allergens)
	}
	return foods
}

func BuildEliminationTable(foods []Food) map[string]StringSet {
	table := make(map[string]StringSet)
	for _, f := range foods {
		for _, allergen := range f.allergens {
			// collect ingredients in set
			newSet := NewStringSet()
			newSet.AddAll(f.ingredients...)

			// if first seen, we add this set as init value
			set, ok := table[allergen]
			if !ok {
				table[allergen] = newSet
				continue
			}

			// if already seen once, we intersect
			set.Intersect(newSet)
			table[allergen] = set
		}
	}
	return table
}

func BuildAllergenMapping(foods []Food) map[string]string {
	table := BuildEliminationTable(foods)
	mapping := make(map[string]string)
	for len(table) > 0 {
		for allergen, ingredients := range table {
			// more than one ingredient? skip
			if ingredients.Len() > 1 {
				continue
			}

			ingredient := ingredients.GetAll()[0]
			mapping[allergen] = ingredient

			// remove allergen and ingredient from elimination table
			delete(table, allergen)
			for key, set := range table {
				if set.Contains(ingredient) {
					set.Remove(ingredient)
					table[key] = set
				}
			}
			break
		}
	}
	return mapping
}

func CountOccurrencesOfHarmlessIngredients(foods []Food, harmfulMapping map[string]string) int {
	harmlessIngredients := NewStringSet()
	for _, f := range foods {
		harmlessIngredients.AddAll(f.ingredients...)
	}

	for _, harmfulIngredient := range harmfulMapping {
		harmlessIngredients.Remove(harmfulIngredient)
	}

	counter := 0
	for _, f := range foods {
		for _, i := range f.ingredients {
			if harmlessIngredients.Contains(i) {
				counter++
			}
		}
	}

	return counter
}
