package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando")
	application := app.Generate()

	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
