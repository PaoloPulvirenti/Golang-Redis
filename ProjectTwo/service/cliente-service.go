package service

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/ProjectTwo/entity"
	"github.com/ProjectTwo/repository"
)

type ClienteService interface {
	Validate(cliente *entity.Cliente) error
	Create(cliente *entity.Cliente) (*entity.Cliente, error)
	CreateWithID(cliente *entity.Cliente, key string) (*entity.Cliente, error)
	Find(key string) (*entity.Cliente, error)
	FindAll() ([]entity.Cliente, error)
}

type clienteservice struct{}

var (
	repoC repository.ClienteRepository = repository.NewClienteRepository()
)

func NewClienteService() ClienteService {
	return &clienteservice{}
}

func (*clienteservice) Validate(cliente *entity.Cliente) error {
	if cliente == nil {
		err := errors.New("Il cliente e` vuoto")
		return err
	}
	if cliente.Nome == "" {
		err := errors.New("Il nome del cliente e` vuoto")
		return err
	}
	return nil
}
func (*clienteservice) Create(cliente *entity.Cliente) (*entity.Cliente, error) {
	cliente.ID = strconv.Itoa(rand.Int())
	return repoC.Save(cliente)
}

func (*clienteservice) CreateWithID(cliente *entity.Cliente, key string) (*entity.Cliente, error) {
	cliente.ID = key
	return repoC.Save(cliente)
}

func (*clienteservice) Find(key string) (*entity.Cliente, error) {
	return repoC.Find(key)
}

func (*clienteservice) FindAll() ([]entity.Cliente, error) {
	return repoC.FindAll()
}
