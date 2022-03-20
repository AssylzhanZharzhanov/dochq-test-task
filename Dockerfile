FROM golang:1.18-alpine

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY ./.bin/app .
COPY ./configs ./configs