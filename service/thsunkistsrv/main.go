package main

import (
	"hi_golang/service/thsunkistsrv/contrib/logger"
	userV1 "hi_golang/service/thsunkistsrv/controller/v1/user"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.Debug("thsunkistsrv Running")
	logger.Info("thsunkistsrv Running")
	logger.Warn("thsunkistsrv Running")
	logger.Error("thsunkistsrv Running")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/v1/user/info", userV1.Info)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
