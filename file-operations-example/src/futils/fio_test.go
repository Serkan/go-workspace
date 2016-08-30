package futils

import (
	"crypto/sha256"
	"io/ioutil"
	"os"
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

func TestContentEqualsDiffContent(t *testing.T) {
	abs1, _ := filepath.Abs("../../testdata/content1.txt")
	abs2, _ := filepath.Abs("../../testdata/content2.txt")
	cmp, err := ContentEquals(abs1, abs2)
	if err != nil {
		t.Fatal(err)
	}
	if cmp {
		t.Fatal("Content of files equals")
	}
}

func TestCopyFile(t *testing.T) {
	filename := "largeFileToCopy"
	abs, _ := filepath.Abs("../../testdata/" + filename)
	tmpdir, _ := ioutil.TempDir("", "")
	err := CopyFile(abs, tmpdir)
	if err != nil {
		t.Fatal(err)
	}
	// hash original file
	from, err := os.Open(abs)
	content, _ := ioutil.ReadAll(from)
	h1 := sha256.Sum256(content)
	// hash copy
	destFileName := tmpdir + "/" + filename
	to, err := os.Open(tmpdir + "/" + filename)
	content, _ = ioutil.ReadAll(to)
	h2 := sha256.Sum256(content)
	if h1 != h2 {
		t.Fatalf("Files does not match source:%s destination:%s", abs, destFileName)
	}
}

func TestCCopyFile(t *testing.T) {
	filename := "largeFileToCopy"
	abs, _ := filepath.Abs("../../testdata/" + filename)
	tmpdir, _ := ioutil.TempDir("", "")
	err := CCopyFile(abs, tmpdir)
	if err != nil {
		t.Fatal(err)
	}
	// hash original file
	from, err := os.Open(abs)
	content, _ := ioutil.ReadAll(from)
	h1 := sha256.Sum256(content)
	sourceSize := len(content)
	// hash copy
	destFileName := tmpdir + "/" + filename
	to, err := os.Open(tmpdir + "/" + filename)
	content, _ = ioutil.ReadAll(to)
	destinationSize := len(content)
	h2 := sha256.Sum256(content)
	if h1 != h2 {
		t.Fatalf("Files does not match source:%s destination:%s source_size:%d destination_size:%d", abs, destFileName, sourceSize, destinationSize)
	}
}

// TODO write a benchmark to compare concurrent and sequential version of CopyFile
