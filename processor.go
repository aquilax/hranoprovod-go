package main

import(
  "log"
  "time"
)

func processor(node *Node){
  ts, err := time.Parse("2006/01/02", Mytrim(node.name));
  if (err != nil){
    log.Print(err)
  }
  log.Print(options)
  log.Print(node)
  log.Print(ts)
}

