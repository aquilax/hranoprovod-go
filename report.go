package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	OUT_DATE_FMT = "2006/01/02"
)

func cNum(num float32) string {
	if options.color {
		if num > 0 {
			return red + fmt.Sprintf("%10.2f", num) + reset
		}
		if num < 0 {
			return green + fmt.Sprintf("%10.2f", num) + reset
		}
	}
	return fmt.Sprintf("%10.2f", num)
}

func printDate(ts time.Time) {
	fmt.Printf("%s\n", ts.Format(IN_DATE_FMT))
}

func printElement(element Element) {
	fmt.Printf("\t%-27s :%s\n", element.name, cNum(element.val))
}

func printIngredient(name string, value float32) {
	fmt.Printf("\t\t%20s %s\n", name, cNum(value))
}

func printTotalHeader() {
	fmt.Printf("\t-- %s %s\n", "TOTAL ", strings.Repeat("-", 52))
}

func printTotalRow(name string, pos float32, neg float32) {
	fmt.Printf("\t\t%20s %s %s =%s\n", name, cNum(pos), cNum(neg), cNum(pos+neg))
}

func printSingleElementRow(ts time.Time, name string, pos float32, neg float32, csv bool) {
	format := "%s %20s %10.2f %10.2f =%10.2f\n"
	if csv {
		format = "%s;\"%s\";%0.2f;%0.2f;%0.2f\n"
	}
	fmt.Printf(format, ts.Format(OUT_DATE_FMT), name, pos, -1*neg, pos+neg)
}

func printSingleFoodRow(ts time.Time, name string, val float32) {
	fmt.Printf("%s\t%s\t%0.2f\n", ts.Format(OUT_DATE_FMT), name, val)
}
func printUnresolvedRow(name string) {
	fmt.Println(name)
}
