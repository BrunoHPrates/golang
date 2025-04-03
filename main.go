package main

import "fmt"

func main(){
	var ages = [4]int{15, 16, 16, 16}
	nomes := [4]string{"Yan", "Eduardo", "Pedro", "Vini"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] = "Miles Morales"
	fmt.Println(nomes)
	// Slice
	var score = []int{100, 200, 300, 400}
	fmt.Println(score)
	score[1] = 2
	fmt.Println(score, len(score), cap(score))
	rangeOne := score[1:3]
	fmt.Println(rangeOne)
	rangetwo := score[2:]
	fmt.Println(rangetwo)
	rangeThree := score[:3]
	fmt.Println(rangeThree)

	var superherois = []string{"Deadpool", "Homem -aranha", "Motoqueiro Fantasma"}
	fmt.Println(superherois)
	superherois = append(superherois, "Ben 10", "Mulher Maravilha")
	fmt.Println(superherois, len(superherois), cap(superherois))
}