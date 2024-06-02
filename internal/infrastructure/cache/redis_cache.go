package cache

import (
	"context"
	"demo/internal/infrastructure/config"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	MyCachePool *RedisCache
)

type RedisCache struct {
	Client        *redis.Client
	Ctx           context.Context
	DefaultExpire time.Duration
}

func NewRedisCache(config *config.Config) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.CacheConfig.Host + ":" + config.CacheConfig.Port,
		Password:     config.CacheConfig.Password,                  // no password set
		DB:           config.CacheConfig.Db,                        // use default DB
		PoolSize:     config.CacheConfig.PoolSize,                  // default pool size
		MinIdleConns: config.CacheConfig.MinIdleCon,                // minimum number of idle connections
		PoolTimeout:  config.CacheConfig.PoolTimeout * time.Second, // pool timeout
		IdleTimeout:  config.CacheConfig.IdleTimeout * time.Minute, // idle timeout
	})

	return &RedisCache{
		Client:        rdb,
		Ctx:           context.Background(),
		DefaultExpire: 60 * time.Second,
	}
}

func (r *RedisCache) SetWithExpire(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r *RedisCache) Set(key string, value interface{}) error {
	expiration := r.DefaultExpire
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisCache) Delete(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *RedisCache) HGet(table string, key string) (string, error) {
	return r.Client.HGet(r.Ctx, table, key).Result()
}

func InitRedisCache(config *config.Config) {
	MyCachePool = NewRedisCache(config)
}
