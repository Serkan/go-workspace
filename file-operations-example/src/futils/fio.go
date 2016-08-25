package futils

import (
	"bytes"
	"os"
)

var readBuffer = int(1024)

// ToString Dumps given file content to a string and returns it
func ToString(fileName string) (content string, err error) {
	f, e := os.Open(fileName)
	defer f.Close()
	if e != nil {
		return "", e
	}
	var b bytes.Buffer
	buffer := make([]byte, readBuffer)
	for read, err := f.Read(buffer); read != 0 && err == nil; read, err = f.Read(buffer) {
		b.Write(bytes.Trim(buffer, "\x00"))
	}
	return b.String(), err
}

// ContentEquals compares two file's content given by their absolute
// filepaths and returns true if and only if there is no error occurred
// and content of two files are identical
func ContentEquals(f1 string, f2 string) (cmp bool, err error) {
	ff1, e1 := os.Open(f1)
	if e1 != nil {
		return false, e1
	}
	ff2, e2 := os.Open(f2)
	if e2 != nil {
		return false, e2
	}
	defer ff1.Close()
	defer ff2.Close()
	b1 := make([]byte, readBuffer)
	b2 := make([]byte, readBuffer)
	var r1 int
	var r2 int
	r1, err = ff1.Read(b1)
	r2, err = ff2.Read(b2)
	for r1 != 0 && r2 != 0 && err == nil {
		if r1 != r2 {
			return false, nil
		} else {
			if !bytes.Equal(b1, b2) {
				return false, nil
			}
		}
		r1, err = ff1.Read(b1)
		r2, err = ff2.Read(b2)
	}
	if r1 == 0 && r2 == 0 {
		return true, nil
	}
	return true, err
}
