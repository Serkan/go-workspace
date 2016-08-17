package main

import (
	"futils"
	"fmt"
)

func main() {
	s, _ := futils.ToString("testdata/test_text_short.txt")
	fmt.Println(s == "This is a text")
}