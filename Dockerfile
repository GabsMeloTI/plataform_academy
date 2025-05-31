# Etapa de build
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar todo o código-fonte, incluindo o diretório cmd
COPY . .

# Compilar o binário (certifique-se de que o caminho está correto)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Etapa final
FROM alpine:latest

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/main .

# Expor a porta que o aplicativo utilizará
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["./main"]
