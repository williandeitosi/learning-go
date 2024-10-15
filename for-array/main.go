package main

import "fmt"

func main() {
	meuArray := [3]int{100, 22, 30}

	for i, v := range meuArray {
		fmt.Printf("Indice: %d valor: %d ", i, v)
	}

	var outroArray [5]string
	outroArray[0] = "Willian"
	outroArray[1] = "Rafaela"
	outroArray[2] = "Jhon"
	outroArray[3] = "Doe"
	outroArray[4] = "Rick"

	for i, v := range outroArray {
		fmt.Printf("Indice: --%d-- nome: %v ", i, v)
	}
}
