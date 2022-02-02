package cache

import "github.com/ProjectTwo/entity"

type BibliCache interface {
	Set(key string, cliente *entity.Cliente) (*entity.Cliente, error)
	Get(key string) *entity.Cliente
	GetAll() []entity.Cliente
	SetL(key string, libro *entity.Libro) (*entity.Libro, error)
	GetL(key string) *entity.Libro
	GetAllL() []entity.Libro
	Del(key string) string
	DelL(key string) string
}
