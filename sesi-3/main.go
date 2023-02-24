package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Taufik", "Hidayat"}
	printMessage("Halo", names)
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
