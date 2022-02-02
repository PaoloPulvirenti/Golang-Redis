package entity

type Libro struct {
	ID       string `json: "ID"`
	Nome     string `json : "Nome"`
	Scaffale string `json : "Scaffale"`
	Prestito string `json : "Prestito"`
}
