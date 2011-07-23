package main

import(
  "flag"
)

func (options *Options)InitFlags() *flag.FlagSet {
  var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
  fs.BoolVar(&(options.help), "help", false, "Shows this message")
  fs.StringVar(&(options.log_file_name), "f", "log.yaml", "Specifies log file name")
  fs.StringVar(&(options.database_file_name), "d", "food.yaml", "Specifies the database file name")
  return fs
}
