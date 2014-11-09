package main

import (
	"flag"
	"fmt"
	"os"
)

type Hranoprovod struct {
	options *Options
	db      *NodeList
}

func NewHranoprovod() *Hranoprovod {
	return &Hranoprovod{}
}

func (hr *Hranoprovod) run() {
	var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
	hr.options = NewOptions(fs)

	if hr.options.version {
		fmt.Println("Hranoprovod version:", VERSION)
		os.Exit(EXIT_OK)
	}

	if hr.options.help {
		fmt.Println("Hranoprovod version:", VERSION)
		fmt.Println("Usage:")
		fs.PrintDefaults()
		os.Exit(EXIT_OK)
	}

	hr.db = NewParser(nil).parseFile(hr.options.database_file_name)
	NewResolver(hr.db).resolve()
	NewParser(NewProcessor(hr.options, hr.db)).parseFile(hr.options.log_file_name)
}
