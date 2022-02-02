package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ProjectTwo/cache"
	"github.com/ProjectTwo/entity"
	"github.com/ProjectTwo/errors"
	"github.com/ProjectTwo/service"
)

type librocontroller struct{}

var (
	libroService   service.LibroService = service.NewLibroService()
	bibliCacheSrvL cache.BibliCache     = cache.NewRedisCache("localhost:6379", 1, 10)
)

type LibroController interface {
	GetLibri(response http.ResponseWriter, request *http.Request)
	AddLibro(response http.ResponseWriter, request *http.Request)
	GetLibroByID(response http.ResponseWriter, request *http.Request)
}

func NewLibroController() LibroController {
	return &librocontroller{}
}

func (*librocontroller) AddLibro(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var libro entity.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i libri"})
		return
	}
	err1 := libroService.Validate(&libro)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := libroService.Create(&libro)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel salvataggio del libro"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*librocontroller) GetLibroByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	libroID := strings.Split(request.URL.Path, "/")[2]
	libro, err := libroService.Find(libroID)
	if libro == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Non esistono libri con questo ID"})
		return
	}
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i libri"})
		return
	}
	if !(len(libro.Scaffale) != 0) {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Per la GET di questo id devi richiamare ../clienteID/{id}"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(libro)
}

func (*librocontroller) GetLibri(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	libro, err := libroService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i clienti"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(libro)
}
