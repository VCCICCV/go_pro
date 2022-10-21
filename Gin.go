package main

import "github.com/gin-gonic/gin"

// PROJECT_NAME:Test
// DATE:2022/10/4 1:01
// USER:21005
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
