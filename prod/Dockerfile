###################################################################
# Создание промежуточного образа
###################################################################
FROM golang:alpine AS builder

# Переменные
ENV PROJECT_NAME="address-list-maker"
WORKDIR /srv/app

##########               Сборка приложения               ##########
#
COPY ./app /srv/app
RUN \
  go mod download; \
  CGO_ENABLED=0 GOOS=linux go build -o build/${PROJECT_NAME} ./main.go

###################################################################
# Создание финального образа
###################################################################
FROM alpine

# Метки
ENV PROJECT_NAME="address-list-maker"
LABEL project-name=$PROJECT_NAME
LABEL maintainer="Aligotr"

# Переменные
ENV \
  TZ="Europe/Moscow" \
  PUID="1000" \
  PGID="1000" \
  USER="mediator" \
  HOME="/srv/app"
WORKDIR ${HOME}

##########      Настройка основного приложения      ##########
#
RUN \
  echo "--- Установка зависимостей и дополнительных программ ---"; \
  apk add --no-cache \
    su-exec \
    tzdata \
    nano; \
  echo "--- Очистка системы ---" && \
  rm -rf \
    /tmp/*

RUN addgroup --gid ${PGID} ${USER} && adduser --disabled-password --uid ${PUID} --ingroup ${USER} ${USER}
RUN chown -R ${PUID}:${PGID} $HOME

COPY --chmod=755 --chown=${PUID}:${PGID} --from=builder /srv/app/build/${PROJECT_NAME} ${HOME}

##########   Настройка и запуск основного процесса   #########
#

COPY --chmod=755 ./prod/docker-entrypoint.sh /opt/docker-entrypoint.sh
ENTRYPOINT ["/opt/docker-entrypoint.sh"]
