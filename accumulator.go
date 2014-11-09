package main

const (
	accNeg = 0
	accPos = 1
)

type Accumulator map[string][2]float32

func NewAccumulator() *Accumulator {
	accumulator := make(Accumulator)
	return &accumulator
}

func (acc *Accumulator) Add(name string, val float32) {
	value_sign := accPos
	if val < 0 {
		value_sign = accNeg
	}
	oldval, exists := (*acc)[name]
	if exists {
		oldval[value_sign] = val + oldval[value_sign]
		(*acc)[name] = oldval
	} else {
		newval := [2]float32{0, 0}
		newval[value_sign] = val
		(*acc)[name] = newval
	}
}
