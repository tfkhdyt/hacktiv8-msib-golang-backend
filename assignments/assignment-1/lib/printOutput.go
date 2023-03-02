package lib

import (
	"assignment-1/entity"
	"fmt"
)

// function untuk mencetak data siswa yang sudah ditemukan
func PrintOutput(student entity.Student) {
	fmt.Printf("ID        : %d\n", student.Id)
	fmt.Printf("Nama      : %s\n", student.Nama)
	fmt.Printf("Alamat    : %s\n", student.Alamat)
	fmt.Printf("Pekerjaan : %s\n", student.Pekerjaan)
	fmt.Printf("Alasan    : %s\n", student.Alasan)
}
