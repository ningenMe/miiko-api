FROM golang:1.19.1-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /miiko-api

CMD ["/miiko-api"]
