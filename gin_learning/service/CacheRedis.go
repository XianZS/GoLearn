package service

import (
	"GoLearn/gin_learning/database"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func CacheOneUser() {
	conn := database.RedisDefaultPool.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	data, err := redis.Bytes(conn.Do("GET", "111"))
	if err != nil {
		fmt.Println(data)
	}
}
