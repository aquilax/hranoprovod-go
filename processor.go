package main

import (
	"github.com/Hranoprovod/parser"
	"regexp"
	"sort"
)

const (
	inDateFormat = "2006/01/02"

	dateBeginning = 0
	dateEnd       = 1
)

// Processor contains the processor data
type Processor struct {
	options  *Options
	db       *NodeList
	reporter *Reporter
}

// NewProcessor creates new node processor
func NewProcessor(options *Options, db *NodeList, reporter *Reporter) *Processor {
	return &Processor{
		options,
		db,
		reporter,
	}
}

func (p *Processor) process(node *parser.Node) error {
	options := p.options
	time, err := parseTime(node.Header)
	if err != nil {
		return err
	}
	logNode := NewLogNode(time, node.Elements)

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
		_, found := (*p.db)[e.Name]
		if !found {
			p.reporter.printUnresolvedRow(e.Name)
		}
	}
	return nil
}

func (p *Processor) singleProcessor(logNode *LogNode) error {
	acc := NewAccumulator()
	singleElement := p.options.singleElement
	for _, e := range *logNode.elements {
		repl, found := (*p.db)[e.Name]
		if found {
			for _, repl := range *repl.Elements {
				if repl.Name == singleElement {
					acc.add(repl.Name, repl.Val*e.Val)
				}
			}
		} else {
			if e.Name == singleElement {
				acc.add(e.Name, e.Val)
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
		matched, err := regexp.MatchString(p.options.singleFood, e.Name)
		if err != nil {
			return err
		}
		if matched {
			p.reporter.printSingleFoodRow(logNode.time, e.Name, e.Val)
		}
	}
	return nil
}

func (p *Processor) defaultProcessor(logNode *LogNode) error {
	acc := NewAccumulator()
	p.reporter.printDate(logNode.time)
	for _, element := range *logNode.elements {
		p.reporter.printElement(element)
		if repl, found := (*p.db)[element.Name]; found {
			for _, repl := range *repl.Elements {
				res := repl.Val * element.Val
				p.reporter.printIngredient(repl.Name, res)
				acc.add(repl.Name, res)
			}
		} else {
			p.reporter.printIngredient(element.Name, element.Val)
			acc.add(element.Name, element.Val)
		}
	}
	if p.options.totals {
		var ss sort.StringSlice
		if len(*acc) > 0 {
			p.reporter.printTotalHeader()
			for name := range *acc {
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
