package Myserver

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leiyuuu/second-test/naive-server/src/Mydatabase"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

func response_Signin() {
	server.POST("/signin", func(context *gin.Context) {
		loggenerator.Trace("get a signin")
		b, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		access_token, err := Mydatabase.Sign_in(m["username"], m["password"])
		if err == 0 { //没有错误
			context.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "",
				"data": gin.H{
					"access_token": access_token,
				},
			})
			loggenerator.Trace("signin return succeed")
		} else if err == 1 { //用户不存在
			context.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "The user does not exist",
				"data": nil,
			})
			loggenerator.Trace("signin error: user not existed")
		} else if err == 2 { //密码错误
			context.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "The password is wrong",
				"data": nil,
			})
			loggenerator.Trace("signin error: password wrong")
		} else if err == 3 { //已经登录
			context.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "The user has signed in",
				"data": nil,
			})
			loggenerator.Trace("signin error: the user has signed in")
		} else { //未知错误
			context.JSON(http.StatusOK, gin.H{
				"code": 999,
				"msg":  "Unknown Exception",
				"data": nil,
			})
			loggenerator.Error("SIGN IN UNKNOWN ERROR!!!!")
		}
	})
}
