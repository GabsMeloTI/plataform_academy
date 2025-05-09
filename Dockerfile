# Etapa de Build
FROM golang:1.22.2-alpine AS builder

# Diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código-fonte
COPY . .

# Compilar o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Etapa Final
FROM alpine:latest

# Instalar certificados SSL
RUN apk --no-cache add ca-certificates

# Diretório de trabalho
WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/app .

# Expor a porta da aplicação
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./app"]
