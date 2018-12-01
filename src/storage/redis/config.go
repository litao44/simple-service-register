package redis

import "github.com/go-redis/redis"

type Config struct {
	Addr     string
	Password string
	DB       int
}

func (cfg *Config) Options() *redis.Options {
	op := redis.Options{}
	if cfg.Addr != "" {
		op.Addr = cfg.Addr
	}
	if cfg.Password != "" {
		op.Password = cfg.Password
	}
	return &op
}

func (cfg *Config) NewClient() *RedisSSRDBImpl {
	return NewRedisSSRDBImplWithOption(cfg.Options())
}
