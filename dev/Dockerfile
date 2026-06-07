FROM golang:alpine

# Метки
ENV PROJECT_NAME="address-list-maker"
LABEL project-name=$PROJECT_NAME
LABEL maintainer="Aligotr"

# Переменные
ENV \
  TZ="Europe/Moscow" \
  PUID="1000" \
  PGID="1000" \
  USERNAME="mediator" \
  APP="/srv/app"
WORKDIR $APP

##########      Настройка основного приложения      ##########
#
RUN \
  echo "--- Установка зависимостей и дополнительных программ ---"; \
  apk add --no-cache \
    su-exec \
    tzdata \
    nano; \
  echo "--- Создание пользователя ---"; \
    adduser -u ${PUID} -g ${PGID} -D ${USERNAME}; \
  echo "--- Очистка системы ---" && \
  rm -rf \
    /tmp/*

##########   Настройка и запуск основного процесса   #########
#
COPY --chmod=755 ./docker-entrypoint.sh /opt/docker-entrypoint.sh
ENTRYPOINT ["/opt/docker-entrypoint.sh"]
