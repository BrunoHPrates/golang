package main

import "fmt"

func main(){
	var ages = [4]int{15, 16, 16, 16}
	nomes := [4]string{"Yan", "Eduardo", "Pedro", "Vini"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] = "Miles Morales"
	fmt.Println(nomes)
}