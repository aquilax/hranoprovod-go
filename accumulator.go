package main

type Accumulator map[string][2]float32

func NewAccumulator() *Accumulator {
	accumulator := make(Accumulator)
	return &accumulator
}

func (acc *Accumulator) Add(name string, val float32) {
	ndx := 1
	if val < 0 {
		ndx = 0
	}
	oldval, exists := (*acc)[name]
	if exists {
		oldval[ndx] = val + oldval[ndx]
		(*acc)[name] = oldval
	} else {
		newval := [2]float32{0, 0}
		newval[ndx] = val
		(*acc)[name] = newval
	}
}
