package futils

import (
	"path/filepath"
	"testing"
)

func TestToString(t *testing.T) {
	absPathOfTestFile, _ := filepath.Abs("../../testdata/test_text_short.txt")
	s, err := ToString(absPathOfTestFile)
	if err != nil {
		t.Fatal(err)
	}
	expected := string("A")
	if s != expected {
		t.Fatalf("Result does not match Actual:%s# Expected:%s# ActualLen: %d ExpectedLen: %d", s, expected, len(s), len(expected))
	}
}
