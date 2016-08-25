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

func TestContentEqualsSameContent(t *testing.T) {
	abs1, _ := filepath.Abs("../../testdata/content1.txt")
	abs2, _ := filepath.Abs("../../testdata/identicalContent1.txt")
	cmp, err := ContentEquals(abs1, abs2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp {
		t.Fatal("Content of files does not equal")
	}
}
