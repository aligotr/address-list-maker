package preset

const (
	Plaintext = "plaintext"
	Proxmox   = "proxmox"
	Mikrotik  = "mikrotik"
)

var platforms = map[string]platform{
	Plaintext: {
		extension: "zone",
	},
	Proxmox: {
		extension: "sh",
	},
	Mikrotik: {
		extension: "rsc",
	},
}

var Protocols = protocols{
	V4: protocol{
		Name: "v4",
		Sep:  ".",
	},
	V6: protocol{
		Name: "v6",
		Sep:  ":",
	},
}

/* Получить строку с именем файла, для текущего генератора списков */
func GetFileName(filePrefix, additionalInfo string, platform string) string {
	if platform == Plaintext {
		return filePrefix + "_" + additionalInfo + "_%s." + platforms[platform].extension
	}
	return filePrefix + "_" + additionalInfo + "_" + platform + "_%s." + platforms[platform].extension
}

/* Получить контент для файла списка, с распределением по типам платформ */
func BuildContentFile(platform string, addresses string, ipsetName string) string {
	switch platform {
	case Proxmox:
		return GetProxmoxScript(ipsetName, addresses)

	default:
		return addresses
	}
}
