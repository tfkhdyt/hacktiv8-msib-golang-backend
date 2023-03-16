package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type user struct {
	FullName string `json:"full_name"`
	Age      uint   `json:"age"`
}

func myJson() {
	// jsonString := `
	//    {
	//      "full_name": "Taufik Hidayat",
	//      "email": "tfkhdyt@proton.me",
	//      "age": 20
	//    }
	//  `
	// jsonString := `[
	//    {
	//      "full_name": "Taufik Hidayat",
	//      "email": "tfkhdyt@proton.me",
	//      "age": 20
	//    },
	//    {
	//      "full_name": "M. Fauzi Fathirohman",
	//      "email": "tafanizer@gmail.com",
	//      "age": 10
	//    }
	//  ]`
	object := []user{{"john wick", 27}, {"ethan hunt", 32}}

	// var result employee
	// var result map[string]any
	// var temp any
	// var result []employee

	// if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// if err := json.Unmarshal([]byte(jsonString), &temp); err != nil {
	// 	log.Fatalln(err.Error())
	// }
	//
	// result := temp.(map[string]any)

	jsonData, err := json.Marshal(object)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// fmt.Println("full_name:", result.FullName)
	// fmt.Println("email:", result.Email)
	// fmt.Println("age:", result.Age)
	// fmt.Println("full_name:", result["full_name"])
	// fmt.Println("email:", result["email"])
	// fmt.Println("age:", result["age"])
	// for i, v := range result {
	// 	fmt.Printf("Index %d: %+v\n", i+1, v)
	// }

	jsonString := string(jsonData)
	fmt.Println("JSON string:", jsonString)
}
