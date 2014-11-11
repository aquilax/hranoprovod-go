package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ParserOptions struct {
	commentChar uint8
}

func NewDefaultParserOptions() *ParserOptions {
	return &ParserOptions{'#'}
}

type Parser struct {
	parserOptions *ParserOptions
	processor     *Processor
}

func NewParser(parserOptions *ParserOptions, processor *Processor) *Parser {
	return &Parser{
		parserOptions,
		processor,
	}
}

func (p *Parser) parseFile(fileName string) (*NodeList, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, NewBreakingError(err.Error(), ERROR_OPENING_FILE)
	}
	defer f.Close()
	return p.parseStream(bufio.NewReader(f))
}

func (p *Parser) parseStream(input *bufio.Reader) (*NodeList, error) {
	var node *Node
	db := NewNodeList()
	line_number := 0

	for {
		bytes, _, err := input.ReadLine()
		// handle errors
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, NewBreakingError(err.Error(), ERROR_IO)
		}

		line := mytrim(string(bytes))

		line_number++

		//skip empty lines and lines starting with #
		if mytrim(line) == "" || line[0] == p.parserOptions.commentChar {
			continue
		}

		//new nodes start at the beginning of the line
		if bytes[0] != 32 && bytes[0] != 8 {
			if node != nil {
				if p.processor != nil {
					p.processor.process(node)
				} else {
					db.Push(node)
				}
			}
			node = NewNode(line)
			continue
		}

		if node != nil {
			line = mytrim(line)
			separator := strings.LastIndexAny(line, "\t ")

			if separator == -1 {
				return nil, NewBreakingError(
					fmt.Sprintf("Bad syntax on line %d, \"%s\".", line_number, line),
					ERROR_BAD_SYNTAX,
				)
			}

			ename := mytrim(line[0:separator])
			snum := mytrim(line[separator:])
			enum, err := strconv.ParseFloat(snum, 32)

			if err != nil {
				return nil, NewBreakingError(
					fmt.Sprintf("Error converting \"%s\" to float on line %d \"%s\".", snum, line_number, line),
					ERROR_CONVERSION,
				)
			}
			if ndx, exists := node.elements.Index(ename); exists {
				(*node.elements)[ndx].val += float32(enum)
			} else {
				node.elements.Add(ename, float32(enum))
			}
		}
	}

	if node != nil {
		if p.processor != nil {
			p.processor.process(node)
		} else {
			db.Push(node)
		}
	}
	return db, nil
}
