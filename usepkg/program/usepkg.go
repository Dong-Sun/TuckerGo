package main

import (
	"fmt"
	"tuckerGo/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 1, 2, 3, 10, 4, 6, 5, 3, 1, 2, 9}

	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
