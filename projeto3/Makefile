# Makefile
.EXPORT_ALL_VARIABLES:	

#GO111MODULE=on
#GOPROXY=https://proxy.golang.org
#GONOSUMDB=off
#GOPRIVATE=*. internal.mycompany.com

build:
	rm -f go.*
	go mod init gitprojetos/organizando.seu.projeto.go/projeto3
	go mod tidy

	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o projeto3 main.go
	@echo "buid completo..."
