package adapters

import (
	"context"
	"demo_time_sheet_server/common/configs"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

// NewRedis initializes and returns a new Redis client instance.
func NewRedis() *Redis {
	config := configs.Get()
	return &Redis{
		client: redis.NewClient(&redis.Options{
			Addr:         config.AddressRedis,
			Password:     config.PasswordRedis,
			DB:           config.DatabaseRedisIndex,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
		}),
	}
}

// Connect attempts to connect to the Redis server and verifies the connection.
func (r *Redis) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.client.Ping(ctx).Err(); err != nil {
		return err
	}
	log.Println("connect redis success")
	return nil
}

// Client returns the underlying Redis client for direct usage.
func (r *Redis) Client() *redis.Client {
	return r.client
}
