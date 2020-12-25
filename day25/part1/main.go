package main

import (
	"fmt"
	"math"
)

const subjectNumber = 7
const primeNumber = 20201227

func main() {
	cardPublicKey := uint(9789649)
	doorPublicKey := uint(3647239)

	sharedKey := StealSharedKey(cardPublicKey, doorPublicKey)
	fmt.Printf("solution code: %d\n", sharedKey)
}

func StealSharedKey(cardPublicKey uint, doorPublicKey uint) uint {
	cardSecretKey := SolveByBabyStepGiantStep(subjectNumber, cardPublicKey, primeNumber)
	return Power(doorPublicKey, cardSecretKey, primeNumber)
}

func SolveByBabyStepGiantStep(base, result, modulo uint) uint {
	lookup := make(map[uint]uint)
	m := uint(math.Ceil(math.Sqrt(float64(modulo))))

	tmp := uint(1)
	for j := uint(0); j < m; j++ {
		lookup[tmp] = j
		tmp = (tmp * base) % modulo
	}

	factor := Power(base, modulo-m-1, modulo)
	tmp = result

	for i := uint(0); i < m; i++ {
		if j, ok := lookup[tmp]; ok {
			return i*m + j
		} else {
			tmp = (tmp * factor) % modulo
		}
	}

	return 0
}
