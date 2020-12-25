package main

type Memory struct {
	Mask Mask
	data map[uint64]uint64
}

func NewMemory() Memory {
	return Memory{
		Mask: Mask{},
		data: make(map[uint64]uint64),
	}
}

func (m Memory) Set(address uint64, value uint64) {
	val := m.Mask.Apply(value)
	m.data[address] = val
}

func (m Memory) SumData() uint64 {
	sum := uint64(0)
	for _, v := range m.data {
		sum += v
	}
	return sum
}
