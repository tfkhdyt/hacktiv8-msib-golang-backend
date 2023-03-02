package lib

import "errors"

func ParseNameFromArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("Argumen tidak valid, masukkan minimal 1 nama")
	}

	return args[1], nil
}
