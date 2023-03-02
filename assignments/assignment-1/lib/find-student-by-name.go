package lib

import (
	"assignment-1/data"
	"assignment-1/entity"
	"errors"
	"strings"
)

// function untuk mencari data siswa berdasarkan nama
func FindStudentByName(name string) (entity.Student, error) {
	for _, student := range data.Students {
		if strings.ToLower(student.Nama) == strings.ToLower(name) {
			return student, nil
		}
	}

	return entity.Student{}, errors.New("Siswa tidak ditemukan")
}
