package lib

import "errors"

// function untuk mendapatkan nama siswa yang ingin dicari dari argumen-argumen
func ParseNameFromArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("Argumen tidak valid, masukkan minimal 1 nama")
	}

	return args[1], nil
}
