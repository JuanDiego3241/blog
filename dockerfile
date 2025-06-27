# 1. Etapa de construcción
FROM golang:1.24-alpine AS builder

# Instala herramientas necesarias
RUN apk add --no-cache git

WORKDIR /app

# Copia go.mod y go.sum y descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del proyecto
COPY . .

# Construye el binario optimizado
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/main.go

# 2. Imagen final
FROM alpine:latest

# Instala certificados (útiles para TLS o Postgres)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia el binario compilado y el .env
COPY --from=builder /app/bin/server .
COPY .env .env

# Expone el puerto configurado (por defecto 8080)
EXPOSE 8080

# Comando por defecto
CMD ["./server"]
