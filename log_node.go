package main

import (
	"github.com/Hranoprovod/parser"
	"time"
)

// LogNode contains log node data
type LogNode struct {
	time     time.Time
	elements *parser.Elements
}

// NewLogNode creates new log node
func NewLogNode(time time.Time, elements *parser.Elements) *LogNode {
	return &LogNode{time, elements}
}
