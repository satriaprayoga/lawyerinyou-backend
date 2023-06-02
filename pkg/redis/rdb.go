package redis

import (
	"fmt"
	"lawyerinyou-backend/pkg/settings"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Setup() {
	now := time.Now()
	conString := fmt.Sprintf("%s:%d", settings.AppConfigSetting.RedisDB.Host, settings.AppConfigSetting.RedisDB.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     conString,
		Password: settings.AppConfigSetting.RedisDB.Password,
		DB:       settings.AppConfigSetting.RedisDB.DB,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config redis is ready in %v", timeSpent)
}
