// Project gochat
// @author    : raisingchen
// @contact   : 20222131003@m.scnu.edu.cn
// @time      : 2025/5/16
package main

import (
	"raisingchen/gochat/cache"
	"raisingchen/gochat/conf"
	"raisingchen/gochat/model"
)

func main() {
	conf.InitConfig()
	model.InitDB()
	cache.InitRedis()
}
