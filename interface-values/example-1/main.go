package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type system struct {
	Host string
}

func (*system) Read() {
}

func (*system) Write() { /* ... */ }

func main() {
	var rw ReadWriter = &system{"127.0.0.1"}
	var r Reader = rw
	fmt.Println(rw, r)
}
