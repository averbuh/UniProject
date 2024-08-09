package recipes

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redisClient *redis.Client
}

func NewRedis(Addr string, Password string, DB int) (*Redis, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password, // no password set
		DB:       DB,       // use default DB
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}
	return &Redis{
		redisClient: redisClient,
	}, err
}

func (r *Redis) Add(name string, recipe Recipe) error {
	err := r.redisClient.Set(context.Background(), name, recipe, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) AddImageURL(name string, image string) error {
	err := r.redisClient.Set(context.Background(), name, image, 5*time.Minute).Err()
	return err
}

func (r *Redis) GetImageURL(name string) (Image, error) {
	url, err := r.redisClient.Get(context.Background(), name).Result()
	if err != nil {
		return Image{}, err
	}
	return Image{URL: url}, nil
}

func (r *Redis) Get(name string) (Recipe, error) {
	var recipe Recipe
	err := r.redisClient.Get(context.Background(), name).Scan(&recipe)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

// func (r *Redis) List() (map[string]Recipe, error) {

// 	return nil, nil
// }

// func (r *Redis) Update(name string, recipe Recipe) error {

// 	return nil
// }

// func (r *Redis) Remove(name string) error {

// 	return nil
// }
