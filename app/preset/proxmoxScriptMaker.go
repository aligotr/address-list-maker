package preset

import (
	"fmt"
)

/* Получить контент для файла списка, платформа: proxmox */
func GetProxmoxScript(ipsetName string, addresses string) string {
	return fmt.Sprintf(`#!/usr/bin/env sh
set -eu

cd /tmp

#
CLUSTER_FILE="/etc/pve/firewall/cluster.fw"
IPSET_NAME="%s"
IPSET_PATTERN="\[IPSET ${IPSET_NAME}\]"
TEMP_LIST_FILE="/tmp/${IPSET_NAME}_list.tmp"
TEMP_FILE="/tmp/cluster.fw.tmp"

push() {
  cp -f "$CLUSTER_FILE" "$CLUSTER_FILE.bak"

  #
  if grep -q "${IPSET_PATTERN}" "$CLUSTER_FILE"; then
    pvesh delete /cluster/firewall/ipset/"${IPSET_NAME}" --force="true"
  fi
  pvesh create /cluster/firewall/ipset --name="${IPSET_NAME}"

  #
  echo "$ADDRESSES" > "$TEMP_LIST_FILE" || return 1

  #
  awk -v file="$TEMP_LIST_FILE" -v pattern="$IPSET_PATTERN" '{
    print;
    if ($0 ~ pattern) {
      while ((getline line < file) > 0) {
        print line;
      }
    }
  }' "$CLUSTER_FILE" > "$TEMP_FILE" || return 1
  cp -f "$TEMP_FILE" "$CLUSTER_FILE" || return 1

  #
  rm -f "$TEMP_FILE"
  rm -f "$TEMP_LIST_FILE"
}

ADDRESSES=$(cat << EOF

%s
EOF
)

#
date '+%%d-%%m-%%Y, %%H:%%M'
if push; then
  systemctl restart pve-firewall.service
	echo "Ipset - ${IPSET_NAME} - Список адресов обновлён"
else
  cp -f "$CLUSTER_FILE.bak" "$CLUSTER_FILE"
	echo "Ipset - ${IPSET_NAME} - Ошибка обновления"
fi
`, ipsetName, addresses)

}
