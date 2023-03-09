package main

import (
	"errors"
	"fmt"
)

func errorPanicRecover() {
	// num, err := strconv.Atoi("123GH")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(num)
	// }
	//
	// num, err = strconv.Atoi("123")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(num)
	// }

	defer catchErr()

	var password string
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
