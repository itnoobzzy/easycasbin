FROM golang:1.19

ENV GO111MODULE on

WORKDIR /data/easycasbin

COPY . .

RUN go mod download

EXPOSE 8080


