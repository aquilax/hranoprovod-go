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

func (db *NodeList) Push(node *Node){
  (*db)[(*node).name] = node;
}

func CreateNode(name string, elements Elements) *Node{ 
  var node Node;
  node.name = name;
  node.elements = elements;
  return &node;
}

func (db *NodeList) ParseFile(file_name string, callback func(node *Node)){

  f, err := os.Open(file_name);
  if (err != nil) {
    log.Print(err)
  }

  input := bufio.NewReader(f)

  var name string = ""
  var elements Elements

  for {
    arr, _, err := input.ReadLine()
    line := Mytrim(string(arr));

    if err == os.EOF {
      break
    }

    if err != nil {
      log.Print(err)
      os.Exit(2)
    }

    //skip empty lines
    if (Mytrim(line) == ""){
      continue
    }

    //new nodes start at the beginning of the line
    if (arr[0] != 32 && arr[0] != 8) {
      if name != "" {
        node := CreateNode(name, elements)
        if (callback != nil) {
          callback(node);
        } else {
          db.Push(node)
        }
      }
      name = line;
      elements = make(Elements)
      continue
    }

    if (name != ""){
      line = Mytrim(line)
      separator := strings.LastIndexAny(line, "\t ")

      ename := Mytrim(line[0:separator])
      snum := Mytrim(line[separator:])
      enum, err := strconv.Atof32(snum)

      if err != nil{
        log.Printf("Error converting %s to float from line \"%s\". %s", snum, line, err)
        continue
      }
      val, exists := elements[ename]
      if exists {
        enum += val
      }
      elements[ename] = enum;
    }
  }

  if (name != "") {
    node := CreateNode(name, elements)
    if callback != nil {
      callback(node)
    } else {
      db.Push(node);
    }
  }
  f.Close();
}
