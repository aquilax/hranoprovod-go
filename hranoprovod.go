package main

import (
  "os"
  "log"
)

func processor(node *Node){
  log.Print(node)
}

func main(){
  var options Options;
  var fs = options.InitFlags();
  fs.Parse(os.Args[1:])

  if (options.help){
    fs.PrintDefaults()
    os.Exit(1)
  }

  var db = make(NodeList)
  db.ParseFile(options.database_file_name, nil)
  db.Resolve();

  var mylog = make(NodeList)
  mylog.ParseFile(options.log_file_name, processor)
  //fmt.Print(mylog)
}
