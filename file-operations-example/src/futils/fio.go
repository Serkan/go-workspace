package futils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
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

// CopyFile Copies the file given with first parameter (absolute path of the file)
// to under the directory given with second parameter.
func CopyFile(filename string, dirname string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return err

	}
	name := filename[strings.LastIndex(filename, "/"):]
	newfile, err := os.Create(dirname + "/" + name)
	defer newfile.Close()
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	var e error
	for read, e := buffer.ReadFrom(f); ; read, e = buffer.ReadFrom(f) {
		if read != 0 && e == nil {
			buffer.WriteTo(newfile)
		} else {
			break
		}
	}
	if e != io.EOF {
		return e
	}
	return nil
}

// CCopyFile Concurrently copies given file to under given directory, a go routine reads the
// file in a channel and function do the writing
func CCopyFile(filename string, dirname string) error {
	from, err := os.Open(filename)
	defer from.Close()
	if err != nil {
		return nil

	}
	name := filename[strings.LastIndex(filename, "/"):]
	to, err := os.Create(dirname + "/" + name)
	if err != nil {
		return nil
	}
	pipe := make(chan *[]byte, 10)
	go func(pipe chan *[]byte, file *os.File) {
		defer file.Close()
		// read in a buffer and pass it to channel
		for {
			buffer := make([]byte, 1024) // create byte array to read in and send this byte array to channel
			read, err := file.Read(buffer)
			if read == 0 || err != nil {
				break
			}
			pipe <- &buffer
		}
		if err != nil && err != io.EOF {
			// signal error
			pipe <- nil
		}
		close(pipe)
	}(pipe, from)
	// write to the file from channel
	for buffer := range pipe {
		if buffer != nil {
			to.Write(*buffer)
		} else {
			return errors.New("Error while reading/writing")
		}
	}
	return nil
}

// CopyDir This function copies all content of the specified source directory to destination directory
func CopyDir(sourcedir string, targetdir string) error {
	e := os.MkdirAll(targetdir, 0777)
	if e != nil {
		return e
	}
	f, err := os.Open(sourcedir)
	if err != nil {
		return err
	}
	fiSlice, err := f.Readdir(0)
	for _, fi := range fiSlice {
		if fi.IsDir() {
			err = CopyDir(sourcedir+"/"+fi.Name(), targetdir+"/"+fi.Name())
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourcedir+"/"+fi.Name(), targetdir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// WriteToFile writes given byte stream to specified file
func WriteToFile(reader io.Reader, filename string) error {
	buffer := make([]byte, readBuffer)
	f, err := os.OpenFile(filename, os.O_RDWR, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	for read, e := reader.Read(buffer); read > 0 && e == nil; read, e = reader.Read(buffer) {
		_, we := f.Write(buffer[0:read])
		if we != nil {
			return we
		}
	}
	return nil
}

// Grep It greps a regex pattern from given file
func Grep(pattern string, filename string) (line []string, err error) {
	exp, e1 := regexp.Compile(pattern)
	if e1 != nil {
		return nil, e1
	}
	f, e2 := os.Open(filename)
	if e2 != nil {
		return nil, e2
	}
	bufreader := bufio.NewReader(f)
	var result []string
	for line, e4 := bufreader.ReadString('\n'); e4 == nil; line, e4 = bufreader.ReadString('\n') {
		match := exp.Find([]byte(line))
		if match != nil {
			result = append(result, string(line))
		}
	}
	return result, nil
}

// StringReader implementation of io.Reader interface which reads from a given string
type StringReader struct {
	data string
	pos  int
}

func (readerptr *StringReader) Read(p []byte) (n int, err error) {
	// if pos over data return EOF
	// else copy from pos to pos + len(p) and set p tp pos + len(p)
	bufferSize := len(p)
	dataSize := len(readerptr.data)
	if readerptr.pos >= dataSize {
		return 0, io.EOF
	}
	remaining := dataSize - readerptr.pos
	byteArray := []byte(readerptr.data)
	var read int
	if remaining <= bufferSize {
		copy(p, byteArray[readerptr.pos:dataSize])
		read = dataSize - readerptr.pos
		readerptr.pos = dataSize
	} else {
		copy(p, byteArray[readerptr.pos:readerptr.pos+bufferSize-1])
		readerptr.pos += bufferSize
		read = bufferSize
	}
	return read, nil
}
