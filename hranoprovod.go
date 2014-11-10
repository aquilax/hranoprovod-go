package main

import (
	"flag"
	"fmt"
	"os"
)

type Hranoprovod struct {
}

func NewHranoprovod() *Hranoprovod {
	return &Hranoprovod{}
}

func (hr *Hranoprovod) run() {
	var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
	options, optionsError := NewOptions(fs)
	if optionsError != nil {
		printError(optionsError)
		os.Exit(ERROR_PARSING_OPTIONS)
	}

	if options.version {
		fmt.Println("Hranoprovod version:", VERSION)
		os.Exit(EXIT_OK)
	}

	if options.help {
		fmt.Println("Hranoprovod version:", VERSION)
		fmt.Println("Usage:")
		fs.PrintDefaults()
		os.Exit(EXIT_OK)
	}

	db, error := NewParser(nil).parseFile(options.databaseFileName)
	if error != nil {
		os.Exit(ERROR_IO)
	}
	NewResolver(db).resolve()

	NewParser(NewProcessor(
		options, 
		db,
		NewReporter(options, os.Stdout),
	)).parseFile(options.logFileName)
}
