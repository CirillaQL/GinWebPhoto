package util

import (
	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func InitRedisPool() {
	Pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
