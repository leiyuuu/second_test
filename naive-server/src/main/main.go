package main

import (
	"github.com/leiyuuu/second-test/naive-server/src/Mydatabase"
	"github.com/leiyuuu/second-test/naive-server/src/Myserver"
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

func main() {
	Mydatabase.Set_auto_sign_in(true) //可选：注册账户后自动登录

	loggenerator.OpenFile() //初始化：建立文件与logrus绑定writer

	//可选：需要细微步骤信息时使用trace_mod;warn及以上信息时用warn_mod;info信息都注释掉就行
	loggenerator.Turn_Trace_mod()
	// loggenerator.Turn_Warn_mod()
	loggenerator.Info("Bind log files succeed")

	Mydatabase.Gendatabase() //建立数据库和初始化数据库
	loggenerator.Info("Gendatabase ALL succeed")

	Myserver.Genserver() //开启服务器
	loggenerator.Info("server gen ALL succeed")
	// loggenerator.Trace("为什么进了server你就不行了???")
	// 妈的，是trace正常根本不会输出出来
	defer loggenerator.CloseFile()
}
