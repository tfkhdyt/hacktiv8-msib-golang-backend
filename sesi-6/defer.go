package main

import (
	"fmt"
	"os"
)

func _defer() {
	// defer fmt.Println("defer function starts to execute")
	// fmt.Println("Hi everyone")
	// fmt.Println("Welcome back to Go learning center")

	// callDeferFunc()
	// fmt.Println("Hi everyone")

	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
