package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

func main() {
	var stringer Stringer
	s := stringer.(*Student) // error
	fmt.Println(s)
}
