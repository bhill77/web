package cache

import "time"

type Repository interface {
	Set(key string, value string, expiration time.Duration) error
}
