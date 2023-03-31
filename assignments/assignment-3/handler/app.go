package handler

import (
	"assignment_3/handler/http_handler"
	"assignment_3/service"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func StartApp() {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	statusService := service.NewStatusService(r)
	statusHandler := http_handler.NewServiceHandler(statusService)

	go statusService.GenerateStatus(1, 100)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", statusHandler.Index)

	log.Println("Server listening to port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Failed to start server.", err.Error())
	}
}
