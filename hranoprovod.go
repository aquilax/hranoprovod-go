package main

import (
//  "fmt"
)

func main(){
  var db = make(NodeList)
  db.ParseFile("food.yaml")
  db.Resolve();
//  fmt.Print(db)

  var log = make(NodeList)
  log.ParseFile("log.yaml")
//  fmt.Print(log)
}
