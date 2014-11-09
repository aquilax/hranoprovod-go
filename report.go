package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	OUT_DATE_FMT = "2006/01/02"
)

type Reporter struct {
	options *Options
}

func NewReporter(options *Options) *Reporter {
	return &Reporter{options}
}

func (r *Reporter) cNum(num float32) string {
	if r.options.color {
		if num > 0 {
			return red + fmt.Sprintf("%10.2f", num) + reset
		}
		if num < 0 {
			return green + fmt.Sprintf("%10.2f", num) + reset
		}
	}
	return fmt.Sprintf("%10.2f", num)
}

func (r *Reporter) printDate(ts time.Time) {
	fmt.Printf("%s\n", ts.Format(IN_DATE_FMT))
}

func (r *Reporter) printElement(element *Element) {
	fmt.Printf("\t%-27s :%s\n", element.name, r.cNum(element.val))
}

func (r *Reporter) printIngredient(name string, value float32) {
	fmt.Printf("\t\t%20s %s\n", name, r.cNum(value))
}

func (r *Reporter) printTotalHeader() {
	fmt.Printf("\t-- %s %s\n", "TOTAL ", strings.Repeat("-", 52))
}

func (r *Reporter) printTotalRow(name string, pos float32, neg float32) {
	fmt.Printf("\t\t%20s %s %s =%s\n", name, r.cNum(pos), r.cNum(neg), r.cNum(pos+neg))
}

func (r *Reporter) printSingleElementRow(ts time.Time, name string, pos float32, neg float32, csv bool) {
	format := "%s %20s %10.2f %10.2f =%10.2f\n"
	if csv {
		format = "%s;\"%s\";%0.2f;%0.2f;%0.2f\n"
	}
	fmt.Printf(format, ts.Format(OUT_DATE_FMT), name, pos, -1*neg, pos+neg)
}

func (r *Reporter) printSingleFoodRow(ts time.Time, name string, val float32) {
	fmt.Printf("%s\t%s\t%0.2f\n", ts.Format(OUT_DATE_FMT), name, val)
}
func (r *Reporter) printUnresolvedRow(name string) {
	fmt.Println(name)
}
