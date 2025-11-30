FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY "internal" "internal"
COPY "cmd" "cmd"
COPY "web" "web"

# -trimpath removes file system paths from the compiled binary
# -ldflags="-s -w" strips debugging information, reducing binary size
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w" -o scarcity-roll "cmd/main.go"

FROM alpine:latest

WORKDIR /app
ENV GIN_MODE=release
COPY --from=builder /app/scarcity-roll .

EXPOSE 8080

CMD ["./scarcity-roll"]