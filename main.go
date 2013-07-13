package main

import (
	"fmt"
	"os"
)

const (
	VERSION = "0.1.2"

	EXIT_OK = iota
	ERROR_IO
	ERROR_BAD_SYNTAX
	ERROR_CONVERSION
	ERROR_SINGLE_FOOD_NOT_FOUND
)

var options Options
var db *NodeList

func main() {
	var fs = options.InitFlags()
	fs.Parse(os.Args[1:])

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

	db = NewNodeList()
	db.ParseFile(options.database_file_name, nil)
	db.Resolve()

	var mylog = make(NodeList)
	mylog.ParseFile(options.log_file_name, processor)
}
