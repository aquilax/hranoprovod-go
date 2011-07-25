package main

import(
  "log"
  "fmt"
  "time"
)

func processor(node *Node){
  if len(options.single_element) > 0 {
    SingleProcessor(*node)
    return
  }
  DefaultProcessor(*node)
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
    fmt.Printf("%s %20s %10.2f %10.2f =%10.2f\n", ts.Format("2006/01/02"), options.single_element, arr[1],arr[0], arr[0]+arr[1]);
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
    if len(acc) > 0 {
    fmt.Printf("\t%s\n", "-- TOTAL ----");
      for name, arr := range acc{
        fmt.Printf("\t\t%20s %10.2f %10.2f =%10.2f\n", name, arr[1],arr[0], arr[0]+arr[1]);
      }
    }
  }
}



