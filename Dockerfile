# syntax=docker/dockerfile:1

################ GOLANG ################

FROM golang:alpine as go_builder

LABEL stage=backend_builder

ENV CGO_ENABLED 0
ENV GOOS linux

COPY ./api /build/api
COPY ./cmd /build/cmd
COPY ./docs /build/docs
COPY ./internal /build/internal
COPY ./api /build/api
COPY ./go.mod /build/go.mod
COPY ./go.sum /build/go.sum

WORKDIR /build

RUN mkdir result
RUN apk update --no-cache && apk add --no-cache tzdata
RUN go mod download
RUN go build -ldflags="-s -w" -o ./result ./cmd/SettingsService

################ RUN ##################

FROM alpine:latest as runner

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=go_builder build/result StadiumSlotBot/

ENV StadiumSlotBotPort 9000
ENV StadiumSlotBotDbConnectionString postgres://pg:1@host.docker.internal:5432/stadiumSlotBot_db
ENV StadiumSlotBotEnv dev

WORKDIR /StadiumSlotBot

##EXPOSE 9000

ENTRYPOINT ["./StadiumSlotBot"]