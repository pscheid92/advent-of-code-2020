package main

type Mask struct {
	setMask   uint64
	clearMask uint64
}

func NewMask(maskSpecification string) Mask {
	mask := Mask{}

	for _, m := range maskSpecification {
		mask.setMask <<= 1
		mask.clearMask <<= 1

		if m == '1' {
			mask.setMask |= 0b1
		}

		if m == '0' {
			mask.clearMask |= 0b1
		}
	}

	mask.clearMask = ^mask.clearMask
	return mask
}

func (m Mask) Apply(x uint64) uint64 {
	x |= m.setMask
	x &= m.clearMask
	return x
}
