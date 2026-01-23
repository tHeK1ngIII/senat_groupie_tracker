# Étape 1 : builder avec Go >= 1.25.4
FROM golang:1.25.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server main.go

# Étape 2 : image finale légère
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
ENV PORT=8080
CMD ["./server"]
