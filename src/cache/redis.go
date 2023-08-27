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

func (c *Cache) CounterRateLimit(ctx context.Context, key string, limit int, perSec int, backoffDuration int) (bool, int) {
	latestCount, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// set counter with max requests (limit) per second if we do not have one
		// save in redis
		err = c.Client.SetEx(ctx, key, limit, time.Second*time.Duration(perSec)).Err()
		if err != nil {
			c.log.ErrorF("failed to set rate limit for key %s: %v", key, err)
			return false, backoffDuration
		}
	}

	// break early if counter hits 0
	// or 1 since we will decrement it in this round from 0 to 1
	if latestCount == "0" || latestCount == "1" {
		// stop user from sending any further request for backoff duration amount
		if err = c.Client.SetEx(ctx, key, -1, time.Second*time.Duration(backoffDuration)).Err(); err != nil {
			return false, backoffDuration
		}
		return false, backoffDuration
	}

	// break early while back off duration has not expired yet
	if latestCount == "-1" {
		durationLeft, err := c.Client.TTL(ctx, key).Result()
		if err != nil {
			return false, backoffDuration
		}
		return false, int(durationLeft / time.Second)
	}

	requests, err := c.Client.Decr(ctx, key).Result()
	if err != nil {
		c.log.ErrorF("failed to decrement rate limit for key %s: %v", key, err)
		return false, backoffDuration
	}

	if requests <= 0 {
		return false, backoffDuration
	}

	return true, backoffDuration
}
