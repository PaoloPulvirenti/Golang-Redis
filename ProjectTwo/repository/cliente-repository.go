package repository

import (
	"log"

	"github.com/ProjectTwo/cache"
	"github.com/ProjectTwo/entity"
)

type ClienteRepository interface {
	Save(cliente *entity.Cliente) (*entity.Cliente, error)
	Find(key string) (*entity.Cliente, error)
	FindAll() ([]entity.Cliente, error)
}

type repoC struct{}

var (
	bibliCacheSrv cache.BibliCache = cache.NewRedisCache("localhost:6379", 1, 10)
)

func NewClienteRepository() ClienteRepository {
	return &repoC{}
}

func (*repoC) Save(cliente *entity.Cliente) (*entity.Cliente, error) {

	cliente, err := bibliCacheSrv.Set(cliente.ID, cliente)
	if err != nil {
		log.Fatalf("Cliente non aggiunto : %v", err)
		return nil, err
	}
	return cliente, nil

}

func (*repoC) Find(key string) (*entity.Cliente, error) {

	clienti := bibliCacheSrv.Get(key)
	return clienti, nil

}

func (*repoC) FindAll() ([]entity.Cliente, error) {
	clienti := bibliCacheSrv.GetAll()
	return clienti, nil
}
