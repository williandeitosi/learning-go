package main

import "fmt"

// posso declarar CONSTANTE todas juntas ou separadas
const (
	name  string = "Willian"
	ideia string = ";f;lajdf;lkaj"
)

// posso colocar tipagem ou nao
const nick = "Setth"

// posso declarar VARIAVEL todas juntas ou separadas
var age int = 28

func main() {
	fmt.Print("teste: ", name, "\n")
	hello()
}

func hello() {
	lastName := "Giovanini"

	fmt.Print("sobre nome: ", lastName, " idade: ", age)

}
