package loggenerator

func Info(s string) { //关键操作
	mylogger.Info(s)
	// writer.WriteString(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
func Warn(s string) { //警告信息

	mylogger.Warn(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
func Debug(s string) { //一般程序中输出的调试信息

	mylogger.Debug(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
func Trace(s string) { //很细粒度的信息，一般用不到 所以根本不会输出出来 6吧

	mylogger.Trace(s)
	// writer.WriteString(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
func Error(s string) { //错误日志，需要查看原因

	mylogger.Error(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
func Panic(s string) { //记录日志，然后panic

	mylogger.Panic(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
func Fatal(s string) { //致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；

	mylogger.Fatal(s)
	writer.Flush() //每次需要执行一次这个才会刷新出去！
}
