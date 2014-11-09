package main

import (
	"flag"
	"os"
)

const (
	DEFAULT_DB_FILENAME  = "food.yaml"
	DEFAULT_LOG_FILENAME = "log.yaml"
)

type Options struct {
	unresolved     bool
	version        bool
	help           bool
	totals         bool
	csv            bool
	color          bool
	single_element string
	single_food    string
	beginning      string
	end            string

	log_file_name      string
	database_file_name string
}

func NewOptions(fs *flag.FlagSet) *Options {
	var options Options
	fs.BoolVar(&(options.help), "help", false, "Shows this message")
	fs.BoolVar(&(options.totals), "total", true, "Shows totals for each day")
	fs.BoolVar(&(options.version), "version", false, "Shows version")
	fs.BoolVar(&(options.unresolved), "unresolved", false, "Shows unresolved elements")
	fs.BoolVar(&(options.csv), "csv", false, "Export in CSV format")
	fs.BoolVar(&(options.color), "color", true, "Color output")

	fs.StringVar(&(options.single_food), "food", "", "Shows single food")
	fs.StringVar(&(options.single_element), "single", "", "Show only single element")
	fs.StringVar(&(options.beginning), "b", "", "Beginning of date interval (YYYY/MM/DD)")
	fs.StringVar(&(options.end), "e", "", "Ending of date interval (YYYY/MM/DD)")

	fs.StringVar(&(options.database_file_name), "d", DEFAULT_DB_FILENAME, "Specifies the database file name")
	fs.StringVar(&(options.log_file_name), "f", DEFAULT_LOG_FILENAME, "Specifies log file name")
	fs.Parse(os.Args[1:])
	return &options
}
