package main

import (
	"fmt"
)

func main() {
	var saldo float64
	var resultado float64
	for {
	fmt.Print("Digite o saldo inicial: ")
	fmt.Scan(&saldo)
	if saldo < 0  {
		fmt.Println("Valor inválido")
	} else { 
		break
	}
	} 
	var opcao string
	for {
	fmt.Println("Você deseja sacar ou depositar? (s/d)")
	fmt.Scan(&opcao)
	if opcao != "s" && opcao != "d" {
		fmt.Println("Opção inválida")
	} else {
		break
	}
}
	if opcao == "s" {
		var saque float64
		for {
		fmt.Println("Digite o valor do saque: ")
		fmt.Scan(&saque)
		if saque > saldo || saque <= 0 {
			fmt.Println("Valor inválido")
		} else {
		resultado = saldo - saque
		fmt.Println("Saque de ", saque," reais realizado com sucesso!")
		fmt.Println("Seu saldo atual é de ", resultado, " reais")
			break
		}
		}
	}
	if opcao == "d" {
		var deposito float64
		for {
		fmt.Println("Digite o valor do deposito: ")
		fmt.Scan(&deposito)
		if deposito <= 0 {
			fmt.Println("Valor inválido")
		} else {
			resultado = saldo + deposito
			fmt.Println("Depósito de ", deposito," reais realizado com sucesso!")
			fmt.Println("Seu saldo atual é de ", resultado, " reais")
			break
		}
		}
	}
}
	