package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/organizando.seu.projeto.go/projeto2/internal/crypt"
	"github.com/jeffotoni/organizando.seu.projeto.go/projeto2/internal/fmts"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/ping", Ping)

	muxhttp := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Run Server:8080")
	log.Fatal(muxhttp.ListenAndServe())

}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"pong"}`))
}
