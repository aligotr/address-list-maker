package config

import (
	"address-list-maker/utils"
	"log"
	"os"
	"strconv"
)

var (
	ListsPath                      = "/srv/app/public"
	IpV6Enable                     = false
	Platforms                      = []string{"plaintext", "proxmox", "mikrotik"}
	CrowdsecEnable                 = true
	CrowdsecCron                   = "0,14,28,42 * * * *"
	CrowdsecAPIUrl                 = "http://crowdsec:8080/"
	CrowdsecAPIKey                 = "1234567890"
	CrowdsecStartDelay             = 5
	CrowdsecOrigins                = []string{"crowdsec"}
	CrowdsecScenariosContaining    = []string{}
	CrowdsecScenariosNotContaining = []string{}
	GeoipEnable                    = true
	GeoipCron                      = "0 21 * * 6"
	GeoipCountries                 = []string{"ru", "tr", "eg"}
)

func setStringEnv(key string, target *string) {
	if v := os.Getenv(key); v != "" {
		*target = v
	}
}

func setBoolEnv(key string, target *bool) {
	if v := os.Getenv(key); v != "" {
		val, err := strconv.ParseBool(v)
		if err != nil {
			log.Fatalf("%s, имеет не верное значение: %s", key, v)
		}
		*target = val
	}
}

func setIntEnv(key string, target *int) {
	if v := os.Getenv(key); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("%s, имеет не верное значение: %s", key, v)
		}
		*target = val
	}
}

func setStringSliceEnv(key string, target *[]string) {
	if v := os.Getenv(key); v != "" {
		slice, err := utils.GetArrayFromString(v, ",")
		if err != nil {
			log.Fatalf("%s: %v", key, err)
		}
		*target = slice
	}
}

func setRequiredStringEnv(key string, target *string) {
	if v := os.Getenv(key); v != "" {
		*target = v
	} else {
		log.Fatalf("%s, не задан!", key)
	}
}

// Инициализация конфигов
func InitConfig() {
	setStringEnv("LISTS_PATH", &ListsPath)
	setBoolEnv("IP_V6_ENABLE", &IpV6Enable)
	setStringSliceEnv("PLATFORMS", &Platforms)
	setBoolEnv("CROWDSEC_ENABLE", &CrowdsecEnable)
	setStringEnv("CROWDSEC_CRON", &CrowdsecCron)
	setStringEnv("CROWDSEC_API_URL", &CrowdsecAPIUrl)
	setRequiredStringEnv("CROWDSEC_API_KEY", &CrowdsecAPIKey)
	setIntEnv("CROWDSEC_START_DELAY", &CrowdsecStartDelay)
	setStringSliceEnv("CROWDSEC_ORIGINS", &CrowdsecOrigins)
	setStringSliceEnv("CROWDSEC_SCENARIOS_CONTAINING", &CrowdsecScenariosContaining)
	setStringSliceEnv("CROWDSEC_SCENARIOS_NOT_CONTAINING", &CrowdsecScenariosNotContaining)
	setBoolEnv("GEOIP_ENABLE", &GeoipEnable)
	setStringEnv("GEOIP_CRON", &GeoipCron)
	setStringSliceEnv("COUNTRIES", &GeoipCountries)
}
