package crowdsec

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	csbouncer "github.com/crowdsecurity/go-cs-bouncer"

	"address-list-maker/config"
	"address-list-maker/preset"
	"address-list-maker/utils"
)

const (
	moduleName = "crowdsec"
)

var (
	dir = config.ListsPath + "/" + moduleName
)

func Run() {
	// Задержка перед запуском, на случай если crowdsec-сервер не успел запуститься
	time.Sleep(time.Duration(config.CrowdsecStartDelay) * time.Second)

	// Создание и Расширение массива
	var c cache
	c.v4 = make(map[string][]addressesData)
	c.v6 = make(map[string][]addressesData)

	//
	bouncer := &csbouncer.StreamBouncer{
		APIKey:                 config.CrowdsecAPIKey,
		APIUrl:                 config.CrowdsecAPIUrl,
		TickerInterval:         "11s",
		Origins:                config.CrowdsecOrigins,
		ScenariosContaining:    config.CrowdsecScenariosContaining,
		ScenariosNotContaining: config.CrowdsecScenariosNotContaining,
	}
	if err := bouncer.Init(); err != nil {
		log.Fatal("Bouncer - Ошибка инициализации")
	}

	// Во избежание проблем с разовой загрузкой адресов, timeout, должен быть меньше чем TickerInterval
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		bouncer.Run(ctx)
	}()

	go func() {
		defer wg.Done()

		decisions := <-bouncer.Stream
		if decisions != nil {
			c.decisionProcess(decisions)
		} else {
			utils.LogError("Ошибка получения данных от Crowdsec")
		}
	}()

	wg.Wait()

	//
	utils.RemoveDirContent(dir)

	//
	makeList(preset.Protocols.V4.Name, c.v4)
	if config.IpV6Enable {
		makeList(preset.Protocols.V6.Name, c.v6)
	}
}

func makeList(protocolName string, responseData origin) {
	for _, platform := range config.Platforms {
		for originName, addressesData := range responseData {
			if len(addressesData) == 0 {
				continue
			}

			addresses := buildAddresses(platform, protocolName, addressesData)
			contentFile := preset.BuildContentFile(platform, addresses, moduleName+"-"+originName)
			fileName := preset.GetFileName(moduleName, originName, platform)
			utils.CreateOrTruncFileAndWriteString(dir, fmt.Sprintf(fileName, protocolName), contentFile)
		}
	}
}

func buildAddresses(platform string, protocolName string, addressesData []addressesData) string {
	var sb strings.Builder

	switch platform {
	case preset.Proxmox:
		for _, item := range addressesData {
			line := item.address
			if item.scenario != "" {
				line += " # " + item.scenario
			}
			line += "\n"
			sb.WriteString(line)
		}

	case preset.Mikrotik:
		suffix := ""
		if protocolName == preset.Protocols.V6.Name {
			suffix = "6"
		}

		for _, item := range addressesData {
			line := "/ip" + suffix + " firewall address-list add list=crowdsec address=" + item.address + " comment=\""
			if item.scenario != "" {
				line += item.scenario
			} else {
				line += "crowdsec/uncknown"
			}
			line += "\" timeout=5d;\n"
			sb.WriteString(line)
		}

	default:
		for _, item := range addressesData {
			sb.WriteString(item.address + "\n")
		}
	}

	return sb.String()
}
