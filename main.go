package main

import ("fmt" 
		"strings"
		"sort")

func main(){
	greeting := "Hello my friends!"
	fmt.Println(strings.Contains(greeting, "friends"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"my"))
	fmt.Println(strings.Split(greeting, " "))
	ages := []int {50, 80, 10}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 50)
	fmt.Println(index)
	names := []string{"Bruno", "Yan", "Eduardo"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Yan"))
	x := 0

	for x < 5 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("for 2", i)
	}

	for i:= 0; i <len(names); i++ {
		fmt.Println(names[1])
	}

	for index, value := range names{
		fmt.Println("O index é", index, "e o valor é", value)
	}

	for indice, valor := range ages {
		fmt.Println("O ", indice, "é ", valor)
	}

	superherois := []string{"Homem Aranha", "Pantera Negra", "Homem de Ferro"}
	for index, value := range superherois {
		fmt.Println("O número do herói:", index, "e o nome do herói: ", value)
	}
}