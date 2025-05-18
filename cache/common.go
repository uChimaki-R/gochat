// Package cache
// @author    : raisingchen
// @contact   : 20222131003@m.scnu.edu.cn
// @time      : 2025/5/16
package cache

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"raisingchen/gochat/conf"
	"strconv"
)

var RedisClient *redis.Client

func InitRedis() {
	redisConfig := conf.Config.Redis
	db, _ := strconv.ParseUint(redisConfig.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.RedisAddr,
		DB:       int(db),
		Password: redisConfig.RedisPassword,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	RedisClient = client
	log.Info("Redis connect success: " + redisConfig.RedisAddr)
}
