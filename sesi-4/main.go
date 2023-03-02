package main

import (
	"fmt"
	"log"
	"reflect"
)

type LoginDto struct {
	Username string
	Password string
}

func isValid(payload any) {
	t := reflect.TypeOf(payload)
	v := reflect.ValueOf(payload)

	if t.Kind() != reflect.Struct {
		log.Fatalf(
			"isValid can only receive struct data, but when you give is a %v data \n",
			t.Kind(),
		)
	}

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("t.Field(i): %#v \n", t.Field(i))

		fieldValue := v.Field(i).Interface()

		switch value := fieldValue.(type) {
		case string:
			fmt.Println("This is string =>", value)
		case int:
			fmt.Println("This is int    =>", value)
		}
	}
}

func main() {
	loginDto := LoginDto{
		Username: "tfkhdyt",
		Password: "bruh123",
	}

	loginDtoType := reflect.TypeOf(loginDto)

	fmt.Println("loginDto field amount:", loginDtoType.NumField())
	fmt.Println("loginDto type        :", loginDtoType.Name())
	fmt.Println("loginDto kind        :", loginDtoType.Kind())

	isValid(loginDto)
}
