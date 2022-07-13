package Config

import (
	"flag"
	"github.com/gomodule/redigo/redis"
	"time"
)

var Pool *redis.Pool

var RedisServer = flag.String("redisServer", ":6379", "")

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
