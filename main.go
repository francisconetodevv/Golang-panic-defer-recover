package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	_, err := os.Open("F:/Golang Estudos/Panic/settings.txt")
	/* Tratamento do erro */
	if err != nil {
		/* Para tudo que está fazendo */
		panic("FileNotExist")
	}
}

func main() {
	/* Funcao anônima */
	/* Recovery = Não muito utilizado no Golang */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Tratativa para recuperação")
		}
	}()

	ReadFile()

	fmt.Println("Fim")
}
