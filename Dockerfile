ARG BUILD_FROM
FROM golang:latest AS builder

ENV LANG C.UTF-8

ENV GO111MODULE=on \
    CGO_ENABLED=0

ADD . /build

# Move to working directory /build
WORKDIR /build

# Download dependencies using go mod
RUN go mod download

RUN go build -o out

FROM $BUILD_FROM

COPY --from=builder /build/out ./

WORKDIR /app

COPY init.sh ./

RUN chmod a+x init.sh
CMD [ "/app/init.sh" ]