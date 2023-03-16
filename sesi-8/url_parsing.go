package main

import (
	"fmt"
	"log"
	"net/url"
)

func urlParsing() {
	urlString := "https://tfkhdyt.my.id:3000/api/messages?text=Hello"
	url, err := url.Parse(urlString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("URL       : %s\n", urlString)
	fmt.Printf("Protocol  : %s\n", url.Scheme)
	fmt.Printf("Host      : %s\n", url.Host)
	fmt.Printf("Path      : %s\n", url.Path)

	name := url.Query()["text"][0]
	fmt.Printf("Text      : %s\n", name)
}
