package main

import (
	"regexp"
	"sort"
)

const (
	IN_DATE_FMT = "2006/01/02"

	dateBeginning = 0
	dateEnd       = 1
)

type Processor struct {
	options  *Options
	db       *NodeList
	reporter *Reporter
}

func NewProcessor(options *Options, db *NodeList, reporter *Reporter) *Processor {
	return &Processor{
		options,
		db,
		reporter,
	}
}

func (p *Processor) process(node *Node) error {
	options := p.options
	time, err := parseTime(node.header)
	if err != nil {
		return err
	}
	logNode := NewLogNode(time, node.elements)

	if (options.hasBeginning && !isGoodDate(logNode.time, options.beginningTime, dateBeginning)) || (options.hasEnd && !isGoodDate(logNode.time, options.endTime, dateEnd)) {
		return nil
	}

	if options.unresolved {
		return p.unresolvedProcessor(logNode)
	}
	if len(options.singleElement) > 0 {
		return p.singleProcessor(logNode)
	}
	if len(options.singleFood) > 0 {
		return p.singleFoodProcessor(logNode)
	}
	return p.defaultProcessor(logNode)

}

func (p *Processor) unresolvedProcessor(logNode *LogNode) error {
	for _, e := range *logNode.elements {
		_, found := (*p.db)[e.name]
		if !found {
			p.reporter.printUnresolvedRow(e.name)
		}
	}
	return nil
}

func (p *Processor) singleProcessor(logNode *LogNode) error {
	acc := NewAccumulator()
	singleElement := p.options.singleElement
	for _, e := range *logNode.elements {
		repl, found := (*p.db)[e.name]
		if found {
			for _, repl := range *repl.elements {
				if repl.name == singleElement {
					acc.Add(repl.name, repl.val*e.val)
				}
			}
		} else {
			if e.name == singleElement {
				acc.Add(e.name, e.val)
			}
		}
	}
	if len(*acc) > 0 {
		arr := (*acc)[singleElement]
		p.reporter.printSingleElementRow(logNode.time, singleElement, arr[accPos], arr[accNeg], p.options.csv)
	}
	return nil
}

func (p *Processor) singleFoodProcessor(logNode *LogNode) error {
	for _, e := range *logNode.elements {
		matched, err := regexp.MatchString(p.options.singleFood, e.name)
		if err != nil {
			return err
		}
		if matched {
			p.reporter.printSingleFoodRow(logNode.time, e.name, e.val)
		}
	}
	return nil
}

func (p *Processor) defaultProcessor(logNode *LogNode) error {
	acc := NewAccumulator()
	p.reporter.printDate(logNode.time)
	for _, element := range *logNode.elements {
		p.reporter.printElement(element)
		if repl, found := (*p.db)[element.name]; found {
			for _, repl := range *repl.elements {
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
				p.reporter.printTotalRow(name, arr[accPos], arr[accNeg])
			}
		}
	}
	return nil
}
