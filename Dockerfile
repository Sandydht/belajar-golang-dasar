# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /belajar-golang-dasar

COPY . /belajar-golang-dasar/

CMD ["go", "run", "."]
