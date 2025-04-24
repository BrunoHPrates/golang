package main

import "fmt"

func main (){
	capitais := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
		"ES": "Espírito Santo",
	}

	capitais["BH"] = "Belo Horizonte"

	for k,v := range capitais {
		fmt.Println("Sigla, Nome", k, v)
	}

	delete(capitais, "AC")

	for k,v := range capitais {
		fmt.Println("Sigla, Nome", k, v)
	}
}