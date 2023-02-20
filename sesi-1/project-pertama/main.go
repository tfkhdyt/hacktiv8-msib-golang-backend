package main

import "fmt"

func main() {
	// ini adalah komentar
	fmt.Println("Hello world")
	// fmt.Println("Hello", "world!", "how", "are", "you")
	// fmt.Print("Hello", "world!", "how", "are", "you")

	// var firstName string = "Taufik"
	// lastName := "Hidayat"
	//
	// fmt.Printf("halo %s %s\n", firstName, lastName)
	//
	// first, second, third := "satu", "dua", "tiga"
	// fmt.Printf("%s %s %s\n", first, second, third)
	//
	// name, _ := "john", "wick"
	// fmt.Println(name)

	var positiveNumber uint8 = 89
	negativeNumber := -1243423644
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	decimalNumber := 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	bool := true
	fmt.Printf("exist? %t\n", bool)

	message := `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`
	fmt.Println(message)
}
