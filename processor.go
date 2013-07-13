package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	IN_DATE_FMT  = "2006/01/02"
	OUT_DATE_FMT = "2006/01/02"
)

func processor(node *Node) {
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
		return
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

func ParseTime(date string) (time.Time, bool) {
	ts, err := time.Parse(IN_DATE_FMT, Mytrim(date))
	ok := true
	if err != nil {
		log.Print(err)
		ok = false
	}
	return ts, ok
}

func GoodDate(name, compare string, ctype int) bool {
	ts, _ := ParseTime(name)
	tsb, _ := ParseTime(compare)
	if ctype == 1 {
		return ts.Unix() >= tsb.Unix()
	}
	return ts.Unix() <= tsb.Unix()
}

func UnresolvedProcessor(node Node) {
	for _, e := range node.elements {
		_, found := (*db)[e.name]
		if !found {
			fmt.Println(e.name)
		}
	}
}

func SingleProcessor(node Node) {
	acc := NewAccumulator()
	ts, _ := ParseTime(node.name)
	for _, e := range node.elements {
		repl, found := (*db)[e.name]
		if found {
			for _, repl := range repl.elements {
				if repl.name == options.single_element {
					acc.Add(repl.name, repl.val*e.val)
				}
			}
		} else {
			if e.name == options.single_element {
				acc.Add(e.name, e.val)
			}
		}
	}
	if len(*acc) > 0 {
		arr := (*acc)[options.single_element]
		if options.csv {
			fmt.Printf("%s;%s;%0.2f;%0.2f;%0.2f\n", ts.Format(OUT_DATE_FMT), options.single_element, arr[ACC_POS], -1*arr[ACC_NEG], arr[ACC_POS]+arr[ACC_NEG])
		} else {
			fmt.Printf("%s %20s %10.2f %10.2f =%10.2f\n", ts.Format(OUT_DATE_FMT), options.single_element, arr[ACC_POS], arr[ACC_NEG], arr[ACC_POS]+arr[ACC_NEG])
		}
	}
}

func SingleFoodProcessor(node Node) {
	ts, _ := ParseTime(node.name)
	for _, e := range node.elements {
		matched, err := regexp.MatchString(options.single_food, e.name)
		if err != nil {
			log.Print(err)
			os.Exit(ERROR_SINGLE_FOOD_NOT_FOUND)
		}
		if matched {
			fmt.Printf("%s\t%s\t%0.2f\n", ts.Format(OUT_DATE_FMT), e.name, e.val)
		}
	}
}

func DefaultProcessor(node Node) {
	acc := NewAccumulator()
	ts, _ := ParseTime(node.name)
	fmt.Printf("%s\n", ts.Format(IN_DATE_FMT))
	for _, e := range node.elements {
		fmt.Printf("\t%-27s :%10.2f\n", e.name, e.val)
		repl, found := (*db)[e.name]
		if found {
			for _, repl := range repl.elements {
				res := repl.val * e.val
				fmt.Printf("\t\t%20s %10.2f\n", repl.name, res)
				acc.Add(repl.name, res)
			}
		} else {
			fmt.Printf("\t\t%20s %10.2f\n", e.name, e.val)
			acc.Add(e.name, e.val)
		}
	}
	if options.totals {
		var ss sort.StringSlice
		if len(*acc) > 0 {
			fmt.Printf("\t-- %s %s\n", "TOTAL ", strings.Repeat("-", 52))
			for name, _ := range *acc {
				ss = append(ss, name)
			}
			sort.Sort(ss)
			for _, name := range ss {
				arr := (*acc)[name]
				fmt.Printf("\t\t%20s %10.2f %10.2f =%10.2f\n", name, arr[ACC_POS], arr[ACC_NEG], arr[ACC_POS]+arr[ACC_NEG])
			}
		}
	}
}
