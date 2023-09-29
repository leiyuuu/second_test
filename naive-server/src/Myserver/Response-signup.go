package Myserver

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leiyuuu/second-test/naive-server/src/Mydatabase"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

func response_Signup() {
	server.POST("/signup", func(context *gin.Context) {
		loggenerator.Trace("get a signup")
		b, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		access_token, err := Mydatabase.Sign_up(m["username"], m["password"])
		if err == 0 { //没有错误
			context.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "",
				"data": gin.H{
					"access_token": access_token,
				},
			})
			loggenerator.Trace("signup return succeed")
		} else if err == 1 { //用户名重复
			context.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Duplicated user name",
				"data": nil,
			})
			loggenerator.Trace("checkin error:Duplicated user name")
		} else if err == 2 { //密码为空
			context.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "empty password",
				"data": nil,
			})
			loggenerator.Trace("checkin error:empty password")
		} else if err == 3 { //用户名为空
			context.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "empty username",
				"data": nil,
			})
			loggenerator.Trace("checkin error:empty username")
		} else { //未知错误
			context.JSON(http.StatusOK, gin.H{
				"code": 999,
				"msg":  "Unknown Exception",
				"data": nil,
			})
			loggenerator.Error("SIGN UP UNKNOWN ERROR!!!!")
		}
	})
}
