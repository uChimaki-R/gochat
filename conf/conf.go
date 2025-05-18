// Package conf
// @author    : raisingchen
// @time      : 2025/5/16
package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

type Conf struct {
	App     AppConf
	Mysql   MysqlConf
	Redis   RedisConf
	MongoDB MongoConf
}

type AppConf struct {
	AppMode  string
	HttpPort string
}

type MysqlConf struct {
	DbAddr     string
	DbUser     string
	DbPassword string
	DbName     string
}

type RedisConf struct {
	RedisDb       string
	RedisAddr     string
	RedisDbName   string
	RedisPassword string
}

type MongoConf struct {
	MongoDbName string
	MongoDbAddr string
}

var Config Conf

func InitConfig() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	loadAppConfig(file, &Config.App)
	loadMySqlConfig(file, &Config.Mysql)
	loadMongoDbConfig(file, &Config.MongoDB)
	loadRedisConfig(file, &Config.Redis)
}

func loadRedisConfig(file *ini.File, redisConfig *RedisConf) {
	redisConfig.RedisAddr = file.Section("redis").Key("RedisAddr").String()
	redisConfig.RedisDbName = file.Section("redis").Key("RedisDbName").String()
	redisConfig.RedisDb = file.Section("redis").Key("RedisDb").String()
	redisConfig.RedisPassword = file.Section("redis").Key("RedisPassword").String()
}

func loadMongoDbConfig(file *ini.File, mongoConfig *MongoConf) {
	mongoConfig.MongoDbName = file.Section("mongodb").Key("MongoDbName").String()
	mongoConfig.MongoDbAddr = file.Section("mongodb").Key("MongoDbAddr").String()
}

func loadMySqlConfig(file *ini.File, mysqlConfig *MysqlConf) {
	mysqlConfig.DbUser = file.Section("mysql").Key("DbUser").String()
	mysqlConfig.DbAddr = file.Section("mysql").Key("DbAddr").String()
	mysqlConfig.DbPassword = file.Section("mysql").Key("DbPassword").String()
	mysqlConfig.DbName = file.Section("mysql").Key("DbName").String()
}

func loadAppConfig(file *ini.File, appConfig *AppConf) {
	appConfig.AppMode = file.Section("service").Key("AppMode").String()
	appConfig.HttpPort = file.Section("service").Key("HttpPort").String()
}
