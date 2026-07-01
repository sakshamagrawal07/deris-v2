package core

import (
	"time"

	"github.com/sakshamsharma/deris-v2/config"
)

var store map[string]*RedisObj
var expires map[*RedisObj]uint64

func init() {
	store = make(map[string]*RedisObj)
	expires = make(map[*RedisObj]uint64)
}

func setExpiry(obj *RedisObj, expDurationMs int64) {
	expires[obj] = uint64(time.Now().UnixMilli()) + uint64(expDurationMs)
}

func NewRedisObj(value interface{}, expDurationMs int64, oType uint8, oEnc uint8) *RedisObj {
	obj := &RedisObj{
		Value: value,
		TypeEncoding: oType | oEnc,
		LastAccessedAt: getCurrentClock(),
	}
	if expDurationMs > 0 {
		setExpiry(obj, expDurationMs)
	}
	return obj
}

func Put(key string, value *RedisObj) {
	if len(store) >= config.KeysLimit {
		evict()
	}
	value.LastAccessedAt = getCurrentClock()
	store[key] = value
	if KeyspaceStat[0] == nil {
		KeyspaceStat[0] = make(map[string]int)
	}
	KeyspaceStat[0]["keys"]++
}

func Get(key string) *RedisObj {
	v := store[key]
	if v != nil {
		if hasExpired(v) {
			Del(key)
			return nil
		}
	}
	v.LastAccessedAt = getCurrentClock()
	return v
}

func Del(key string) bool {
	if obj, ok := store[key]; ok {
		delete(store, key)
		delete(expires, obj)
		KeyspaceStat[0]["keys"]--
		return true
	}
	return false
}
