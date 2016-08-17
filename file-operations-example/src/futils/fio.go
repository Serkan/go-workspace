package futils

import (
	"os"
)

func ToString(fileName string) (content string, err error){
	f, e := os.Open(fileName)
	defer f.Close()
	if e != nil {
		return "", err
	}
	buffer := make([]byte,1024)
	for read, err := f.Read(buffer); read != 0 && err == nil; read, err = f.Read(buffer){
		content += content + string(buffer)
	}
	return content, err
}
