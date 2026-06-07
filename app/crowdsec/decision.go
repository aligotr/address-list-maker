package crowdsec

import (
	"address-list-maker/preset"
	"fmt"
	"strings"

	"github.com/crowdsecurity/crowdsec/pkg/models"
)

func (c *cache) decisionProcess(streamDecision *models.DecisionsStreamResponse) {
	for _, decision := range streamDecision.New {
		c.add(decision)
	}
}

func (c *cache) add(decision *models.Decision) {
	isCrowdsecurity := strings.Contains(*decision.Scenario, "crowdsecurity")

	origin := strings.ToLower(*decision.Origin)
	if origin == "capi" && !isCrowdsecurity {
		origin = "plus"
	}
	if origin == "crowdsec" {
		origin = "plus"
	}

	scenario := ""
	if !isCrowdsecurity {
		scenario = *decision.Scenario
	}

	responseData := addressesData{
		address:  "",
		scenario: scenario,
	}

	if strings.Contains(*decision.Value, preset.Protocols.V6.Sep) {
		responseData.address = fmt.Sprintf("%s/128", *decision.Value)
		c.v6[origin] = append(c.v6[origin], responseData)
	} else if strings.Contains(*decision.Value, preset.Protocols.V4.Sep) {
		responseData.address = *decision.Value
		c.v4[origin] = append(c.v4[origin], responseData)
	}
}
