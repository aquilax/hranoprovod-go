package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ParserOptions contains the parser related options
type ParserOptions struct {
	commentChar uint8
}

// NewDefaultParserOptions returns the default set of parser options
func NewDefaultParserOptions() *ParserOptions {
	return &ParserOptions{'#'}
}

// Parser is the parser data structure
type Parser struct {
	parserOptions *ParserOptions
	nodes         chan *Node
	errors        chan *BreakingError
	done          chan bool
}

// NewParser returns new parser
func NewParser(parserOptions *ParserOptions) *Parser {
	return &Parser{
		parserOptions,
		make(chan *Node),
		make(chan *BreakingError),
		make(chan bool),
	}
}

func (p *Parser) parseFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		p.errors <- NewBreakingError(err.Error(), exitErrorOpeningFile)
		return
	}
	defer f.Close()
	p.parseStream(bufio.NewReader(f))
}

func (p *Parser) parseStream(input *bufio.Reader) {
	var node *Node
	lineNumber := 0

	for {
		bytes, _, err := input.ReadLine()
		// handle errors
		if err == io.EOF {
			// push last node
			if node != nil {
				p.nodes <- node
			}
			break
		}
		if err != nil {
			p.errors <- NewBreakingError(err.Error(), exitErrorIO)
			return
		}

		line := mytrim(string(bytes))

		lineNumber++

		//skip empty lines and lines starting with #
		if mytrim(line) == "" || line[0] == p.parserOptions.commentChar {
			continue
		}

		//new nodes start at the beginning of the line
		if bytes[0] != 32 && bytes[0] != 8 {
			if node != nil {
				p.nodes <- node
			}
			node = NewNode(line)
			continue
		}

		if node != nil {
			line = mytrim(line)
			separator := strings.LastIndexAny(line, "\t ")

			if separator == -1 {
				p.errors <- NewBreakingError(
					fmt.Sprintf("Bad syntax on line %d, \"%s\".", lineNumber, line),
					exitErrorBadSyntax,
				)
				return
			}

			ename := mytrim(line[0:separator])
			snum := mytrim(line[separator:])
			enum, err := strconv.ParseFloat(snum, 32)

			if err != nil {
				p.errors <- NewBreakingError(
					fmt.Sprintf("Error converting \"%s\" to float on line %d \"%s\".", snum, lineNumber, line),
					exitErrorConversion,
				)
				return
			}
			if ndx, exists := node.elements.index(ename); exists {
				(*node.elements)[ndx].val += float32(enum)
			} else {
				node.elements.add(ename, float32(enum))
			}
		}
	}
	p.done <- true
}
