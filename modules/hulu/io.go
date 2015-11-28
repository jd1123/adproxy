package hulu

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// struct to implement io.ReadCloser
type ClosingBuffer struct {
	io.Reader
}

func (cb ClosingBuffer) Close() (err error) {
	return
}

func LoadFakeResponse() string {
	f, err := os.OpenFile("/etc/adproxy/dat/response", os.O_RDONLY, 0660)
	if err != nil {
		fmt.Println("File Error: ", err)
		os.Exit(1)
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("File Error: ", err)
		os.Exit(1)
	}
	return string(buff)
}

func LoadFilterList() []string {
	f, err := os.OpenFile("/etc/adproxy/dat/filterlist.txt", os.O_RDONLY, 0660)
	if err != nil {
		fmt.Println("File Errror: ", err)
		os.Exit(1)
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("File Error: ", err)
		os.Exit(1)
	}
	ns := make([]string, 0)
	s := strings.Split(string(buff), "\n")
	for i := range s {
		if s[i] != "" {
			ns = append(ns, s[i])
		}
	}
	return ns

}

func LogToFile(filename string, b []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(b); err != nil {
		return err
	}
	return nil
}
