# Etapa de build
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar o código-fonte
COPY . .

# Compilar o binário a partir do arquivo main.go em ./cmd/main.go
RUN go build -o main ./cmd/main.go

# Expor a porta que o aplicativo utilizará
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["./main"]
