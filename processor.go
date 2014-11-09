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

type Processor struct {
	options  *Options
	db       *NodeList
	reporter *Reporter
}

func NewProcessor(options *Options, db *NodeList) *Processor {
	return &Processor{
		options,
		db,
		NewReporter(options),
	}
}

func (p *Processor) process(node *Node) {
	options := p.options
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
		p.unresolvedProcessor(*node)
		return
	}
	if len(options.single_element) > 0 {
		p.singleProcessor(*node)
		return
	}
	if len(options.single_food) > 0 {
		p.singleFoodProcessor(*node)
		return
	}
	p.defaultProcessor(*node)
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

func (p *Processor) unresolvedProcessor(node Node) {
	for _, e := range node.elements {
		_, found := (*p.db)[e.name]
		if !found {
			printUnresolvedRow(e.name)
		}
	}
}

func (p *Processor) singleProcessor(node Node) {
	acc := NewAccumulator()
	ts, _ := ParseTime(node.name)
	for _, e := range node.elements {
		repl, found := (*p.db)[e.name]
		if found {
			for _, repl := range repl.elements {
				if repl.name == p.options.single_element {
					acc.Add(repl.name, repl.val*e.val)
				}
			}
		} else {
			if e.name == p.options.single_element {
				acc.Add(e.name, e.val)
			}
		}
	}
	if len(*acc) > 0 {
		arr := (*acc)[p.options.single_element]
		p.reporter.printSingleElementRow(ts, p.options.single_element, arr[ACC_POS], arr[ACC_NEG], p.options.csv)
	}
}

func (p *Processor) singleFoodProcessor(node Node) {
	ts, _ := ParseTime(node.name)
	for _, e := range node.elements {
		matched, err := regexp.MatchString(p.options.single_food, e.name)
		if err != nil {
			log.Print(err)
			os.Exit(ERROR_SINGLE_FOOD_NOT_FOUND)
		}
		if matched {
			printSingleFoodRow(ts, e.name, e.val)
		}
	}
}

func (p *Processor) defaultProcessor(node Node) {
	acc := NewAccumulator()
	ts, _ := ParseTime(node.name)
	p.reporter.printDate(ts)
	for _, element := range node.elements {
		p.reporter.printElement(element)
		repl, found := (*p.db)[element.name]
		if found {
			for _, repl := range repl.elements {
				res := repl.val * element.val
				p.reporter.printIngredient(repl.name, res)
				acc.Add(repl.name, res)
			}
		} else {
			p.reporter.printIngredient(element.name, element.val)
			acc.Add(element.name, element.val)
		}
	}
	if p.options.totals {
		var ss sort.StringSlice
		if len(*acc) > 0 {
			p.reporter.printTotalHeader()
			for name, _ := range *acc {
				ss = append(ss, name)
			}
			sort.Sort(ss)
			for _, name := range ss {
				arr := (*acc)[name]
				p.reporter.printTotalRow(name, arr[ACC_POS], arr[ACC_NEG])
			}
		}
	}
}
