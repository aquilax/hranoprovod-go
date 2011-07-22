package main

import (
)

type Elements map[string] float32

type Node struct{
  name string
  elements Elements
}

type NodeList map[string] Node
