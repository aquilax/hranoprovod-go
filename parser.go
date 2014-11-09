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

type Parser struct {
	processor *Processor
}

func NewParser(processor *Processor) *Parser {
	return &Parser{processor}
}

func (p *Parser) parseFile(fileName string) *NodeList {
	f, err := os.Open(fileName)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	return p.parseStream(bufio.NewReader(f))
}

func (p *Parser) parseStream(input *bufio.Reader) *NodeList {
	db := NewNodeList()
	line_number := 0

	var node = new(Node)

	for {
		arr, _, err := input.ReadLine()
		line := mytrim(string(arr))

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Print(err)
			os.Exit(ERROR_IO)
		}

		line_number++

		//skip empty lines and lines starting with #
		if mytrim(line) == "" || line[0] == COMMENT_CHAR {
			continue
		}

		//new nodes start at the beginning of the line
		if arr[0] != 32 && arr[0] != 8 {
			if node.name != "" {
				if p.processor != nil {
					p.processor.process(node)
				} else {
					db.Push(node)
				}
			}
			node = NewNode()
			node.name = line
			continue
		}

		if node != nil {
			line = mytrim(line)
			separator := strings.LastIndexAny(line, "\t ")

			if separator == -1 {
				log.Printf("Bad syntax on line %d, \"%s\".", line_number, line)
				os.Exit(ERROR_BAD_SYNTAX)
			}

			ename := mytrim(line[0:separator])
			snum := mytrim(line[separator:])
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
		if p.processor != nil {
			p.processor.process(node)
		} else {
			db.Push(node)
		}
	}
	return db
}
