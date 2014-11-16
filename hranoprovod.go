package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	resolverMaxDepth = 9
)

// Hranoprovod is the main program structure
type Hranoprovod struct {
}

// NewHranoprovod creates new program structure
func NewHranoprovod() *Hranoprovod {
	return &Hranoprovod{}
}

func readDatabase(parser *Parser) (*NodeList, error) {
	nodeList := NewNodeList()
	for {
		select {
		case node := <-parser.nodes:
			nodeList.push(node)
		case breakingError := <-parser.errors:
			return nil, breakingError
		case <-parser.done:
			return nodeList, nil
		}
	}
}

// Run runs the program
func (hr *Hranoprovod) Run(version string) error {
	var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
	options, optionsError := NewOptions(fs)
	if optionsError != nil {
		return optionsError
	}

	if options.version {
		fmt.Println("Hranoprovod version:", version)
		return nil
	}

	if options.help {
		fmt.Println("Hranoprovod version:", version)
		fmt.Println("Usage:")
		fs.PrintDefaults()
		return nil
	}

	parserOptions := NewDefaultParserOptions()

	parser := NewParser(parserOptions)
	go parser.parseFile(options.databaseFileName)
	nodeList, err := readDatabase(parser)
	if err != nil {
		return err
	}
	NewResolver(nodeList, resolverMaxDepth).resolve()

	processor := NewProcessor(
		options,
		nodeList,
		NewReporter(options, os.Stdout),
	)

	go parser.parseFile(options.logFileName)
	for {
		select {
		case node := <-parser.nodes:
			processor.process(node)
		case breakingError := <-parser.errors:
			return breakingError
		case <-parser.done:
			return nil
		}
	}
	return nil
}
