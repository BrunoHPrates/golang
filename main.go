package main

import ("fmt")

func main(){
	numeros:= [10]float32{}
	fmt.Println("Digite 5 números números.")
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&numeros[0])
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&numeros[1])
	fmt.Println("Digite o terceiro número:")
	fmt.Scan(&numeros[2])
	fmt.Println("Digite o quarto número:")
	fmt.Scan(&numeros[3])
	fmt.Println("Digite o quinto número:")
	fmt.Scan(&numeros[4])
	fmt.Println("A soma dos números é", numeros[0]+numeros[1]+numeros[2]+numeros[3]+numeros[4])
}