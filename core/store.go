package core

import "time"

var store map[string]*RedisObj

type RedisObj struct {
	Value     interface{}
	ExpiresAt int64
}

func init() {
	store = make(map[string]*RedisObj)
}

func NewRedisObj(value interface{}, durationMs int64) *RedisObj {
	var expiresAt int64 = -1
	if durationMs > 0 {
		expiresAt = time.Now().UnixMilli() + durationMs
	}
	return &RedisObj{
		Value:     value,
		ExpiresAt: expiresAt,
	}
}

func Put(key string, value *RedisObj) {
	store[key] = value
}

func Get(key string) *RedisObj {
	return store[key]
}
