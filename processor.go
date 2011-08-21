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
  for element, _ := range node.elements {
    _, found := db[element]
    if !found {
      fmt.Println(element)
    }
  }
}

func SingleProcessor(node Node) {
  acc := make(Accumulator);
  ts, err := time.Parse("2006/01/02", Mytrim(node.name));
  if (err != nil){
    log.Print(err)
  }
  for element, coef := range node.elements {
    repl, found := db[element]
    if found {
      for repl, val := range repl.elements{
        if repl == options.single_element {
          acc.Add(repl, val* coef)
        }
      }
    } else {
      if element == options.single_element {
        acc.Add(element, coef)
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
  for element, coef := range node.elements {
    matched, err := regexp.MatchString(options.single_food, element)
    if err != nil {
      log.Print(err)
      os.Exit(3)
    }
    if matched {
      fmt.Printf("%s\t%s\t%0.2f\n", ts.Format("2006/01/02"), element, coef)
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
  for element, coef := range node.elements {
    fmt.Printf("\t%s : %01.2f\n", element, coef);
    repl, found := db[element]
    if found {
      for repl, val := range repl.elements{
        res := val*coef;
        fmt.Printf("\t\t%20s %10.2f\n", repl, res)
        acc.Add(repl, res)
      }
    } else  {
      fmt.Printf("\t\t%20s %10.2f\n", element, coef)
      acc.Add(element, coef)
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



