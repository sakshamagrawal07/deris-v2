package core

import (
	"time"

	"github.com/sakshamsharma/deris-v2/config"
)

var store map[string]*RedisObj

func init() {
	store = make(map[string]*RedisObj)
}

func NewRedisObj(value interface{}, durationMs int64, oType uint8, oEnc uint8) *RedisObj {
	var expiresAt int64 = -1
	if durationMs > 0 {
		expiresAt = time.Now().UnixMilli() + durationMs
	}
	return &RedisObj{
		TypeEncoding: oType | oEnc,
		Value:     value,
		ExpiresAt: expiresAt,
	}
}

func Put(key string, value *RedisObj) {
	if len(store) >= config.KeysLimit {
		evict()
	}
	store[key] = value
	if KeyspaceStat[0] == nil {
		KeyspaceStat[0] = make(map[string]int)
	}
	KeyspaceStat[0]["keys"]++
}

func Get(key string) *RedisObj {
	v := store[key]
	if v != nil {
		if v.ExpiresAt <= time.Now().UnixMilli() {
			Del(key)
			return nil
		}
	}
	return v
}

func Del(key string) bool {
	if _, ok := store[key]; ok {
		delete(store, key)
		KeyspaceStat[0]["keys"]--
		return true
	}
	return false
}
