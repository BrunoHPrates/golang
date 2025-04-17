package main

import (
	"fmt"
)

func dadosPessoa(nome string, idade int)(string, int) {
	if idade < 0 {
		return "Idade inválida", 0
	} else if idade >= 0 && idade <= 18 {
		return "Menor de idade", idade
	} else  {
		return nome, idade
	}
}

func main() {
	var nome string
	var idade int
	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)
	fmt.Println("Digite sua idade:")
	fmt.Scanln(&idade)
	fmt.Println(dadosPessoa(nome, idade))

}