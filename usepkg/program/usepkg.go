package main

import (
	"fmt"
	"tuckerGo/usepkg/custompkg"
	"tuckerGo/usepkg/exinit"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()

	s := custompkg.Student{}
	s.Name = "화랑"
	s.Age = 32
	fmt.Println(s.Name, s.Age)
	fmt.Println(custompkg.PI)
	custompkg.Ratio = 10

	fmt.Println(custompkg.Ratio)
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 1, 2, 3, 10, 4, 6, 5, 3, 1, 2, 9}

	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	exinit.PrintD()
}
