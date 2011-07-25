package main

type Options struct{
  help bool
  totals bool
  single_element string

  log_file_name string
  database_file_name string
}

type Elements map[string] float32

type Node struct {
  name string
  elements Elements
}

type NodeList map[string] Node

type Accumulator map[string] [2]float32
