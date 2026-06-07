package geoip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"address-list-maker/config"
	"address-list-maker/preset"
	"address-list-maker/utils"
)

const (
	filePrefix = "ipdeny"
)

var (
	dir = config.ListsPath + "/" + "geoip"
	v4  = preset.Protocols.V4
	v6  = preset.Protocols.V6
)

func Run() {
	var c cache

	//
	utils.RemoveDirContent(dir)

	//
	c.makeList(v4.Name, v4.Sep, "https://www.ipdeny.com/ipblocks/data/countries/%s.zone")
	if config.IpV6Enable {
		c.makeList(v6.Name, v6.Sep, "https://www.ipdeny.com/ipv6/ipaddresses/blocks/%s.zone")
	}
}

func (c *cache) makeList(protocolName string, protocolSep string, baseUrl string) {
	for _, country := range config.GeoipCountries {
		url := fmt.Sprintf(baseUrl, country)
		// utils.LogDebug(url)

		err := c.fetchString(url)
		if err != nil {
			utils.LogError(err.Error())
			continue
		}

		//
		if len(c.responseBody) == 0 {
			utils.LogError("Список адресов - пуст <%s>: %s", c.responseBody)
			continue
		}

		//
		addresses := c.checkResponse(protocolSep)

		//
		for _, platform := range config.Platforms {
			contentFile := preset.BuildContentFile(platform, addresses, country)
			fileName := preset.GetFileName(filePrefix, country, platform)
			utils.CreateOrTruncFileAndWriteString(dir, fmt.Sprintf(fileName, protocolName), contentFile)
		}
	}
}

func (c *cache) fetchString(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	c.responseBody = string(body)
	// utils.LogDebug(c.responseBody)

	return nil
}

func (c *cache) checkResponse(sep string) string {
	lines, _ := utils.GetArrayFromString(c.responseBody, "\n")
	// utils.LogDebug("%s", lines)

	var sb strings.Builder
	for _, line := range lines {
		if _, _, err := net.ParseCIDR(line); err != nil {
			continue
		}

		if strings.Contains(line, sep) {
			sb.WriteString(line + "\n")
		}
	}

	return sb.String()
}
