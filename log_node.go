package main

import (
	"time"
)

type LogNode struct {
	time     time.Time
	elements *Elements
}

func NewLogNode(time time.Time, elements *Elements) *LogNode {
	return &LogNode{time, elements}
}
