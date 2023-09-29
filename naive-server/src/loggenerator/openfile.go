package loggenerator

import (
	"bufio"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var writer *bufio.Writer
var file *os.File
var mylogger *logrus.Logger

// var mylogrus logrus.new()
func OpenFile() {
	timestr := time.Now().Format("2006-01-02---15-04-05") //windows can't use : as a part of filename
	timestr += "logs.txt"
	// os.Create()
	filePath := "../../logs/" + timestr
	// fmt.Println(filePath)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //创建日志文件
	if err != nil {
		Fatal("open file error")
		return
	}
	writer = bufio.NewWriter(file)
	mylogger = logrus.New()
	mylogger.SetOutput(writer)

	Info("Welcome to naive-server")
	Info("version:1.0.0")
	Info("system:" + runtime.GOOS)
	//创建日志文件，写入日志
}
func Turn_Trace_mod() {
	mylogger.Level = logrus.TraceLevel
}
func Turn_Warn_mod() {
	mylogger.Level = logrus.WarnLevel
}
func CloseFile() { //最后close!!
	file.Close()
}
