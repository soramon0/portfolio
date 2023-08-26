package cache

import "github.com/redis/go-redis/v9"

type Cache struct {
	Client *redis.Client
}

func NewCache(url string) (*Cache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opts)

	return &Cache{Client: rdb}, nil
}
