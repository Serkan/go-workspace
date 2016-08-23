package futils

import (
	"bytes"
	"os"
	"strings"
)

// ToString Dumps given file content to a string and returns it
func ToString(fileName string) (content string, err error) {
	f, e := os.Open(fileName)
	defer f.Close()
	if e != nil {
		return "", e
	}
	var b bytes.Buffer
	buffer := make([]byte, 1024)
	for read, err := f.Read(buffer); read != 0 && err == nil; read, err = f.Read(buffer) {
		b.Write(buffer)
	}
	content = strings.Trim(string(buffer), "\x00")
	return content, err
}
