package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisSSRDBImpl struct {
	cli *redis.Client
}

func NewRedisSSRDBImpl(cfg *Config) *RedisSSRDBImpl {
	return NewRedisSSRDBImplWithOption(cfg.Options())
}

func NewRedisSSRDBImplWithOption(op *redis.Options) *RedisSSRDBImpl {
	cli := RedisSSRDBImpl{
		cli: redis.NewClient(op),
	}
	return &cli
}

func (impl *RedisSSRDBImpl) Put(key, value string, expire time.Duration) error {
	return impl.cli.Set(key, value, expire).Err()
}

func (impl *RedisSSRDBImpl) Get(key string) (string, error) {
	return impl.cli.Get(key).Result()
}
