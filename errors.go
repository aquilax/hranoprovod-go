package main

import (
	"os"
)

// process exit codes
const (
	exitOk = iota
	exitErrorIO
	exitErrorBadSyntax
	exitErrorConversion
	exitErrorSingleFoodNotFound
	exitErrorParsingOptions
	exitErrorOpeningFile
	exitErrorGeneral
)

// BreakingError contains breaking error data
type BreakingError struct {
	msg      string
	exitCode int
}

// NewBreakingError creates new BreakingError
func NewBreakingError(msg string, exitCode int) *BreakingError {
	return &BreakingError{msg, exitCode}
}

// Error returns the error message
func (e *BreakingError) Error() string {
	return e.msg
}

func handleResult(err error) {
	if err == nil {
		os.Exit(exitOk)
	}
	println(err)
	if bErr, ok := err.(*BreakingError); ok {
		os.Exit(bErr.exitCode)
	}
	os.Exit(exitErrorGeneral)
}
