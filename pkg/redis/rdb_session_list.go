package redis

import (
	"fmt"
	"time"
)

func GetList(key string) ([]string, error) {
	list, err := rdb.SMembers(key).Result()
	return list, err
}

func RemoveList(key string, val interface{}) error {
	_, err := rdb.SRem(key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

func AddList(key, val string) error {
	_, err := rdb.SAdd(key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

func TurncateList(key string) error {
	_, err := rdb.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

func AddSession(key string, val interface{}, mn time.Duration) error {
	set, err := rdb.Set(key, val, mn).Result()
	if err != nil {
		return err
	}
	fmt.Println(set)
	return nil
}

func GetSession(key string) interface{} {
	value := rdb.Get(key).Val()
	fmt.Println(value)
	return value
}
