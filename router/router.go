// Package router
// @author    : raisingchen
// @contact   : 20222131003@m.scnu.edu.cn
// @time      : 2025/5/19
package router

import (
	"github.com/gin-gonic/gin"
	"raisingchen/gochat/api"
)

func NewRouter() (r *gin.Engine) {
	r = gin.Default()
	{
		group := r.Group("/")
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		group.POST("/user/register", api.UserRegister)
	}
	return
}
