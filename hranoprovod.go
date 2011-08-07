package main

import (
  "os"
  "fmt"
)

const version string = "0.1"

var options Options
var db NodeList

func main(){
  var fs = options.InitFlags()
  fs.Parse(os.Args[1:])

  if (options.version) {
    fmt.Println("Hranoprovod version:", version)
    os.Exit(1);
  }

  if (options.help) {
    fmt.Println("Hranoprovod version:", version)
    fmt.Println("Usage:");
    fs.PrintDefaults()
    os.Exit(1)
  }

  db = make(NodeList)
  db.ParseFile(options.database_file_name, nil)
  db.Resolve()

  var mylog = make(NodeList)
  mylog.ParseFile(options.log_file_name, processor)
}
