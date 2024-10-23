package main

import "fmt"

func main() {
	s := []int{1, 20, 35, 69, 100}

	primitivo := struct{ name string }{name: "willian"}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
