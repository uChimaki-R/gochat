// Package model
// @author    : raisingchen
// @contact   : 20222131003@m.scnu.edu.cn
// @time      : 2025/5/17
package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"raisingchen/gochat/conf"
	"strings"
	"time"
)

var DB *gorm.DB
var MongoClient *mongo.Client

func InitDB() {
	loadMysql(&conf.Config.Mysql)
	loadMongoDb(&conf.Config.MongoDB)
}

func loadMysql(mysqlConfig *conf.MysqlConf) {
	connString := strings.Join([]string{mysqlConfig.DbUser, ":", mysqlConfig.DbPassword, "@tcp(", mysqlConfig.DbAddr, ")/", mysqlConfig.DbName, "?charset=utf8&parseTime=true&loc=Local"}, "")
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
	log.Info("Mysql connect success: " + connString)
}

func loadMongoDb(mongoConfig *conf.MongoConf) {
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoConfig.MongoDbAddr)
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("MongoDb connect success: " + mongoConfig.MongoDbAddr)
}
