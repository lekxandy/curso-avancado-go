package main

import "fmt"

func main() {

	matriz1 := [][]int{
		{1, 2},
		{3, 4},
	}

	matriz2 := make([][]int, len(matriz1))

	copy(matriz2, matriz1)

	matriz2[0][0] = 99 // Modificando a matriz2, mas a matriz1 também é modificada

	fmt.Println(matriz1)
	fmt.Println(matriz2)

	matriz3 := [][]int{
		{1, 2},
		{3, 4},
	}

	matriz4 := make([][]int, len(matriz3))

	for i := range matriz3 {
		matriz4[i] = make([]int, len(matriz3[i]))
		copy(matriz4[i], matriz3[i])
	}
	matriz4[0][0] = 99 // Modificando a matriz4, mas a matriz3 não é modificada

	fmt.Println(matriz3)
	fmt.Println(matriz4)

}
