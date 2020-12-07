package dao

import (
	"github.com/go-redis/redis"
)

var (
	RDb *redis.Client
)

func InitRedis()  error{
	RDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_ ,err := RDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// redis operation hmset
func SetMapForever(key string, field map[string]interface{}) (string, error) {
	return RDb.HMSet(key, field).Result()
}

// redis operation hmget
func GetMap(key string, fields ...string) ([]interface{}, error) {
	return RDb.HMGet(key, fields...).Result()
}

// redis SADD
func SetAdd(key string, field string) (int64, error) {
	return RDb.SAdd(key, field).Result()
}

// redis SISMEMBER
func SetIsMember(key string, field string) (bool, error) {
	return RDb.SIsMember(key, field).Result()
}

// redis SMEMBERS
func GetSetMembers(key string) ([]string, error) {
	return RDb.SMembers(key).Result()
}
func uniqueOccurrences(arr []int) bool {
	 num := make(map[int]int)
	 simple := make(map[int]int)
	for _ , value := range arr{
		num[value] = num[value]+1
	}

	for _ ,value := range num{
		simple[value] = simple[value]+1
		if simple[value] > 1{
			return false
		}
	}
	return true
}
