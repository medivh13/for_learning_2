FROM golang:1.20-alpine AS build-env

LABEL maintainer="jody almaida<jody.almaida@gmail.com>"

ENV APP_NAME=for_learning_2
ENV GO111MODULE=on
ENV GOPRIVATE=github.com/medivh13
ENV TZ=Asia/Jakarta
ENV GIT_TERMINAL_PROMPT=0
ENV CGO_ENABLED=0

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    curl \
    busybox-extras \
    make \
    git \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk update && apk add --no-cache coreutils

WORKDIR /src

RUN ls -ls

RUN mkdir -p /src/for_learning_2
COPY . /src/for_learning_2
WORKDIR /src/for_learning_2

RUN go mod tidy -compat=1.20

# RUN mkdir -p bin
RUN go build

# RUN chmod 755 ./bin/for_learning_2

EXPOSE 8080

CMD "./for_learning_2"