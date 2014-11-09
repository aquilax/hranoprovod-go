package main

import (
	"flag"
	"os"
	"time"
)

const (
	DEFAULT_DB_FILENAME  = "food.yaml"
	DEFAULT_LOG_FILENAME = "log.yaml"
)

type Options struct {
	unresolved       bool
	version          bool
	help             bool
	totals           bool
	csv              bool
	color            bool
	singleElement    string
	singleFood       string
	hasBeginning     bool
	beginning        string
	beginningTime    time.Time
	hasEnd           bool
	end              string
	endTime          time.Time
	logFileName      string
	databaseFileName string
}

func NewOptions(fs *flag.FlagSet) (*Options, error) {
	var options Options
	fs.BoolVar(&(options.help), "help", false, "Shows this message")
	fs.BoolVar(&(options.totals), "total", true, "Shows totals for each day")
	fs.BoolVar(&(options.version), "version", false, "Shows version")
	fs.BoolVar(&(options.unresolved), "unresolved", false, "Shows unresolved elements")
	fs.BoolVar(&(options.csv), "csv", false, "Export in CSV format")
	fs.BoolVar(&(options.color), "color", true, "Color output")

	fs.StringVar(&(options.singleFood), "food", "", "Shows single food")
	fs.StringVar(&(options.singleElement), "single", "", "Show only single element")
	fs.StringVar(&(options.beginning), "b", "", "Beginning of date interval (YYYY/MM/DD)")
	fs.StringVar(&(options.end), "e", "", "Ending of date interval (YYYY/MM/DD)")

	fs.StringVar(&(options.databaseFileName), "d", DEFAULT_DB_FILENAME, "Specifies the database file name")
	fs.StringVar(&(options.logFileName), "f", DEFAULT_LOG_FILENAME, "Specifies log file name")
	fs.Parse(os.Args[1:])
	return &options, options.processOptions()
}

func (o *Options) processOptions() error {
	var err error
	o.hasBeginning = false
	o.hasEnd = false
	if len(o.beginning) > 0 {
		o.beginningTime, err = parseTime(o.beginning)
		if err != nil {
			return err
		}
		o.hasBeginning = true
	}
	if len(o.end) > 0 {
		o.endTime, err = parseTime(o.end)
		if err != nil {
			return err
		}
		o.hasEnd = true
	}
	return nil
}
