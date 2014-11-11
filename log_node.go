package main

import (
	"time"
)

// LogNode contains log node data
type LogNode struct {
	time     time.Time
	elements *Elements
}

// NewLogNode creates new log node
func NewLogNode(time time.Time, elements *Elements) *LogNode {
	return &LogNode{time, elements}
}
