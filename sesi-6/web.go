package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT string = ":8080"

type Response struct {
	Message string `json:"message"`
}

type RequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("static/templates/*"))
}

func web() {
	http.HandleFunc("/", allAboutJson)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Name    string
			Address string
			Hobby   string
		}{
			Name:    r.URL.Query().Get("name"),
			Address: r.URL.Query().Get("address"),
			Hobby:   r.URL.Query().Get("hobby"),
		}
		// message := r.URL.Query().Get("greet")
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			fmt.Println(err)
		}
	})

	fmt.Println("Listening on port:", PORT)
	log.Fatalln(http.ListenAndServe(PORT, nil))
}

func allAboutJson(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		response := Response{
			Message: "success",
		}

		bs, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(bs))
	} else if r.Method == http.MethodPost {
		name := r.FormValue("name")
		address := r.FormValue("address")
		response := Response{
			Message: name + " " + address,
		}

		bs, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(bs))
	} else if r.Method == http.MethodPut {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		var newRequestBody RequestBody

		if err := json.Unmarshal(requestBody, &newRequestBody); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%#v\n", newRequestBody)

		bs, _ := json.Marshal(newRequestBody)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, string(bs))
	}
}
