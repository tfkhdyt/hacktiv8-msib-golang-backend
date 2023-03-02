package main

import (
	"assignment-1/lib"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Argumen tidak valid, masukkan minimal 1 nama")
	}

	name := args[1]

	student, err := lib.FindStudentByName(name)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("ID        : %d\n", student.Id)
	fmt.Printf("Nama      : %s\n", student.Nama)
	fmt.Printf("Alamat    : %s\n", student.Alamat)
	fmt.Printf("Pekerjaan : %s\n", student.Pekerjaan)
	fmt.Printf("Alasan    : %s\n", student.Alasan)
}
