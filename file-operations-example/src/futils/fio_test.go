package futils

import (
	"testing"
)

func TestToString(t *testing.T) {
	s, err := ToString("testdata/test_text_short.txt")
	if err != nil || s != "This is a text" {
		t.Fail()
	}
}
