package Myserver

import (
	"github.com/gin-gonic/gin"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

var server = gin.Default()

func gen_a_server() {
	server.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func Genserver() { //设置API，开启服务器
	loggenerator.Info("begin genserver")
	gin.SetMode(gin.ReleaseMode)
	response_Ping()
	response_Checkin()
	response_Signup()
	response_Signin()
	response_Clear()
	loggenerator.Info("server init succeed")
	gen_a_server() //先设置各种服务，最后再gen！

}
