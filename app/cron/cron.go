package cron

import (
	"address-list-maker/config"
	"address-list-maker/crowdsec"
	"address-list-maker/geoip"
	"log"

	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	ranAny := false

	//
	if config.GeoipEnable {
		go geoip.Run()
		_, err := c.AddFunc(config.GeoipCron, func() { go geoip.Run() })
		if err != nil {
			panic(err)
		}

		ranAny = true
	}

	//
	if config.CrowdsecEnable {
		go crowdsec.Run()
		_, err := c.AddFunc(config.CrowdsecCron, func() { go crowdsec.Run() })
		if err != nil {
			panic(err)
		}

		ranAny = true
	}

	//
	if !ranAny {
		log.Fatalln("Завершение работы приложения: все задачи для обработки отключены!")
	}

	c.Start()
}
