FROM golang:1.19

ENV GO111MODULE on

WORKDIR /app/easycasbin

COPY go.mod ./
COPY go.sum ./

RUN go mod download

EXPOSE 8080


