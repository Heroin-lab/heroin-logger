package logger

import (
	"fmt"
	"time"
)

func getTime() string {
	fullTime := time.Now().Format("01-02-2006 15:04:05")
	return fullTime
}

func Info(message string, params interface{}) {
	currentTime := getTime()
	fmt.Println(currentTime, "\u001B[36m[INFO]\u001B[0m", message, " params: ", params)
}

func Error(message string, params interface{}) {
	currentTime := getTime()
	fmt.Println(currentTime, "\033[31m[ERROR]\u001B[0m", message, " params: ", params)
}

func Warning(message string, params interface{}) {
	currentTime := getTime()
	fmt.Println(currentTime, "\033[33m[WARNING]\u001B[0m", message, " params: ", params)
}

func Fatal(message string, params interface{}) {
	currentTime := getTime()
	fmt.Println(currentTime, "\033[31m[FATAL]\u001B[0m", message, " params: ", params)
}
