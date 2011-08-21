package main

import(
  "sort"
  "os"
  "log"
  "fmt"
  "time"
  "regexp"
)

func processor(node *Node){
  if len(options.beginning) > 0 {
    if !GoodDate(node.name, options.beginning, 1) {
      return
    }
  }
  if len(options.end) > 0 {
    if !GoodDate(node.name, options.end, 2) {
      return
    }
  }
  if options.unresolved {
    UnresolvedProcessor(*node)
    return;
  }
  if len(options.single_element) > 0 {
    SingleProcessor(*node)
    return
  }
  if len(options.single_food) > 0 {
    SingleFoodProcessor(*node)
    return
  }
  DefaultProcessor(*node)
}

func ParseTime(date string) (*time.Time){
  ts, err := time.Parse("2006/01/02", Mytrim(date));
  if (err != nil){
    log.Print(err)
  }
  return ts
}

func GoodDate(name, compare string, ctype int) bool {
  ts := ParseTime(name)
  tsb := ParseTime(compare)
  if ctype ==1 {
    return ts.Seconds() >= tsb.Seconds()
  }
  return ts.Seconds() <= tsb.Seconds()
}

func UnresolvedProcessor(node Node) {
  for _, e := range node.elements {
    _, found := db[e.name]
    if !found {
      fmt.Println(e.name)
    }
  }
}

func SingleProcessor(node Node) {
  acc := make(Accumulator);
  ts, err := time.Parse("2006/01/02", Mytrim(node.name));
  if (err != nil){
    log.Print(err)
  }
  for _, e := range node.elements {
    repl, found := db[e.name]
    if found {
      for _, repl := range repl.elements{
        if repl.name == options.single_element {
          acc.Add(repl.name, repl.val * e.val)
        }
      }
    } else {
      if e.name == options.single_element {
        acc.Add(e.name, e.val)
      }
    }
  }
  if len(acc) > 0 {
    arr := acc[options.single_element]
    if options.csv {
      fmt.Printf("%s;%s;%0.2f;%0.2f;%0.2f\n", ts.Format("2006/01/02"), options.single_element, arr[1], -1*arr[0], arr[0]+arr[1]);
    } else {
      fmt.Printf("%s %20s %10.2f %10.2f =%10.2f\n", ts.Format("2006/01/02"), options.single_element, arr[1],arr[0], arr[0]+arr[1]); 
    }
  }
}

func SingleFoodProcessor(node Node) {
  ts, err := time.Parse("2006/01/02", Mytrim(node.name));
  if (err != nil){
    log.Print(err)
  }
  for _, e := range node.elements {
    matched, err := regexp.MatchString(options.single_food, e.name)
    if err != nil {
      log.Print(err)
      os.Exit(3)
    }
    if matched {
      fmt.Printf("%s\t%s\t%0.2f\n", ts.Format("2006/01/02"), e.name, e.val)
    }
  }
}

func DefaultProcessor(node Node) {
  acc := make(Accumulator);
  ts, err := time.Parse("2006/01/02", Mytrim(node.name));
  if (err != nil){
    log.Print(err)
  }
  fmt.Printf("%s\n", ts.Format("2006/01/02"))
  for _, e := range node.elements {
    fmt.Printf("\t%s : %01.2f\n", e.name, e.val);
    repl, found := db[e.name]
    if found {
      for _, repl := range repl.elements{
        res := repl.val * e.val;
        fmt.Printf("\t\t%20s %10.2f\n", repl.name, res)
        acc.Add(repl.name, res)
      }
    } else  {
      fmt.Printf("\t\t%20s %10.2f\n", e.name, e.val)
      acc.Add(e.name, e.val)
    }
  }
  if options.totals {
    var ss sort.StringSlice
    if len(acc) > 0 {
    fmt.Printf("\t%s\n", "-- TOTAL ----");
      for name, _ := range acc{
        ss = append(ss, name)
      }
      sort.Sort(ss)
      for _, name := range ss{
        arr := acc[name]
        fmt.Printf("\t\t%20s %10.2f %10.2f =%10.2f\n", name, arr[1],arr[0], arr[0]+arr[1]);
      }
    }
  }
}



