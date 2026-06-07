package main

import (
	"address-list-maker/config"
	"address-list-maker/cron"
)

var (
	AppName = "address-list-maker"
	Version = "1.0.1"
)

func main() {
	// Инициализация параметров приложения
	config.InitConfig()

	// Запуск основных горутин с некоторой перодичностью
	cron.Run()

	//
	select {}
}
