package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"time"
)

const (
	IN_DATE_FMT = "2006/01/02"

	DATE_START = 1
	DATE_STOP  = 2
)

func processor(node *Node) {
	if len(options.beginning) > 0 {
		if !GoodDate(node.name, options.beginning, DATE_START) {
			return
		}
	}
	if len(options.end) > 0 {
		if !GoodDate(node.name, options.end, DATE_STOP) {
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
	if ctype == DATE_START {
		return ts.Unix() >= tsb.Unix()
	}
	return ts.Unix() <= tsb.Unix()
}

func UnresolvedProcessor(node Node) {
	for _, e := range node.elements {
		_, found := (*db)[e.name]
		if !found {
			printUnresolvedRow(e.name)
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
		printSingleElementRow(ts, options.single_element, arr[ACC_POS], arr[ACC_NEG], options.csv)
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
			printSingleFoodRow(ts, e.name, e.val)
		}
	}
}

func DefaultProcessor(node Node) {
	acc := NewAccumulator()
	ts, _ := ParseTime(node.name)
	printDate(ts)
	for _, element := range node.elements {
		printElement(element)
		repl, found := (*db)[element.name]
		if found {
			for _, repl := range repl.elements {
				res := repl.val * element.val
				printIngredient(repl.name, res)
				acc.Add(repl.name, res)
			}
		} else {
			printIngredient(element.name, element.val)
			acc.Add(element.name, element.val)
		}
	}
	if options.totals {
		var ss sort.StringSlice
		if len(*acc) > 0 {
			printTotalHeader()
			for name, _ := range *acc {
				ss = append(ss, name)
			}
			sort.Sort(ss)
			for _, name := range ss {
				arr := (*acc)[name]
				printTotalRow(name, arr[ACC_POS], arr[ACC_NEG])
			}
		}
	}
}
