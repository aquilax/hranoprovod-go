package main

import (
	"strings"
	"time"
)

const (
	reset       = "\x1B[0m"
	bold        = "\x1B[1m"
	dim         = "\x1B[2m"
	under       = "\x1B[4m"
	reverse     = "\x1B[7m"
	hide        = "\x1B[8m"
	clearscreen = "\x1B[2J"
	clearline   = "\x1B[2K"
	black       = "\x1B[30m"
	red         = "\x1B[31m"
	green       = "\x1B[32m"
	yellow      = "\x1B[33m"
	blue        = "\x1B[34m"
	magenta     = "\x1B[35m"
	cyan        = "\x1B[36m"
	white       = "\x1B[37m"
	bblack      = "\x1B[40m"
	bred        = "\x1B[41m"
	bgreen      = "\x1B[42m"
	byellow     = "\x1B[43m"
	bblue       = "\x1B[44m"
	bmagenta    = "\x1B[45m"
	bcyan       = "\x1B[46m"
	bwhite      = "\x1B[47m"
	newline     = "\r\n\x1B[0m"
)

func mytrim(s string) string {
	return strings.Trim(s, "\t \n:")
}

func parseTime(date string) (time.Time, error) {
	return time.Parse(inDateFormat, mytrim(date))
}

func isGoodDate(time, compareTime time.Time, compareType int) bool {
	if time.Equal(compareTime) {
		return true
	}
	if compareType == dateBeginning {
		return time.After(compareTime)
	}
	return time.Before(compareTime)
}

func printError(err error) {
	print(err)
}
