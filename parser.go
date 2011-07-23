package main

import (
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

func Mytrim(s string) string{
  return strings.Trim(s, "\t \n:");
}

func (db *NodeList) Push(node Node){
  (*db)[node.name] = node;
}

func (db *NodeList) ParseFile(file_name string, callback func(node Node)){
  f, err := os.Open(file_name);
  if (err != nil) {
    log.Print(err)
  }

  input := bufio.NewReader(f)

  var node Node

  for {
    line, err := input.ReadString('\n')

    if err == os.EOF {
      break
    }

    if err != nil {
      log.Print(err)
      os.Exit(2)
    }

    //skip empty lines
    if (line[0] == 10){
      continue
    }
    //new nodes start at the beginning of the line
    if(line[0] != 32 && line[0] != 8){
      if node.name != "" {
        if (callback != nil) {
          callback(node);
        } else {
          db.Push(node)
        }
      }
      node.name = Mytrim(line)
      node.elements = make(Elements)
      continue
    }

    separator := strings.LastIndexAny(line, "\t ")

    ename := Mytrim(line[0:separator])
    enum, err := strconv.Atof32(Mytrim(line[separator:]))

    if err != nil{
      log.Println(err)
      continue
    }
    node.elements[ename] = enum;
  }
  if (node.name != ""){
    if callback != nil {
      callback(node)
    } else {
      db.Push(node);
    }
  }
  f.Close();
}
