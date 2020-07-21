package utils

import (
	"log"
	"runtime"
)

func Warn(message string) {
	logger("warn", message, nil)
}

func Error(message string, err error) {
	logger("error", message, err)
}

func Info(message string) {
	logger("info", message, nil)
}

func Log(level, message, method string, err error) {
	if err != nil {
		log.Printf("level[%s] method[%s] msg='%s' err='%s'\n", level, method, message, err.Error())
	} else {
		log.Printf("level[%s] method[%s] msg='%s'\n", level, method, message)
	}
}

func logger(level, message string, err error) {
	if err != nil {
		log.Printf("level[%s] method[%s] msg='%s' err='%s'\n", level, getCaller(2), message, err.Error())
	} else {
		log.Printf("level[%s] method[%s] msg='%s'\n", level, getCaller(2), message)
	}
}

func getCaller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if ok {
		return runtime.FuncForPC(pc).Name()
	}
	return ""
}
