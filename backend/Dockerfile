# -------- Stage 1: Build --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# -------- Stage 2: Runtime --------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
