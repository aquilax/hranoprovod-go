package main

import(
  "flag"
)

func (options *Options)InitFlags() *flag.FlagSet {
  var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
  fs.BoolVar(&(options.help), "help", false, "Shows this message")
  fs.BoolVar(&(options.totals), "total", true, "Shows totals for each day")
  fs.BoolVar(&(options.version), "version", false, "Shows version")
  fs.StringVar(&(options.single_element), "single", "", "Show only single element")
  fs.StringVar(&(options.beginning), "b", "", "Beginning of date interval (YYYY/MM/DD)")
  fs.StringVar(&(options.end), "e", "", "Ending of date interval (YYYY/MM/DD)")

  fs.StringVar(&(options.log_file_name), "f", "log.yaml", "Specifies log file name")
  fs.StringVar(&(options.database_file_name), "d", "food.yaml", "Specifies the database file name")
  return fs
}
