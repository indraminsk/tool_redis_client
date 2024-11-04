package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func New(host, user, pass string, port, db, protocol int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: user,
		Password: pass,
		DB:       db,
		Protocol: protocol,
	})
}
