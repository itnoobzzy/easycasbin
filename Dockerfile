FROM golang:1.19

ENV GO111MODULE on

WORKDIR /data/easycasbin

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

EXPOSE 8080


