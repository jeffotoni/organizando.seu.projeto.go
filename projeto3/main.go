package main

import (
	"fmt"

	"github.com/jeffotoni/gcolorfake"
)

func main() {

	fmt.Println("Projeto 3, entendendo um pouco go mod...")
	// gcolorfake
	color := gcolorfake.YellowCor("My Color Yellow...")
	fmt.Println("Yellow:", color)
}
