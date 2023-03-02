package main

import (
	"assignment-1/lib"
	"fmt"
	"log"
	"os"
)

func main() {
	name, err := lib.ParseNameFromArgs(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}

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
