package main

import "fmt"

var nota1 float64
var nota2 float64
var media float64

func analisarNotas (nota1, nota2 float64) (float64, string) {
    media = (nota1 + nota2) / 2
	if media < 0 || media > 10 {
		fmt.Println("Nota inválida! Digite novamente")
		main()
	} else if media >= 6 {
		return media, "Aprovado"
	} else if media < 6 {
		return media, "Reprovado"
	}
	return media, "Nota inválida"
}

func main (){
	fmt.Println("Digite a primeira nota:")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a segunda nota:")
	fmt.Scanln(&nota2)
	media, resultado := analisarNotas(nota1, nota2)
	fmt.Printf("A média é: %.2f\n", media)
	fmt.Println("Resultado: ", resultado)
	}
