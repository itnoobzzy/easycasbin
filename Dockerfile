FROM golang:1.19

ENV GO111MODULE on

WORKDIR /app/easycasbin

COPY . .

RUN go mod download

RUN go build -o /easycasbin

EXPOSE 8080
CMD ["/easycasbin"]

