package main

const (
	accNeg = 0
	accPos = 1
)

// AccValues contains accumulator values
type AccValues [2]float32

// Accumulator accumulates element values by name
type Accumulator map[string]*AccValues

// NewAccumulator returns new accumulator
func NewAccumulator() *Accumulator {
	return &Accumulator{}
}

func (acc *Accumulator) add(name string, val float32) {
	sign := accPos
	if val < 0 {
		sign = accNeg
	}
	if _, exists := (*acc)[name]; exists {
		(*acc)[name][sign] += val
	} else {
		newVal := &AccValues{0, 0}
		newVal[sign] = val
		(*acc)[name] = newVal
	}
}
