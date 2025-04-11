package main

import "fmt"

func saldoInicial() float64 {
	var saldo float64
	for {
		fmt.Println("Bem-Vindo (a)!")
		fmt.Println("Digite seu saldo inicial:")
		fmt.Scan(&saldo)
		if saldo < 0 {
			fmt.Println("Saldo inválido!")
		} else {
			break
		}
	}
	return saldo
}

func menu() int {
	var opcao int
	fmt.Println("\nEscolha uma opção:")
	fmt.Println("1. Depositar")
	fmt.Println("2. Sacar")
	fmt.Println("3. Ver saldo")
	fmt.Println("4. Sair")
	fmt.Scan(&opcao)
	return opcao
}

func main() {
	saldo := saldoInicial()

	for {
		opcao := menu()
		switch opcao {
		case 1:
			var deposito float64
			fmt.Println("Digite quanto você deseja depositar:")
			fmt.Scan(&deposito)
			if deposito <= 0 {
				fmt.Println("Valor inválido!")
			} else {
				saldo += deposito
				fmt.Println("Depósito realizado. Seu saldo atual é:", saldo)
			}
		case 2:
			var saque float64
			fmt.Println("Digite quanto você deseja sacar:")
			fmt.Scan(&saque)
			if saque > saldo || saque <= 0 {
				fmt.Println("Valor inválido!")
			} else {
				saldo -= saque
				fmt.Println("Saque realizado. Seu saldo atual é:", saldo)
			}
		case 3:
			fmt.Println("Seu saldo atual é:", saldo)
		case 4:
			fmt.Println("Saindo... Obrigado por usar o sistema!")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
