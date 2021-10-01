FROM golang:1.14.4 AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:3.12.0
ARG APP_VERSION
ARG APP_REVISION
RUN apk add --no-cache tzdata
ENV TZ Asia/Tokyo
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.1/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe
ENV APP_VERSION ${APP_VERSION}
ENV APP_REVISION ${APP_REVISION}
COPY --from=builder /build/main /usr/local/bin/server
EXPOSE 50051
ENTRYPOINT ["/usr/local/bin/server"]
