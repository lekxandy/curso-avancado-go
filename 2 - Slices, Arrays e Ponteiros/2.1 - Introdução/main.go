package main

import "fmt"

func main() {

	var (
		slice = []int{1, 2, 3, 4, 5}  //tamanho dinâmico
		array = [5]int{1, 2, 3, 4, 5} //tamanho fixo
	)

	fmt.Println("Tamanho do array: ", len(array))
	fmt.Println("Capacidade do array: ", cap(array))

	fmt.Println("Tamanho do slice: ", len(slice))
	fmt.Println("Capacidade do slice: ", cap(slice))

	slice = append(slice, 6) //adiciona um elemento ao slice

	fmt.Println("Slice após adicionar um elemento: ", slice)
	fmt.Println("Tamanho do slice após adicionar um elemento: ", len(slice))
	fmt.Println("Capacidade do slice após adicionar um elemento: ", cap(slice))

	printSlice(slice)

	// append nao funciona com arrays, pois eles tem tamanho fixo
	// array = append(array, 6) //isso vai dar erro
	// append só funciona com slices

}

func printSlice(slice []int) {
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
