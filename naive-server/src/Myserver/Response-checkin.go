package Myserver

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leiyuuu/second-test/naive-server/src/Mydatabase"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

func response_Checkin() {
	server.POST("/checkin", func(context *gin.Context) {
		loggenerator.Trace("get a checkin")
		b, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		err := Mydatabase.Check_in(m["access_token"])
		if err == 0 { //成功签到
			context.JSON(200, gin.H{
				"point": rand.Int31n(100) + 1, //随机数生成分数，保底1分hhh
			})
			loggenerator.Trace("checkin return succeed")
		} else if err == 1 { //token没有或者无效
			context.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "invalid access_token",
				"data": nil,
			})
			loggenerator.Trace("checkin error:invalid token")
		} else if err == 2 { //已经签到过了
			context.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "The user has already checked in",
				"data": nil,
			})
			loggenerator.Trace("checkin error:already checked in")
		} else if err == 3 { // 没有登录
			context.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "The user hasn't signed in",
				"data": nil,
			})
			loggenerator.Trace("checkin error:not signed in")
		} else { // 未知错误
			context.JSON(http.StatusOK, gin.H{
				"code": 999,
				"msg":  "Unknown Exception",
				"data": nil,
			})
			loggenerator.Error("CHECK IN UNKNOWN ERROR!!!!")
		}
	})
}
