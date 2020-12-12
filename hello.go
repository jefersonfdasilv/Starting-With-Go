package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Jeferson"
	var peso = 74.5
	var pesoAtual float32 = 77.0
	fmt.Println("Olá sr.", nome)
	fmt.Println("Peso em 2019", peso)
	fmt.Println("O tipo de dado da var peso é:", reflect.TypeOf(peso))
	fmt.Println("O tipo de dado da var peso atual é:", reflect.TypeOf(pesoAtual))
}
