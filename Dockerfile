FROM golang:1.24.4 AS builder

RUN apt-get update \
    && apt-get install -y \
        wget \
        ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/thavlik/mirbase
COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /mirbase
COPY scripts/download_mirbase.sh /tmp/download_mirbase.sh
RUN chmod +x /tmp/download_mirbase.sh /mirbase \
    && /tmp/download_mirbase.sh \
    && rm -f /tmp/download_mirbase.sh

WORKDIR /go/src/github.com/thavlik/mirbase
COPY cmd cmd
COPY pkg pkg

WORKDIR /go/src/github.com/thavlik/mirbase/cmd
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o mirbase
RUN ./mirbase build -o /mirbase.sqlite \
    && bash -c 'if [[ "$(stat -c%s /mirbase.sqlite)" -lt "100000" ]]; then echo "Database file is too small: $(stat -c%s /mirbase.sqlite) bytes"; ls -al / | grep mirbase; exit 1; fi'

FROM debian:bookworm-slim
COPY --from=builder /mirbase.sqlite /mirbase.sqlite
