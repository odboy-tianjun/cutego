package common

import "github.com/druidcaesa/gotool"

func ErrorLog(v ...interface{}) {
	gotool.Logs.ErrorLog().Println(v)
}

func FatalfLog(format string, v ...interface{}) {
	gotool.Logs.ErrorLog().Fatalf(format, v)
}

func InfoLog(v ...interface{}) {
	gotool.Logs.InfoLog().Println(v)
}

func InfoLogf(format string, v ...interface{}) {
	gotool.Logs.InfoLog().Printf(format+"\n", v)
}

func DebugLogf(format string, v ...interface{}) {
	gotool.Logs.DebugLog().Printf(format+"\n", v)
}
