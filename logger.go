package logger

import "fmt"

func Info(message string, params interface{}) {
	fmt.Println("\u001B[36mINFO:\u001B[0m", message, " params: ", params)
}

func Error(message string, params interface{}) {
	fmt.Println("\033[31mERROR:\u001B[0m", message, " params: ", params)
}

func Warning(message string, params interface{}) {
	fmt.Println("\033[33mWARNING:\u001B[0m", message, " params: ", params)
}

func Fatal(message string, params interface{}) {
	fmt.Println("\033[31mFATAL:\u001B[0m", message, " params: ", params)
}
