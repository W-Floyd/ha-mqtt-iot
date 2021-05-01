ARG BUILD_FROM
FROM golang:latest AS builder

ENV LANG C.UTF-8

# Move to working directory /build
WORKDIR /build

# Copy static assets into the container
COPY go.mod go.sum ./
# Download dependencies using go mod
RUN go mod download

RUN go build -o out

FROM $BUILD_FROM

WORKDIR /app

COPY init.sh ./

RUN chmod a+x init.sh
CMD [ "/app/init.sh" ]