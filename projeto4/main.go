package main

import (
	"fmt"

	"github.com/jeffotoni/gcolorfake2"
)

func main() {

	fmt.Println("Projeto 3, entendendo um pouco go mod...")
	// gcolorfake
	color := gcolorfake2.Yellow("My Color Yellow...")
	fmt.Println("Yellow:", color)
}
