package main

const (
	accNeg = 0
	accPos = 1
)

type AccValue [2]float32

type Accumulator map[string]*AccValue

func NewAccumulator() *Accumulator {
	return &Accumulator{}
}

func (acc *Accumulator) Add(name string, val float32) {
	sign := accPos
	if val < 0 {
		sign = accNeg
	}
	if _, exists := (*acc)[name]; exists {
		(*acc)[name][sign] += val
	} else {
		newVal := &AccValue{0, 0}
		newVal[sign] = val
		(*acc)[name] = newVal
	}
}
