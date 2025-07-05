package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo2/ch14/expkg"
	"goproject/usepkg/custompkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 8, 1, 2, 5, 6, 3, 33, 22, 1}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}
