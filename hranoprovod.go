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

	db, errp1 := NewParser(nil).parseFile(options.databaseFileName)
	if errp1 != nil {
		return errp1
	}
	NewResolver(db).resolve()

	_, errp2 := NewParser(NewProcessor(
		options,
		db,
		NewReporter(options, os.Stdout),
	)).parseFile(options.logFileName)
	return errp2
}
