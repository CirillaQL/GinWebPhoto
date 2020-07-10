package util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T){
	var conn = Pool.Get()
	defer conn.Close()

	conn.Do("set", "cat1", "tom")
	line, _ := redis.String(conn.Do("get", "cat1"))

	fmt.Println(line)
}
