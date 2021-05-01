ARG BUILD_FROM
FROM $BUILD_FROM

ENV LANG C.UTF-8

COPY --from=golang:alpine /usr/local/go/ /usr/local/go/

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Move to working directory /build
WORKDIR /build

# Copy static assets into the container
COPY go.mod go.sum ./

# Download dependencies using go mod
RUN go mod download

# Copy the code into the container
COPY . ./cmd
COPY init.sh ./

# Build the application
RUN cd cmd && go build -o /out 

RUN rm -r /usr/local/go/

RUN chmod a+x init.sh
CMD [ "/init.sh" ]