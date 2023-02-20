package main

import "fmt"

func main() {
	// ini adalah komentar
	// fmt.Println("Hello world")
	// fmt.Println("Hello", "world!", "how", "are", "you")
	// fmt.Print("Hello", "world!", "how", "are", "you")

	var firstName string = "Taufik"
	lastName := "Hidayat"

	fmt.Printf("halo %s %s\n", firstName, lastName)

	first, second, third := "satu", "dua", "tiga"
	fmt.Printf("%s %s %s\n", first, second, third)

	name, _ := "john", "wick"
	fmt.Println(name)
}
