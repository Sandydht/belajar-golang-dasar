# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o /belajar-golang-dasar

EXPOSE 8080

CMD [ "/belajar-golang-dasar" ]
