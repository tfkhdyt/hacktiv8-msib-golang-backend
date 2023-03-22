package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func main() {
	users, err := fetchUsers()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user.ID, user.Name, user.Grade)
	}

	user1, err := fetchUser("E001")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}

func fetchUsers() ([]student, error) {
	client := &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (*student, error) {
	var data student
	client := &http.Client{}

	param := url.Values{}
	param.Set("id", ID)
	payload := bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return &data, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
