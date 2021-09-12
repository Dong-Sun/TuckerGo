package main

import (
	"fmt"
	"os"
)

type Writer func(string)
type WriterInterface interface {
	write(string)
}

func writeHello(writer Writer) {
	writer("Hello world")
}
func writeHello2(writer WriterInterface) {
	writer.write("Hello world")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
	//writeHello(func(msg string) {
	//fmt.Println(msg)
	//})
}
