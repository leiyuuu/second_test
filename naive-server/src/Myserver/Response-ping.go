package Myserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

func response_Ping() {
	server.GET("/ping", func(context *gin.Context) {
		loggenerator.Trace("get a ping")
		// fmt.Println("GET A PING!!!")
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": gin.H{
				"msg": "pong!",
			},
		})
		loggenerator.Trace("ping return succeed")
	})
}
