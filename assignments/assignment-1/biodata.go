package main

import (
	"assignment-1/lib"
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

	lib.PrintOutput(student)
}
