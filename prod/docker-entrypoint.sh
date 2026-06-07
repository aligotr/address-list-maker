#!/usr/bin/env sh
set -eu

#
cat <<EOF
────────────────────────────────────
                             O o
                                o
 ______   ______   ______  [O]__ST
|""""""|_|""""""|_|""""""|_|======}
'-0--0-'"'-0--0-'"'-0--0-'"'000--o\.
────────────────────────────────────
Пользователь контейнера
UID:  $(id -u)
GID:  $(id -g)
────────────────────────────────────
Пользователь основного процесса
UID:  ${PUID}
GID:  ${PGID}
────────────────────────────────────
EOF

# Изменение владельца рабочей папки
if [ -d "$HOME" ]; then
  echo "Проверка прав: $HOME"
  find "$HOME" ! -user $PUID \
    -exec sh -c 'printf "chown: %s\n" {}' \; \
    -exec chown $PUID:$PGID {} +
else
  echo "Путь не найден, пропускаем: $HOME"
fi

echo "────────────────────────────────────"

# Запуск основного процесса
MAIN_OBJECT="${HOME}/${PROJECT_NAME}"
LOG_PATH="${HOME}/logs"

if [ -f $MAIN_OBJECT ]; then
  [ ! -d $LOG_PATH ] && mkdir ${LOG_PATH}
  [ -f ${LOG_PATH}/${PROJECT_NAME}.log ] && truncate -s 0 ${LOG_PATH}/${PROJECT_NAME}.log

  exec su-exec $PUID:$PGID $MAIN_OBJECT 2>&1 | tee -ai ${LOG_PATH}/${PROJECT_NAME}.log
else
  echo "Основной объект программы не найден: ${MAIN_OBJECT}"
fi
