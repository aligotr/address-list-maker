#!/usr/bin/env sh
set -eu

# Переменные
DEVELOPER="aligotr"
PROJECT_NAME="address-list-maker"
VERSION=$(grep "Version = " ./app/main.go | sed 's/.*= "\(.*\)".*/\1/')
TYPES="dev, prod"

# Функции сборки
dev() {
  docker build -f dev/Dockerfile -t ${DEVELOPER}/${PROJECT_NAME}-dev:latest dev/
}

prod() {
  docker build -f prod/Dockerfile -t ${DEVELOPER}/${PROJECT_NAME}:latest -t ${DEVELOPER}/${PROJECT_NAME}:${VERSION} ./
}

# Вспомогательные функции
default() {
  echo "Выберите тип образа: ${TYPES}"
}

clean() {
  echo "ВНИМАНИЕ! Команда удалит не используемые docker-ресурсы:\n${clean_details}"
  prompt
  docker image prune --all --force
  docker system prune --volumes --force
}

# Утилиты
prompt() {
  read -p "Продолжить выполнение? [y/n] " choice
  case $choice in
  [Yy]*)
    break
    ;;
  *)
    exit
    ;;
  esac
}

help() {
  cat <<EOF
────────────────────────────────────
                             O o
                                o
 ______   ______   ______  [O]__ST
|""""""|_|""""""|_|""""""|_|======}
'-0--0-'"'-0--0-'"'-0--0-'"'000--o\.
────────────────────────────────────
Проект: ${PROJECT_NAME}
Тип: docker
Типы итоговых образов: ${TYPES}
────────────────────────────────────
Вывести типы сборок: types
Удалить не используемые docker-ресурсы: clean
${clean_details}
────────────────────────────────────
EOF
}

clean_details=$(
  cat <<EOF
- Остановленные контейнеры;
- Неиспользуемые образы;
- Неиспользуемые тома;
- Неиспользуемые сети;
- Кэш сборки.
EOF
)

# Сопоставление: Аргумент-Функция
if [ $# -eq 0 ]; then
  default
else
  for arg in "$@"; do
    case $arg in
    "dev")
      dev
      ;;
    "prod")
      prod
      ;;
    "clean")
      clean
      ;;
    "help")
      help
      ;;
    *)
      echo "────────────────────────────────────"
      echo "Неизвестный аргумент: $arg"
      help
      ;;
    esac
  done
fi
