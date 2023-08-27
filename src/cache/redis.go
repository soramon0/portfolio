package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/soramon0/portfolio/src/lib"
	"golang.org/x/net/context"
)

type Cache struct {
	Client *redis.Client
	log    *lib.AppLogger
}

func NewCache(url string, l *lib.AppLogger) (*Cache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opts)

	return &Cache{Client: rdb, log: l}, nil
}

func (c *Cache) CounterRateLimit(ctx context.Context, key string, limit int, perSec int) (int, bool) {
	latestCount, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// set counter with max requests (limit) per second if we do not have one
		// save in redis
		err = c.Client.SetEx(ctx, key, limit, time.Second*time.Duration(perSec)).Err()
		if err != nil {
			c.log.ErrorF("failed to set rate limit for key %s: %v", key, err)
			return 0, false
		}
	}

	// break early if counter hits 0
	if latestCount == "0" {
		return 0, false
	}

	requests, err := c.Client.Decr(ctx, key).Result()
	if err != nil {
		c.log.ErrorF("failed to decrement rate limit for key %s: %v", key, err)
		return 0, false
	}

	if requests <= 0 {
		return int(requests), false
	}

	return int(requests), true
}
