package cache

import (
	"encoding/json"
	"time"

	"github.com/ProjectTwo/entity"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) BibliCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
func (cache *redisCache) Set(key string, Cliente *entity.Cliente) (*entity.Cliente, error) {

	client := cache.GetClient()

	json, err := json.Marshal(Cliente)
	if err != nil {
		panic(err)
	}
	client.HSet("Clienti", key, json)
	//client.Set(key, json, 100*cache.expires*time.Second)

	return Cliente, nil
}

func (cache *redisCache) Get(key string) *entity.Cliente {
	client := cache.GetClient()

	//val, err := client.Get(key).Result()
	val, err := client.HGet("Clienti", key).Result()

	if err != nil {
		return nil
	}

	cliente := entity.Cliente{}

	err = json.Unmarshal([]byte(val), &cliente)

	if err != nil {
		panic(err)
	}

	return &cliente
}

func (cache *redisCache) GetAll() []entity.Cliente {
	client := cache.GetClient()

	val, err := client.HGetAll("Clienti").Result()
	if err != nil {
		return nil
	}
	cliente := entity.Cliente{}
	var clienteArr []entity.Cliente
	for k, v := range val {
		println(k)
		err := json.Unmarshal([]byte(v), &cliente)
		clienteArr = append(clienteArr, cliente)
		if err != nil {
			panic(err)
		}
	}
	return clienteArr
}

func (cache *redisCache) Del(key string) string {
	client := cache.GetClient()
	//val, err := client.Del(key).Result()
	val, err := client.HDel("Clienti", key).Result()
	if err != nil {
		return "errore"
	}
	if val < 0 {
		return "errore valore"
	}

	return key
}

func (cache *redisCache) DelL(key string) string {
	client := cache.GetClient()
	//val, err := client.Del(key).Result()
	val, err := client.HDel("Libro", key).Result()
	if err != nil {
		return "errore"
	}
	if val < 0 {
		return "errore valore"
	}

	return key
}

func (cache *redisCache) SetL(key string, Libro *entity.Libro) (*entity.Libro, error) {

	client := cache.GetClient()

	json, err := json.Marshal(Libro)
	if err != nil {
		panic(err)
	}
	client.HSet("Libro", key, json)
	//client.Set(key, json, 100*cache.expires*time.Second)
	return Libro, nil
}

func (cache *redisCache) GetL(key string) *entity.Libro {
	client := cache.GetClient()

	//val, err := client.Get(key).Result()
	val, err := client.HGet("Libro", key).Result()

	if err != nil {
		return nil
	}

	libro := entity.Libro{}
	err = json.Unmarshal([]byte(val), &libro)

	if err != nil {
		panic(err)
	}

	return &libro
}

func (cache *redisCache) GetAllL() []entity.Libro {
	client := cache.GetClient()

	val, err := client.HGetAll("Libro").Result()
	if err != nil {
		return nil
	}
	libro := entity.Libro{}
	var libroArr []entity.Libro
	for k, v := range val {
		println(k)
		err := json.Unmarshal([]byte(v), &libro)
		libroArr = append(libroArr, libro)
		if err != nil {
			panic(err)
		}
	}
	return libroArr
}
