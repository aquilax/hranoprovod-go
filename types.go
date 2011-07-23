package main

type Options struct{
  help bool

  log_file_name string
  database_file_name string
}


type Elements map[string] float32

type Node struct {
  name string
  elements Elements
}

type NodeList map[string] *Node
