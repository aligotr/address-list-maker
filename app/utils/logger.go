package utils

import (
	"fmt"
	"log"
)

/* Вывод рабочей информации в консоль */
func LogInfo(message string, v ...interface{}) {
	fmt.Printf(message+"\n", v...)
}

/* Вывод ошибки в консоль */
func LogError(message string, v ...interface{}) {
	funcName := GetFuncName(2)
	log.SetFlags(log.LstdFlags)
	log.Printf("[Error] - "+funcName+"\n"+message+"\n", v...)
}

/* Вывод отладочной информации в консоль */
func LogDebug(message string, v ...interface{}) {
	funcName := GetFuncName(2)
	log.SetFlags(log.LstdFlags)
	log.Printf("[Debug] - "+funcName+"\n"+message+"\n\n", v...)
}
