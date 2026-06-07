#!/usr/bin/env sh
set -eu

#####
cat <<EOF
────────────────────────────────────
                             O o
                                o
 ______   ______   ______  [O]__ST
|""""""|_|""""""|_|""""""|_|======}
'-0--0-'"'-0--0-'"'-0--0-'"'000--o\.
────────────────────────────────────
UID: $(id -u $USERNAME)
GID: $(id -g $USERNAME)
────────────────────────────────────
EOF

# Изменение владельца рабочей папки
if [ -d "$APP" ]; then
  echo "Проверка прав: $APP"
  find "$APP" \
    -path "/srv/app/src" -prune -o \
    ! -user $PUID \
    -exec sh -c 'printf "chown: %s\n" {}' \; \
    -exec chown $PUID:$PGID {} +
else
  echo "Путь не найден, пропускаем: $APP"
fi

echo "────────────────────────────────────"

# Запуск основного процесса
MAIN_OBJECT="$APP/go.mod"
OTHER_OBJECT_1="$APP/go.sum"

if [ ! -f $MAIN_OBJECT ]; then
  echo "Инициализация приложения"
  exec su-exec $PUID:$PGID go mod init $PROJECT_NAME
  echo "────────────────────────────────────"
fi

if [ ! -f $OTHER_OBJECT_1 ]; then
  echo "Установка пакетов"
  exec su-exec $PUID:$PGID go mod download &
  wait
  echo "────────────────────────────────────"
fi

if [ -f $MAIN_OBJECT ]; then
  echo "Основной процесс: exec su-exec ${PUID}:${PGID} tail -f /dev/null"
  exec su-exec $PUID:$PGID tail -f /dev/null
else
  echo "Основной объект программы не найден: ${MAIN_OBJECT}"
fi
