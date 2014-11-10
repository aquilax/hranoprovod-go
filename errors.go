package main

const (
	EXIT_OK = iota
	ERROR_IO
	ERROR_BAD_SYNTAX
	ERROR_CONVERSION
	ERROR_SINGLE_FOOD_NOT_FOUND
	ERROR_PARSING_OPTIONS
	ERROR_OPENING_FILE
)

type BreakingError struct {
	msg      string
	exitCode int
}

func NewBreakingError(msg string, exitCode int) *BreakingError {
	return &BreakingError{msg, exitCode}
}

func (e *BreakingError) Error() string {
	return e.msg
}
