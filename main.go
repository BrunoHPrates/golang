package main

import "fmt"

func main(){
	var a float32
	var b float32
	fmt.Println("DIgite dois números: ", a, b)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("A soma desses 2 números é: ", a + b)
	fmt.Println("A subtração desses 2 números é: ", a - b)
	fmt.Println("A multiplicação desses 2 números é: ", a * b)
	fmt.Println("A divisão desses 2 números é: ", a / b)
}