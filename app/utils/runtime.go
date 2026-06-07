package utils

import "runtime"

/*
Получить имя функции
  - skip(int): Количество фреймов стека, на которое нужно подняться, при этом 0 указывает вызывающего абонента Caller.
*/
func GetFuncName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name()
}
