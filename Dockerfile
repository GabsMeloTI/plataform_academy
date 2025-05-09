# Etapa de build
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar o código-fonte
COPY . .

# Compilar o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# Etapa final
FROM alpine:latest

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/app .

# Expor a porta que o aplicativo utilizará
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["./app"]
