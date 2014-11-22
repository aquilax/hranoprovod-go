package main

import (
	"flag"
	"fmt"
	"github.com/Hranoprovod/parser"
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

	nodeList := NewNodeList()
	prsr := parser.NewParser(parser.NewDefaultOptions())
	go prsr.ParseFile(options.databaseFileName)
	err := func() error {
		for {
			select {
			case node := <-prsr.Nodes:
				nodeList.push(node)
			case breakingError := <-prsr.Errors:
				return breakingError
			case <-prsr.Done:
				return nil
			}
		}
	}()

	if err != nil {
		return err
	}
	NewResolver(nodeList, resolverMaxDepth).resolve()

	processor := NewProcessor(
		options,
		nodeList,
		NewReporter(options, os.Stdout),
	)

	go prsr.ParseFile(options.logFileName)
	for {
		select {
		case node := <-prsr.Nodes:
			processor.process(node)
		case breakingError := <-prsr.Errors:
			return breakingError
		case <-prsr.Done:
			return nil
		}
	}
	return nil
}
