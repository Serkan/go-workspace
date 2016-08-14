package main

import (
	"log"
	"os"
	"search"
	_"matchers" // even "matchers" package does not called directly from this package (hence the "_" char) this
	// package must be imported to trigger its "init()" function which fills out matcher hash table
	// in search package. Think dependencies as a DAG and main entry point is the root which entry
	// point of the program
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // change log formatting, add line of the source code
}

// Entry point of the program
func main() {
	search.Run("President")
}
