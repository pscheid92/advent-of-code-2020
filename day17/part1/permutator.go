package main

type Permutation struct {
	N          int
	Dimensions []int
	state      []int
	init       bool
	done       bool
}

func InitPermutation(dimensions ...int) Permutation {
	return Permutation{
		N:          len(dimensions),
		Dimensions: dimensions,
		state:      make([]int, len(dimensions)),
		init:       true,
		done:       false,
	}
}

func (p Permutation) Value() []int {
	result := make([]int, p.N)
	copy(result, p.state)
	return result
}

func (p *Permutation) Next() bool {
	// skip if already at end
	if p.done {
		return false
	}

	// make sure the initial state is returned on first call
	if p.init {
		p.init = false
		return true
	}

	// increment dimension and iteratively overflow next dimension
	// if no overflow occurs, we are done
	// if an overflow occurs, we reset this dimension and increment the next one
	// if all dimensions overflow, we run trough the whole loop and leave it (without return)
	for i, d := range p.Dimensions {
		p.state[i]++
		if p.state[i] < d {
			return true
		}
		p.state[i] = 0
	}

	// all dimensions overflowed, we are done and there is no next element
	p.done = true
	return false
}

func (p *Permutation) Reset() {
	p.init = true
	p.done = false
	p.state = make([]int, p.N)
}
