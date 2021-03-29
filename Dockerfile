FROM golang:alpine as builder

WORKDIR /github.com/hieronimusbudi/go-fiber-bookstore-order-api
COPY go.mod go.sum ./
RUN go mod download && go get github.com/codegangsta/gin
COPY . .

CMD gin --immediate -a 9010 -p 9011 run server.go