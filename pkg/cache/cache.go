package cache

import "time"

type Cache interface {
	Get(key string) (value string)
	Set(key string, value interface{}, ttl time.Duration) error
}
