FROM golang:1.17.2-alpine AS build
WORKDIR /http_server_demo
COPY go.mod /http_server_demo
RUN go mod download
COPY /http_server_demo/*.go /http_server_demo
RUN go build -o /http_server_demo/httpserver

FROM alpine:latest
WORKDIR /http_server_demo
COPY --from=build /http_server_demo/httpserver /http_server_demo/
EXPOSE 8082
ENTRYPOINT ["./httpserver"]
