package main

import "fmt"

func main() {
	//arrays
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[3:])

	//append

	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	//append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// recorrer un array o un slice

	slice2 := []string{"Vamos", "a ver", "que pasa"}

	for i, valor := range slice2 {
		fmt.Println(i, valor)
	}

	//

}
