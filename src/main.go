package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//recorrer un map

	for i, v := range m {
		fmt.Println(i, v)

	}
	// los maps realizar concurrencia.

}
