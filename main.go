package main

const (
	VERSION = "0.1.3"

	EXIT_OK = iota
	ERROR_IO
	ERROR_BAD_SYNTAX
	ERROR_CONVERSION
	ERROR_SINGLE_FOOD_NOT_FOUND
)
func main() {
	NewHranoprovod().run()
}