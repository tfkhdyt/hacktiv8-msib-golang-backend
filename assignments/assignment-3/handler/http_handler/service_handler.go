package http_handler

import (
	"assignment_3/service"
	"html/template"
	"net/http"
)

type serviceHandler struct {
	serviceService service.StatusService
}

func NewServiceHandler(serviceService service.StatusService) *serviceHandler {
	return &serviceHandler{serviceService: serviceService}
}

func (s *serviceHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := s.serviceService.ReadStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, "index", response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
