package main

func LoopPower(base, exponent, modulo uint) uint {
	res := base
	for i := uint(1); i < exponent; i++ {
		res *= base
		res %= modulo
	}
	return res
}

func Power(base, exponent uint, modulo uint) uint {
	// square and multiply algorithm
	result := uint(1)
	for exponent > 0 {
		if exponent&1 == 1 {
			result *= base
			result %= modulo
		}
		base *= base
		base %= modulo
		exponent >>= 1
	}
	return result
}
