package entity

type Cliente struct {
	ID         string  `json: "ID"`
	Nome       string  `json : "Nome"`
	Cognome    string  `json : "Cognome"`
	Email      string  `json: "email"`
	ListaLibri []Libro `json: "listaLibri"`
}
