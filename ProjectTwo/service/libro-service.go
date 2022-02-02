package service

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/ProjectTwo/entity"
	"github.com/ProjectTwo/repository"
)


type LibroService interface{
	Validate(libro *entity.Libro) error
	Create(libro *entity.Libro) (*entity.Libro, error)
	CreateWithID(libro *entity.Libro, key string) (*entity.Libro, error)
	Find(key string) (*entity.Libro, error)
	FindAll() ([]entity.Libro, error)
}


type libroservice struct{}

var (
	repoL repository.LibroRepository=repository.NewLibroRepository()
)

func NewLibroService() LibroService{
	return &libroservice{}
}


func(*libroservice) Validate(libro *entity.Libro) error {
	if libro==nil{
		err:=errors.New("Il cliente e` vuoto")
		return err
	}
	if libro.Nome==""{
		err:=errors.New("Il nome del cliente e` vuoto")
		return err
	}
	return nil
}
func(*libroservice) Create(libro *entity.Libro) (*entity.Libro, error) {
	libro.ID=strconv.Itoa(rand.Int())
	return repoL.Save(libro)
}
func(*libroservice) Find(key string) (*entity.Libro, error) {
	return repoL.Find(key)
}

func (*libroservice) CreateWithID(libro *entity.Libro, key string) (*entity.Libro, error) {
	libro.ID = key
	return repoL.Save(libro)
}

func(*libroservice) FindAll() ([]entity.Libro, error) {
	return repoL.FindAll()
}