package wredis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Instance *redis.Client
	Host     string
	User     string
	Pass     string
	Port     int
	DB       int
	Protocol int
}

func New(host, user, pass string, port, db, protocol int) *Client {
	instance := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: user,
		Password: pass,
		DB:       db,
		Protocol: protocol,
	})

	return &Client{
		Instance: instance,
		Host:     host,
		User:     user,
		Pass:     pass,
		Port:     port,
		DB:       db,
		Protocol: protocol,
	}
}
