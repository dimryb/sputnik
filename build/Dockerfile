# Собираем в гошке
FROM golang:1.24 AS builder

ENV BIN_FILE=/opt/sputnik/sputnik-app
ENV CODE_DIR=/go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
#COPY go.sum .
RUN go mod download

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /go/src

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN go build \
    -ldflags "$LDFLAGS" \
    -o ${BIN_FILE} cmd/app/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL ORGANIZATION="DimRyb"
LABEL SERVICE="sputnik"
LABEL MAINTAINERS="dimryb@bk.ru"

ENV BIN_FILE=/opt/sputnik/sputnik-app
COPY --from=builder ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE=/etc/sputnik/config.yaml
COPY ./configs/config.yaml ${CONFIG_FILE}

RUN apk add --no-cache curl

CMD ["/opt/sputnik/sputnik-app", "-config", "/etc/sputnik/config.yaml"]
