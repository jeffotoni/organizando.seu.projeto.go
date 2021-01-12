package main

import (
	"fmt"

	"github.com/jeffotoni/organizando.seu.projeto.go/internal/crypt"
	"github.com/jeffotoni/organizando.seu.projeto.go/internal/fmts"

	"github.com/jeffotoni/gcolor"
)

func main() {

	fmt.Println("Projeto 1, entendendo um pouco go mod...")

	// sha1
	sha1 := crypt.Sha1("vamos passar algo aqui para ele!!!")

	// gerar uma crc32
	cr32 := crypt.Crc32() // gera automatico

	fmt.Println(sha1)
	fmt.Println(cr32)

	// concat
	text := fmts.Concat("jeff", "otoni")
	fmt.Println(text)

	// gcolor
	color := gcolor.CyanCor("My Color Cyan...")
	fmt.Println("Cyan:", color)
}
