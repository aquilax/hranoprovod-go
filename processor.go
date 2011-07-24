package main

import(
  "log"
  "fmt"
  "time"
)

func processor(node Node){
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

