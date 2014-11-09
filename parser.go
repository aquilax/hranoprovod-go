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

func (p *Parser) parseFile(fileName string) (*NodeList, error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	return p.parseStream(bufio.NewReader(f))
}

func (p *Parser) parseStream(input *bufio.Reader) (*NodeList, error) {
	db := NewNodeList()
	line_number := 0

	node := NewNode()

	for {
		bytes, _, err := input.ReadLine()
		// handle errors
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		line := mytrim(string(bytes))

		line_number++

		//skip empty lines and lines starting with #
		if mytrim(line) == "" || line[0] == COMMENT_CHAR {
			continue
		}

		//new nodes start at the beginning of the line
		if bytes[0] != 32 && bytes[0] != 8 {
			if node.header != "" {
				if p.processor != nil {
					p.processor.process(node)
				} else {
					db.Push(node)
				}
			}
			node = NewNode()
			node.header = line
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
				(*node.elements)[ndx].val += float32(enum)
			} else {
				node.elements.Add(ename, float32(enum))
			}
		}
	}

	if node.header != "" {
		if p.processor != nil {
			p.processor.process(node)
		} else {
			db.Push(node)
		}
	}
	return db, nil
}
