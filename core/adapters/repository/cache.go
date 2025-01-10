package repository

import (
	"context"
	"rices/core/adapters"
	"rices/core/adapters/cache"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	db *adapters.Redis
}

func NewRepositoryCache(db *adapters.Redis) cache.CacheOperations {
	return &cacheRepository{
		db: db,
	}
}

// Delete implements cache.CacheOperations.
func (c *cacheRepository) Delete(ctx context.Context, key string) error {
	// Use Redis DEL command to delete the key.
	err := c.db.Client().Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

// Exists implements cache.CacheOperations.
func (c *cacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	// Use Redis EXISTS command to check if the key exists.
	result, err := c.db.Client().Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Expire implements cache.CacheOperations.
func (c *cacheRepository) Expire(ctx context.Context, key string, expiration time.Duration) error {
	// Use Redis EXPIRE command to set the expiration for the key.
	err := c.db.Client().Expire(ctx, key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get implements cache.CacheOperations.
func (c *cacheRepository) Get(ctx context.Context, key string, dest interface{}) error {
	// Use Redis GET command to retrieve the value for the given key.
	val, err := c.db.Client().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// If the key doesn't exist, return a nil value.
			return nil
		}
		return err
	}
	// Assuming dest is a pointer to a string for simplicity (you can modify based on your need).
	*dest.(*string) = val
	return nil
}

// HGet implements cache.CacheOperations.
func (c *cacheRepository) HGet(ctx context.Context, key string, field string) (string, error) {
	// Use Redis HGET command to get a field from a hash.
	val, err := c.db.Client().HGet(ctx, key, field).Result()
	if err != nil {
		if err == redis.Nil {
			// If the field doesn't exist, return an empty string.
			return "", nil
		}
		return "", err
	}
	return val, nil
}

// HGetAll implements cache.CacheOperations.
func (c *cacheRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	// Use Redis HGETALL command to retrieve all fields and values from a hash.
	result, err := c.db.Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HSet implements cache.CacheOperations.
func (c *cacheRepository) HSet(ctx context.Context, key string, values ...interface{}) error {
	// Convert the values to a map.
	// Assuming the values are in key-value pairs and each key and value are strings.
	hashValues := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		hashValues[values[i].(string)] = values[i+1].(string)
	}

	// Use Redis HSET command to set fields in the hash.
	err := c.db.Client().HSet(ctx, key, hashValues).Err()
	if err != nil {
		return err
	}
	return nil
}

// Set implements cache.CacheOperations.
func (c *cacheRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// Use Redis SET command to set the key-value pair with expiration time.
	err := c.db.Client().Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
