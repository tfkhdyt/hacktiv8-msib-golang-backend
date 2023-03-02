package lib

import (
	"assignment-1/data"
	"assignment-1/domain"
	"errors"
	"strings"
)

// function untuk mencari data siswa berdasarkan nama
func FindStudentByName(name string) (domain.Student, error) {
	for _, student := range data.Students {
		if strings.ToLower(student.Nama) == strings.ToLower(name) {
			return student, nil
		}
	}

	return domain.Student{}, errors.New("Siswa tidak ditemukan")
}
