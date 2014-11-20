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
	p.parseStream(f)
}

func (p *Parser) parseStream(reader io.Reader) {
	var node *Node
	lineNumber := 0
	lineScanner := bufio.NewScanner(reader)
	for lineScanner.Scan() {
		lineNumber++
		line := lineScanner.Text()
		trimmedLine := mytrim(line)

		//skip empty lines and lines starting with #
		if trimmedLine == "" || line[0] == p.parserOptions.commentChar {
			continue
		}

		//new nodes start at the beginning of the line
		if line[0] != ' ' && line[0] != '\t' {
			if node != nil {
				p.nodes <- node
			}
			node = NewNode(trimmedLine)
			continue
		}

		if node != nil {
			separator := strings.LastIndexAny(trimmedLine, "\t ")

			if separator == -1 {
				p.errors <- NewBreakingError(
					fmt.Sprintf("Bad syntax on line %d, \"%s\".", lineNumber, line),
					exitErrorBadSyntax,
				)
				return
			}
			ename := mytrim(trimmedLine[0:separator])

			//get element value
			snum := mytrim(trimmedLine[separator:])
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
	// push last node
	if node != nil {
		p.nodes <- node
	}
	p.done <- true
}
