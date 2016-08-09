package main

import (
	"log"
	"os"
	"search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("President")
}
