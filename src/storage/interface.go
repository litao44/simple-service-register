package storage

import (
	"time"
)

type SSRDB interface {
	Put(key string, value string, expire time.Duration) error
	Get(key string) (string, error)
}
