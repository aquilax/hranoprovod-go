package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	COMMENT_CHAR = '#'
)

func Mytrim(s string) string {
	return strings.Trim(s, "\t \n:")
}

func (db *NodeList) Push(node *Node) {
	(*db)[(*node).name] = node
}

func CreateNode() *Node {
	var node Node
	return &node
}

func (db *NodeList) ParseFile(file_name string, callback func(node *Node)) {
	f, err := os.Open(file_name)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()

	input := bufio.NewReader(f)

	var line_number = 0

	var node = new(Node)

	for {
		arr, _, err := input.ReadLine()
		line := Mytrim(string(arr))

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Print(err)
			os.Exit(ERROR_IO)
		}

		line_number++

		//skip empty lines and lines starting with #
		if Mytrim(line) == "" || line[0] == COMMENT_CHAR {
			continue
		}

		//new nodes start at the beginning of the line
		if arr[0] != 32 && arr[0] != 8 {
			if node.name != "" {
				if callback != nil {
					callback(node)
				} else {
					db.Push(node)
				}
			}
			node = CreateNode()
			node.name = line
			continue
		}

		if node != nil {
			line = Mytrim(line)
			separator := strings.LastIndexAny(line, "\t ")

			if separator == -1 {
				log.Printf("Bad syntax on line %d, \"%s\".", line_number, line)
				os.Exit(ERROR_BAD_SYNTAX)
			}

			ename := Mytrim(line[0:separator])
			snum := Mytrim(line[separator:])
			enum, err := strconv.ParseFloat(snum, 32)

			if err != nil {
				log.Printf("Error converting \"%s\" to float on line %d \"%s\".", snum, line_number, line)
				os.Exit(ERROR_CONVERSION)
			}
			ndx, exists := node.elements.Index(ename)
			if exists {
				node.elements[ndx].val += float32(enum)
			} else {
				node.elements.Add(ename, float32(enum))
			}
		}
	}

	if node.name != "" {
		if callback != nil {
			callback(node)
		} else {
			db.Push(node)
		}
	}
	f.Close()
}
