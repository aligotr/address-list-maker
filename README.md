# Address List Maker

Генератор адресных списков.

Источники: IPdeny, CrowdSec
Конечные системы: MikroTik, Proxmox.

Примеры выходных файлов:
`ipdeny_ru_v4.zone`, `ipdeny_ru_proxmox_v4.sh`, `crowdsec_plus_mikrotik_v4.rsc`.

## Примеры использования

### Proxmox - Cron

```cron
00 05 * * 5 root wget -O - https://mydomain.local/geoip/ipdeny_ru_proxmox_v4.sh | sh > /var/log/geoip-ru.log 2>&1
```

### Mikrotik - Script

```routeros
:local name "[crowdsec]"
:local url "https://mydomain.local/crowdsec/crowdsec_plus_mikrotik_v4.rsc"
:local fileName "blocklist.rsc"
:log info "$name fetch blocklist from $url"
/tool fetch url="$url" mode=https dst-path=$fileName idle-timeout="30s"
:if ([:len [/file find name=$fileName]] > 0) do={
    :log info "$name removing old ipv4 blocklist"
     /ip/firewall/address-list/remove [ find where list="crowdsec" ];
    :log info "$name import;start"
    /import file-name=$fileName
    :log info "$name import:done"
} else={
    :log error "$name failed to fetch the blocklist"
}
/file remove $fileName
```
