package repository

import (
	"log"

	"github.com/ProjectTwo/cache"
	"github.com/ProjectTwo/entity"
)

type LibroRepository interface {
	Save(libro *entity.Libro) (*entity.Libro, error)
	Find(key string) (*entity.Libro, error)
	FindAll() ([]entity.Libro, error)
}

type repoL struct{}

var (
	bibliCacheSrvL cache.BibliCache = cache.NewRedisCache("localhost:6379", 1, 10)
)

func NewLibroRepository() LibroRepository {
	return &repoL{}
}

func (*repoL) Save(libro *entity.Libro) (*entity.Libro, error) {

	libro, err := bibliCacheSrvL.SetL(libro.ID, libro)
	if err != nil {
		log.Fatalf("Cliente non aggiunto : %v", err)
		return nil, err
	}
	return libro, nil

}

func (*repoL) Find(key string) (*entity.Libro, error) {

	libri := bibliCacheSrvL.GetL(key)
	return libri, nil

}

func (*repoL) FindAll() ([]entity.Libro, error) {
	libri := bibliCacheSrvL.GetAllL()
	return libri, nil
}
