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

type clientecontroller struct{}

var (
	clientiService service.ClienteService = service.NewClienteService()
	bibliCacheSrvC cache.BibliCache       = cache.NewRedisCache("localhost:6379", 1, 10)
)

type ClienteController interface {
	GetClienteByID(response http.ResponseWriter, request *http.Request)
	AddCliente(response http.ResponseWriter, request *http.Request)
	GetClienti(response http.ResponseWriter, request *http.Request)
}

func NewClienteController() ClienteController {
	return &clientecontroller{}
}

func (*clientecontroller) GetClienti(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	cliente, err := clientiService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i clienti"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(cliente)
}

func (*clientecontroller) GetClienteByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	clienteID := strings.Split(request.URL.Path, "/")[2]
	cliente, err := clientiService.Find(clienteID)
	if cliente == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Non ho trovato clienti con questo ID"})
		return
	}
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i clienti"})
		return
	}
	if !(len(cliente.Cognome) != 0) {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Per la GET di questo id devi richiamare ../libroID/{id}"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(cliente)
}

func (*clientecontroller) AddCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var cliente entity.Cliente
	err := json.NewDecoder(request.Body).Decode(&cliente)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i clienti"})
		return
	}
	err1 := clientiService.Validate(&cliente)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := clientiService.Create(&cliente)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel salvataggio del cliente"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
