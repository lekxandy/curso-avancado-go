package main

import "fmt"

func main() {

	fmt.Println("--- Cópia insegura de slices ---")

	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1 //slice2 aponta para o mesmo slice (array interno) que slice1

	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)

	slice1[0] = 10 //modificando o primeiro elemento do slice1

	fmt.Println("Slice1 após modificação: ", slice1)
	fmt.Println("Slice2 após modificação: ", slice2) //slice2 também é afetado, pois aponta para o mesmo slice (array interno)

	slice2 = append(slice2, 6) //adicionando um elemento ao slice2, isso cria um novo slice (array interno) para slice2
	slice2[0] = 99             //modificando o primeiro elemento do slice2, isso não afeta slice1, pois slice2 agora aponta para um novo slice (array interno)

	fmt.Println("Slice1 após append em slice2: ", slice1)
	fmt.Println("Slice2 após append e modificação: ", slice2)

	fmt.Println("--- Cópia segura de slices ---")

	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, len(slice3)) //criando um novo slice com o mesmo tamanho de slice3

	//copiando os elementos de slice3 para slice4
	_ = copy(slice4, slice3)

	fmt.Println("Slice3: ", slice3)
	fmt.Println("Slice4: ", slice4)

	slice3[0] = 10 //modificando o primeiro elemento do slice3

	fmt.Println("Slice3 após modificação: ", slice3)
	fmt.Println("Slice4 após modificação em slice3: ", slice4) //slice4 não é afetado, pois é um slice diferente (array interno diferente)

	slice4 = append(slice4, 6) //adicionando um elemento ao slice4, isso cria um novo slice (array interno) para slice4
	slice4[0] = 99             //modificando o primeiro elemento do slice4, isso não afeta slice3, pois slice4 já é um slice diferente (array interno diferente)

	fmt.Println("Slice3 após append em slice4: ", slice3)
	fmt.Println("Slice4 após append e modificação: ", slice4)

	var slice5 = make([]int, 3)
	copy(slice5, slice3) //copiando os elementos de slice3 para slice5, isso cria um novo slice (array interno) para slice5

	fmt.Println("Slice5: ", slice5)

	var slice6 = make([]int, 20)
	copy(slice6, slice3) //copiando os elementos de slice3 para slice6, isso cria um novo slice (array interno) para slice6 e os elementos restantes de slice6 serão preenchidos com zero (0) (valor zero para int)

	fmt.Println("Slice6: ", slice6)

}
