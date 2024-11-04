package wredis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Instance *redis.Client
}

func New(host, user, pass string, port, db, protocol int) *Client {
	instance := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: user,
		Password: pass,
		DB:       db,
		Protocol: protocol,
	})

	return &Client{Instance: instance}
}
