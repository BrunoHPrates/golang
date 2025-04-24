package main

import "fmt"

func main (){
	estoque := make(map[string]int)

	estoque["coxinha"] = 10
	estoque["pâo de queijo"] = 15
	estoque["refrigerante"] = 20	

	for k,v := range estoque {
	fmt.Println("Produto, Quantidade", k, v)
	}
	
	estoque["coxinha"] -= 2
	estoque["pâo de queijo"] -= 1
	 
	fmt.Printf("\nProdutos após a venda: \n")
	for k,v := range estoque {
		fmt.Println("Produto, Quantidade", k, v)
		}

}
