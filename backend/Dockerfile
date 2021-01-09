# Global argument(s)
ARG APP=app

# Stage 1: Build the backend
FROM golang:1.15.6-alpine AS build

ARG APP
ADD . ${GOPATH}/src/${APP}
WORKDIR ${GOPATH}/src/${APP}

# Build with removing the symbols and debug info
RUN go mod verify && \
    go get github.com/swaggo/swag/cmd/swag

RUN swag init -g main.go && \
    go build -ldflags "-s -w" -mod vendor -o app

# Stage 2: Only copy the compiled binary and config to run the backend
FROM alpine:3.12.3
# v3.12.3 was created on Dec 16, 2020

ARG APP
RUN apk --no-cache add ca-certificates
RUN mkdir -p /home/nobody/app
WORKDIR /home/nobody/app
USER nobody
COPY --from=build --chown=nobody:nobody /go/src/${APP}/app ./
CMD ["./app"]