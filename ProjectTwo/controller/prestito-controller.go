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

type prestito struct{}

var (
	libroServiceP   service.LibroService   = service.NewLibroService()
	clientiServiceP service.ClienteService = service.NewClienteService()
	bibliCacheSrvP  cache.BibliCache       = cache.NewRedisCache("localhost:6379", 1, 10)
)

type PrestitoController interface {
	PrestaLibro(response http.ResponseWriter, request *http.Request)
	PrestaLibroPOST(response http.ResponseWriter, request *http.Request)
	RestituisciLibro(response http.ResponseWriter, request *http.Request)
	RestituisciLibroPOST(response http.ResponseWriter, request *http.Request)
}

func NewPrestitoController() PrestitoController {
	return &prestito{}
}

func (*prestito) PrestaLibro(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	clienteID := strings.Split(request.URL.Path, "/")[2]
	cliente, _ := clientiService.Find(clienteID)
	var libro *entity.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)
	libro, _ = libroService.Find(libro.ID)
	if len(libro.Prestito) != 0 {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Il libro e' gia` prenotato."})
		return
	}
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i libri"})
		return
	}
	if cliente == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono clienti con questo ID"})
		return
	}
	if libro == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono libri con questo ID"})
		return
	}
	cliente.ListaLibri = append(cliente.ListaLibri, *libro)
	var key string = bibliCacheSrvP.Del(clienteID)
	result, err2 := clientiService.CreateWithID(cliente, key)
	libro.Prestito = "Occupato"
	var key2 string = bibliCacheSrvP.DelL(libro.ID)
	libroServiceP.CreateWithID(libro, key2)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel creare il cliente"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}

func (*prestito) PrestaLibroPOST(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var jsonString struct {
		IDLibro   string `json: "IDLibro"`
		IDCliente string `json : "IDCliente"`
	}

	var libro *entity.Libro
	var cliente *entity.Cliente

	err := json.NewDecoder(request.Body).Decode(&jsonString)
	cliente, _ = clientiService.Find(jsonString.IDCliente)
	libro, _ = libroService.Find(jsonString.IDLibro)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i dati"})
		return
	}

	if libro == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono libri con questo ID"})
		return
	}

	if cliente == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono clienti con questo ID"})
		return
	}
	cliente.ListaLibri = append(cliente.ListaLibri, *libro)
	var key string = bibliCacheSrvP.Del(jsonString.IDCliente)
	result, err2 := clientiService.CreateWithID(cliente, key)
	libro.Prestito = "Occupato"
	var key2 string = bibliCacheSrvP.Del(jsonString.IDLibro)
	libroServiceP.CreateWithID(libro, key2)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel creare il cliente"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}

func (*prestito) RestituisciLibro(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	clienteID := strings.Split(request.URL.Path, "/")[2]
	cliente, _ := clientiService.Find(clienteID)
	var libro *entity.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)
	libro, _ = libroService.Find(libro.ID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i dati"})
		return
	}
	if cliente == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono clienti con questo ID"})
		return
	}
	if libro == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono libri con questo ID"})
		return
	}

	for i := 0; i < len(cliente.ListaLibri); i++ {
		if cliente.ListaLibri[i].ID == libro.ID {
			cliente.ListaLibri[i] = cliente.ListaLibri[len(cliente.ListaLibri)-1]
			cliente.ListaLibri = cliente.ListaLibri[:len(cliente.ListaLibri)-1]
		}
		var key string = bibliCacheSrvP.Del(clienteID)
		result, err2 := clientiService.CreateWithID(cliente, key)
		libro.Prestito = ""
		var key2 string = bibliCacheSrvP.Del(libro.ID)
		libroServiceP.CreateWithID(libro, key2)
		if err2 != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel creare il cliente"})
			return
		}
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(result)

	}
}

	func (*prestito) RestituisciLibroPOST(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		var jsonString struct {
			IDLibro   string `json: "IDLibro"`
			IDCliente string `json : "IDCliente"`
		}
	
		var libro *entity.Libro
		var cliente *entity.Cliente
	
		err := json.NewDecoder(request.Body).Decode(&jsonString)
		cliente, _ = clientiService.Find(jsonString.IDCliente)
		libro, _ = libroService.Find(jsonString.IDLibro)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel prendere i dati"})
			return
		}
	
		if libro == nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono libri con questo ID"})
			return
		}
	
		if cliente == nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore non ci sono clienti con questo ID"})
			return
		}
	
		for i := 0; i < len(cliente.ListaLibri); i++ {
			if cliente.ListaLibri[i].ID == libro.ID {
				cliente.ListaLibri[i] = cliente.ListaLibri[len(cliente.ListaLibri)-1]
				cliente.ListaLibri = cliente.ListaLibri[:len(cliente.ListaLibri)-1]
			}
			var key string = bibliCacheSrvP.Del(jsonString.IDCliente)
			result, err2 := clientiService.CreateWithID(cliente, key)
			libro.Prestito = ""
			var key2 string = bibliCacheSrvP.Del(jsonString.IDLibro)
			libroServiceP.CreateWithID(libro, key2)
			if err2 != nil {
				response.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(response).Encode(errors.ServiceError{Message: "Errore nel creare il cliente"})
				return
			}
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(result)
	
		}
}
