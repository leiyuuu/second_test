package Myserver

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leiyuuu/second-test/naive-server/src/Mydatabase"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

func response_Clear() { //[POST] /clear操作 软删除数据库中所以数据，用于调试等操作
	//需要在json文件中提供admin_key 为 114514 才能操作成功
	server.POST("/clear", func(context *gin.Context) {
		loggenerator.Trace("get a clear")
		b, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		if m["admin_key"] != "114514" { //admin_key 不正确
			context.JSON(http.StatusOK, gin.H{
				"code": 444,
				"msg":  "Who are you?operation denied",
				"data": nil,
			})
			loggenerator.Warn("someone else tried to clear the database but was denied")
		} else { //admin_key 正确，删除数据
			Mydatabase.Remove_all_data()
			context.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "",
				"data": gin.H{
					"msg": "operation succeeded",
				},
			})
			loggenerator.Info("clear succeed")
		}
	})
}
