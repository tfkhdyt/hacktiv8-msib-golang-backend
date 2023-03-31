package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type StatusDetail struct {
	Value       int    `json:"value"`
	Description string `json:"description"`
}

type Response struct {
	Water StatusDetail `json:"water"`
	Wind  StatusDetail `json:"wind"`
}

func generateStatus(min int, max int) {
	for {
		status := Status{
			Water: rand.Intn(max-min) + min,
			Wind:  rand.Intn(max-min) + min,
		}

		b, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			log.Fatalln("Failed to marshal status to json.", err.Error())
		}

		if err := os.WriteFile("data.json", b, 0644); err != nil {
			log.Fatalln("Failed to write json to file.", err.Error())
		}

		time.Sleep(15 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go generateStatus(1, 100)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		b, err := os.ReadFile("data.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var status Status
		if err := json.Unmarshal(b, &status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response := Response{}
		response.Water.Value = status.Water
		response.Wind.Value = status.Wind

		if status.Water <= 5 {
			response.Water.Description = "Aman"
		} else if status.Water >= 6 && status.Water <= 8 {
			response.Water.Description = "Siaga"
		} else if status.Water > 8 {
			response.Water.Description = "Bahaya"
		}

		if status.Wind <= 6 {
			response.Wind.Description = "Aman"
		} else if status.Wind >= 7 && status.Wind <= 15 {
			response.Wind.Description = "Siaga"
		} else if status.Wind > 15 {
			response.Wind.Description = "Bahaya"
		}

		if err := tmpl.Execute(w, response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server listening to port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Failed to start server.", err.Error())
	}
}
