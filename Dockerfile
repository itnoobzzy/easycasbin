FROM golang:1.19.3-alpine

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /data
COPY . .
EXPOSE 8000
