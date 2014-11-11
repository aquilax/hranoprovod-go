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

	db, errp1 := NewParser(parserOptions, nil).parseFile(options.databaseFileName)
	if errp1 != nil {
		return errp1
	}
	NewResolver(db, resolverMaxDepth).resolve()

	_, errp2 := NewParser(parserOptions, NewProcessor(
		options,
		db,
		NewReporter(options, os.Stdout),
	)).parseFile(options.logFileName)
	return errp2
}
