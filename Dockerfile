# syntax=docker/dockerfile:1

FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /ecommerce-backend

EXPOSE 8080

CMD [ "/ecommerce-backend" ]
